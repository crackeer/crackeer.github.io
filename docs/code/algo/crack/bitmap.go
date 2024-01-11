package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(0 >> 5)
	object := NewFigure(128)
	object.Set(100)
	object.Set(99)
	object.Set(10)
	fmt.Println(object.Exists(100))
	fmt.Println(object.Exists(99))
	object.Delete(100)
	fmt.Println(object.Exists(100))
	fmt.Println(object.Exists(99))

}

type Figure struct {
	Storage []uint
	MaxSize int
}

func NewFigure(maxSize int) *Figure {
	index, offset := cal(maxSize)
	if offset > 0 {
		index += 1
	}
	return &Figure{
		Storage: make([]uint, index),
		MaxSize: index*8 - 1,
	}
}

func cal(num int) (int, int) {
	return num / 8, num % 8
}

func (f *Figure) Set(num int) bool {
	index, offset := cal(num)
	f.Storage[index] |= 1 << offset
	return true
}

func (f *Figure) Echo() {
	for _, item := range f.Storage {
		fmt.Println(strconv.FormatInt(int64(item), 2))
	}
}

func (f *Figure) Exists(num int) bool {
	index, offset := cal(num)
	return f.Storage[index]>>offset > 0
}

func (f *Figure) Delete(num int) bool {
	index, offset := cal(num)
	f.Storage[index] &= ^(1 << offset)
	return true
}
