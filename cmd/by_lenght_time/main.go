package main

import (
	"context"
	"fmt"
	"time"

	rslice "custom-buffer/pkg"
	"custom-buffer/pkg/observers"
	"custom-buffer/pkg/rbuffer"
)

func main() {
	rb := rbuffer.NewRBuffer()
	lo := observers.NewLoadByLenght(3, rb)
	to := observers.NewLoadByTimeObserver(1000, rb)
	rs := rslice.NewRSlice(context.Background(), lo, to)
	fmt.Println("Add 1, 2, 3, 4")
	rs.Push(1, 2, 3, 4)
	fmt.Println("Add 5, 6")
	rs.Push(5, 6)
	time.Sleep(time.Duration(2) * time.Second)
	fmt.Println("Add 7, 8, 9")
	rs.Push(7, 8, 9)

}
