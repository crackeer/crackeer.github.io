## 实现Bitmap

```go
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
```

## 16进制转2进制

```go
package main

import (
	"fmt"
	"math"
  "strings"
)

func foo(slice []int) {
  slice[0] = 100
}

func main() {
  
	fmt.Println(transferBit2Str("11101"))
  fmt.Println(transferNum2Bit(29, 16, "0123456789ABCDEF"))
  fmt.Println(transfer2To16("11101"))
}


func transfer2To16(str string) string {
  if len(str) < 1 {
    return ""
  }

  var (
     mapping string = "0123456789ABCDEF"
     size16 int = 16
     rest int = len(str) % 4
     segmant int = len(str) / 4
  )
 

  result := []string{}
  if rest > 0 {
    result = append(result, transferNum2Bit(transferBit2Str(str[0:rest]), size16, mapping), )
  }

  if segmant > 0 {
    for i := 0; i< segmant; i++ {
       from := i * 4 + rest
          result = append(result, transferNum2Bit(transferBit2Str(str[from:from + 4]), size16, mapping), )

    }
  }

  return strings.Join(result, "")
  
}


func transferNum2Bit(num int, size int, chars string) string {

   if len(chars) < size {
     panic("chars mapping length error")
   }

  result := []byte{}
  for num > 0{
    res := num % size
    result = append(result, chars[res])
    num = num / size
  }
  for i :=0 ;i < len(result)/2;i++ {
    result[i], result[len(result) - 1- i] = result[len(result) - 1- i], result[i]
  }
  return string(result)
}

func transferBit2Str(str string) int {
  offset := len(str) - 1
  result := 0
	for _, b := range str {
    if b == '1' {
      result += int(math.Pow(2, float64(offset)))
    }
    offset--
	}
  return result
}

```