package mutex

type Mutex struct {
	ch chan struct{}
}

func NewMutex() *Mutex {
	return &Mutex{ch: make(chan struct{}, 1)}
}

func (m *Mutex) Lock() {
	m.ch <- struct{}{}
}

func (m *Mutex) Unlock() {
	select {
	case <-m.ch:
	default:
		panic("unlock of unlocked mutex")
	}
}
