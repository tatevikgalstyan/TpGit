package main

import (
	"errors"
	"fmt"
	"math"
)

type myInt int32

func main() {
	var mi myInt
	mi = 30
	fmt.Println(mi)
	fmt.Println(mi.Divide(2))
	fmt.Println(mi.Add(2))
	fmt.Println(mi.Sub(2))
	fmt.Println(mi.Multiply(2))
}
func (mi myInt) Divide(n int) myInt {
	return mi / myInt(n)
}
func (mi myInt) Add(n int) (myInt, error) {
	if n >= math.MaxInt32 || int64(n)+int64(n) > math.MaxInt32 {
		return 0, errors.New("out of range")
	}

	return mi + myInt(n), nil
}
func (mi myInt) Sub(n int) myInt {
	return mi - myInt(n)
}
func (mi myInt) Multiply(n int) myInt {
	return mi * myInt(n)
}
