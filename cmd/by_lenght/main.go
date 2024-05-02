package main

import (
	"context"
	"fmt"

	rslice "custom-buffer/pkg"
	"custom-buffer/pkg/observers"
	"custom-buffer/pkg/rbuffer"
)

func main() {
	rb := rbuffer.NewRBuffer()
	lo := observers.NewLoadByLenght(3, rb)
	rs := rslice.NewRSlice(context.Background(), lo)
	fmt.Println("Add 1, 2, 3, 4")
	rs.Push(1, 2, 3, 4)
	fmt.Println("Add 5, 6")
	rs.Push(5, 6)
	fmt.Println("Add 7, 8, 9")
	rs.Push(7, 8, 9)

}
