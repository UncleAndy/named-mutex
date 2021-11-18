package named

import "sync"

type Mutex string

var (
	namedMutexes = make(map[Mutex]*sync.Mutex)
	lockCounters = make(map[Mutex]int)
	mapMutex     = sync.Mutex{}
)

func (nm Mutex) Lock() {
	mapMutex.Lock()

	m, ok := namedMutexes[nm]
	if !ok {
		m = &sync.Mutex{}
		namedMutexes[nm] = m
		lockCounters[nm] = 0
	}
	lockCounters[nm]++

	mapMutex.Unlock()

	m.Lock()
}

func (nm Mutex) Unlock() {
	mapMutex.Lock()

	m, ok := namedMutexes[nm]
	if !ok {
		mapMutex.Unlock()
		return
	}
	lockCounters[nm]--

	// Remove unused mutexes
	if lockCounters[nm] <= 0 {
		delete(namedMutexes, nm)
		delete(lockCounters, nm)
	}

	mapMutex.Unlock()

	m.Unlock()
}
