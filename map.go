package golang_map_bench

import (
	"sync"
	"sync/atomic"
	"time"
)

// Concurrent safe map for benchmarking
type SafeMaper interface {
	Get() map[int]int
	Produce()
}

type NativeMap struct {
	Table map[int]int
}

func NewNativeMap() *NativeMap {
	m := &NativeMap{}
	m.Table = make(map[int]int)
	return m
}

func (m *NativeMap) Get() map[int]int {
	return m.Table
}

func (m *NativeMap) Produce() {
	tmp := make(map[int]int)
	for i := 0; i < 100; i++ {
		tmp[i] = i + 5566
	}
	m.Table = tmp
}

type AtomicMap struct {
	Table atomic.Value
}

func NewAtomicMap() *AtomicMap {
	m := &AtomicMap{}
	m.Table.Store(make(map[int]int))
	return m
}

func (m *AtomicMap) Get() map[int]int {
	return m.Table.Load().(map[int]int)
}

func (m *AtomicMap) Produce() {
	tmp := make(map[int]int)
	for i := 0; i < 100; i++ {
		tmp[i] = i + 5566
	}
	m.Table.Store(tmp)
}

type RWLockMap struct {
	Table map[int]int
	mtx   sync.RWMutex
}

func NewRWLockMap() *RWLockMap {
	m := &RWLockMap{}
	m.Table = make(map[int]int)
	return m
}

func (m *RWLockMap) Get() map[int]int {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	tmp := m.Table
	return tmp
}

func (m *RWLockMap) Produce() {
	tmp := make(map[int]int)
	for i := 0; i < 100; i++ {
		tmp[i] = i + 5566
	}
	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.Table = tmp
}

type SyncMap struct {
	Table sync.Map
}

func NewSyncMap() *SyncMap {
	m := &SyncMap{}
	return m
}

func (m *SyncMap) Get() map[int]int {
	_, _ = m.Table.Load(0)
	return make(map[int]int)
}

func (m *SyncMap) Produce() {
	for i := 0; i < 100; i++ {
		m.Table.Store(i, i+5566)
	}
}

// StartProducer simulates an periodic timer to update whole map.
func StartProducer(m SafeMaper) {
	go func() {
		for {
			start := time.Now().UnixNano()
			m.Produce()
			time.Sleep(time.Millisecond - time.Duration(time.Now().UnixNano()-start))
		}
	}()
}
