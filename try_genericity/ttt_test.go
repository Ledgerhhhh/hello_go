package main

import (
	"sync"
	"testing"
)

type Broadcast[T any] struct {
	msg map[string]chan T
	mut sync.Mutex
}

func NewBroadcast[T any](msg map[string]chan T, mut sync.Mutex) *Broadcast[T] {
	return &Broadcast[T]{msg: msg, mut: mut}
}

func TestTTTT(t *testing.T) {

	broadcast := NewBroadcast(make(map[string]chan int), sync.Mutex{})
	broadcast.mut.Lock()
}
