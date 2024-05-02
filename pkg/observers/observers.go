package observers

import (
	"context"
	"fmt"
	"sync"
	"time"

	"custom-buffer/pkg/rbuffer"
)

type DataFlow struct {
	Data []int
	Mu   *sync.Mutex
}

type Observer interface {
	Run(ctx context.Context, wg *sync.WaitGroup, d *DataFlow)
}

type LoadByTime struct {
	duration  time.Duration
	buffer    rbuffer.Buffer
	isRunning bool
}

func NewLoadByTimeObserver(ms int, buffer rbuffer.Buffer) *LoadByTime {
	return &LoadByTime{
		duration: time.Duration(ms) * time.Millisecond,
		buffer:   buffer,
	}
}

func (l *LoadByTime) Run(ctx context.Context, wg *sync.WaitGroup, d *DataFlow) {
	defer wg.Done()
	if !l.isRunning {
		ticker := time.NewTicker(l.duration)
		go func() {
			l.isRunning = true
			for {
				select {
				case <-ctx.Done():
					return
				case t := <-ticker.C:
					fmt.Println("Load at", t)
					d.Mu.Lock()
					err := l.buffer.Load(d.Data)
					if err != nil {
						fmt.Printf("error load to buffer %+v\n", d.Data)
					} else {
						d.Data = []int{}
					}
					d.Mu.Unlock()
				}
			}
		}()
	}
}

type LoadByLenght struct {
	lenght int
	buffer rbuffer.Buffer
}

func NewLoadByLenght(lenght int, buffer rbuffer.Buffer) *LoadByLenght {
	return &LoadByLenght{
		lenght: lenght,
		buffer: buffer,
	}
}

func (l *LoadByLenght) Run(ctx context.Context, wg *sync.WaitGroup, d *DataFlow) {
	defer wg.Done()
	select {
	case <-ctx.Done():
		return
	default:
		if len(d.Data) >= l.lenght {
			d.Mu.Lock()
			defer d.Mu.Unlock()
			err := l.buffer.Load(d.Data)
			if err != nil {
				fmt.Printf("error load to buffer %+v\n", d.Data)
				return
			}
			d.Data = []int{}

		}
	}
}
