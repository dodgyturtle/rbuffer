package rbuffer

import "fmt"

type Buffer interface {
	Load(data []int) error
}

type RBuffer struct {
	buffer []int
}

func NewRBuffer() *RBuffer {
	return &RBuffer{}
}

func (r *RBuffer) Load(data []int) error {
	r.buffer = append(r.buffer, data...)
	fmt.Printf("Buffer %+v \n", r.buffer)
	fmt.Printf("Buffer lenght %d\n", len(r.buffer))
	return nil
}
