package leaf

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"github.com/iam1912/gemseries/gemleaf/logger"
	"github.com/iam1912/gemseries/gemleaf/models"
	"gorm.io/gorm"
)

const (
	defaultStep = 2000
	MAXTRY      = 3
)

type GemLeaf struct {
	DB    *gorm.DB
	Cache *LeafCache
	mu    sync.Mutex
}

func NewGemLeaf(db *gorm.DB) *GemLeaf {
	return &GemLeaf{
		DB:    db,
		Cache: NewLeafCache(),
	}
}

func (l *GemLeaf) Create(key string, maxID uint64, step int, desc string) error {
	if maxID == 0 {
		maxID = 1
	}
	if step == 0 {
		step = 2000
	}
	if err := models.LeafCreate(l.DB, key, desc, maxID, step); err != nil {
		logger.Errorf("%s created failed error:%s\n", key, err.Error())
		return err
	}
	logger.Infof("%s created success\n", key)
	return nil
}

func (l *GemLeaf) GetID(key string) (uint64, error) {
	l.mu.Lock()
	alloc, ok := l.Cache.Get(key)
	if !ok {
		var err error
		alloc, err = l.InitLeafAlloc(key)
		if err != nil {
			logger.Errorf("%s init failed error:%s", key, err.Error())
			return 0, err
		}
		logger.Infof("%s success init and cursor is %d\n", key, alloc.Buffer[alloc.CurrentPos].Cursor)
	}
	l.mu.Unlock()
	id, err := l.GetNextID(alloc)
	if err != nil {
		return 0, err
	}
	ok = l.Cache.Update(key, alloc)
	if !ok {
		return 0, errors.New("alloc update failed")
	}
	return id, nil
}

func (l *GemLeaf) InitLeafAlloc(key string) (*models.LeafAlloc, error) {
	leaf, err := models.LeafGet(l.DB, key)
	if err != nil {
		return nil, err
	}
	alloc := models.NewLeafAlloc(leaf)
	alloc.Buffer = append(alloc.Buffer, models.NewLeafSegment(leaf))
	l.Cache.Add(key, alloc)
	return alloc, nil
}

func (l *GemLeaf) GetNextID(alloc *models.LeafAlloc) (uint64, error) {
	alloc.Lock()
	defer alloc.Unlock()
	var id uint64
	currentBuffer := alloc.Buffer[alloc.CurrentPos]
	if currentBuffer.HasID() {
		id = currentBuffer.Cursor
		atomic.AddUint64(&alloc.Buffer[alloc.CurrentPos].Cursor, 1)
	}
	if float32(id/currentBuffer.Max) >= 0.1 && len(alloc.Buffer) == 1 && !alloc.IsPreload {
		alloc.IsPreload = true
		cancel, _ := context.WithTimeout(context.Background(), time.Second*3)
		go l.Preload(cancel, alloc.Key, alloc)
	}
	if id == currentBuffer.Max {
		if len(alloc.Buffer) == 2 && alloc.Buffer[alloc.CurrentPos+1].InitOk {
			alloc.Buffer = alloc.Buffer[1:]
		}
	}
	if id != 0 {
		return id, nil
	}
	ch := make(chan struct{}, 1)
	alloc.Ch[alloc.Key] = ch
	select {
	case <-ch:
		logger.Info("preload is sucess")
	case <-time.After(time.Millisecond * 500):
	}
	if len(alloc.Buffer) <= 1 {
		return 0, errors.New("get failed id")
	}
	alloc.Buffer = alloc.Buffer[1:]
	newBuffer := alloc.Buffer[alloc.CurrentPos]
	if newBuffer.HasID() {
		id = newBuffer.Cursor
		atomic.AddUint64(&alloc.Buffer[alloc.CurrentPos].Cursor, 1)
	}
	return id, nil
}

func (l *GemLeaf) Preload(ctx context.Context, key string, alloc *models.LeafAlloc) error {
	for i := 0; i < MAXTRY; i++ {
		leaf, err := models.LeafNextSegment(l.DB, key)
		if err != nil {
			continue
		}
		segment := models.NewLeafSegment(leaf)
		alloc.Buffer = append(alloc.Buffer, segment)
		l.Cache.Update(key, alloc)
		alloc.Wakeup()
		break
	}
	alloc.IsPreload = false
	return nil
}
