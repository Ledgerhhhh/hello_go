package test

import "sync"

type Event[T any] struct {
	Type string
	Data T
}
type Broadcast[T any] struct {
	msg map[string]chan Event[T]
	mut sync.Mutex
}

func (b *Broadcast[T]) Subscribe(id string) {
	_, ok := b.msg[id]
	if !ok {
		return
	}
	b.mut.Lock()
	defer b.mut.Unlock()
	_ = make(chan Event[T], 5)

}

func NewBroadcast[T any]() *Broadcast[T] {
	return &Broadcast[T]{
		msg: make(map[string]chan Event[T]),
		mut: sync.Mutex{},
	}
}
