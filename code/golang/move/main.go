package main

import "fmt"

const ByteSize int64 = 8

type Cat struct {
	Data []byte
	Size int64
}

func calcLen(maxSize int64) int64 {
	return maxSize/ByteSize + 1
}

func NewCat(maxSize int64) *Cat {
	return &Cat{
		Data: make([]byte, calcLen(maxSize)),
		Size: maxSize,
	}
}

func (cat *Cat) Set(value int64) bool {
	if value > cat.Size {
		return false
	}

	index := value / ByteSize
	offset := value % ByteSize
	cat.Data[index] |= 1 << offset
	return true
}

func (cat *Cat) Del(value int64) bool {
	if value > cat.Size {
		return false
	}

	index := value / ByteSize
	offset := value % ByteSize
	cat.Data[index] &= 0 << offset
	return true
}

func (cat *Cat) Exists(value int64) bool {
	if value > cat.Size {
		return false
	}

	index := value / ByteSize
	offset := value % ByteSize
	return cat.Data[index]&(1<<offset) > 0
}

func main() {
	cat := NewCat(100)
	cat.Set(99)
	cat.Set(88)

	fmt.Println(cat.Exists(1))
	fmt.Println(cat.Exists(99))
	fmt.Println(cat.Exists(88))
	fmt.Println(cat.Del(88))
	fmt.Println(cat.Exists(88))

}
