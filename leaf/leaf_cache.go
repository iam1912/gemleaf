package leaf

import (
	"sync"
	"time"

	"github.com/iam1912/gemseries/gemleaf/models"
)

type LeafCache struct {
	mu    sync.RWMutex
	cache map[string]*models.LeafAlloc
}

func NewLeafCache() *LeafCache {
	return &LeafCache{
		cache: make(map[string]*models.LeafAlloc, 10),
	}
}

func (l *LeafCache) Get(key string) (*models.LeafAlloc, bool) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	leafAlloc, ok := l.cache[key]
	if !ok {
		return nil, false
	}
	return leafAlloc, true
}

func (l *LeafCache) Add(key string, alloc *models.LeafAlloc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.cache[key] = alloc
}

func (l *LeafCache) Update(key string, alloc *models.LeafAlloc) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	if _, ok := l.cache[key]; ok {
		val := l.cache[key]
		val.Buffer = alloc.Buffer
		val.UpdatedAt = time.Now()
		return true
	}
	return false
}
