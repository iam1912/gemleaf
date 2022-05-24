package models

import (
	"sync"
	"time"
)

type LeafAlloc struct {
	Key        string
	Step       int
	CurrentPos int
	Buffer     []*LeafSegment
	IsPreload  bool
	UpdatedAt  time.Time
	Ch         map[string]chan struct{}
	mu         sync.Mutex
}

func NewLeafAlloc(leaf *Leaf) *LeafAlloc {
	return &LeafAlloc{
		Key:        leaf.BizTag,
		Step:       leaf.Step,
		CurrentPos: 0,
		Buffer:     make([]*LeafSegment, 0),
		Ch:         make(map[string]chan struct{}),
		IsPreload:  false,
		UpdatedAt:  time.Now(),
	}
}

func (l *LeafAlloc) Lock() {
	l.mu.Lock()
}

func (l *LeafAlloc) Unlock() {
	l.mu.Unlock()
}

func (l *LeafAlloc) Wakeup() {
	l.mu.Lock()
	defer l.mu.Unlock()
	close(l.Ch[l.Key])
	delete(l.Ch, l.Key)
}
