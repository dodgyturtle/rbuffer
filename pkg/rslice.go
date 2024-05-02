package rslice

import (
	"context"
	"sync"

	"custom-buffer/pkg/observers"
)

type RSlice struct {
	DataFlow  *observers.DataFlow
	observers []observers.Observer
	ctx       context.Context
}

func NewRSlice(ctx context.Context, o ...observers.Observer) *RSlice {
	var mutex sync.Mutex
	return &RSlice{
		DataFlow: &observers.DataFlow{
			Mu:   &mutex,
			Data: []int{},
		},
		observers: o,
		ctx:       ctx,
	}
}

func (r *RSlice) Push(elem ...int) {
	select {
	case <-r.ctx.Done():
		return
	default:
		r.DataFlow.Mu.Lock()
		r.DataFlow.Data = append(r.DataFlow.Data, elem...)
		r.DataFlow.Mu.Unlock()
		r.notifyAll()
	}

}

func (r *RSlice) notifyAll() {
	var wg sync.WaitGroup
	wg.Add(len(r.observers))
	for _, o := range r.observers {
		go o.Run(r.ctx, &wg, r.DataFlow)
	}
	wg.Wait()
}
