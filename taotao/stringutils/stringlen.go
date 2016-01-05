package stringutils

import (
	"fmt"
)

func Getstringlen(str string){
	strLength :=len([]rune(str))
	fmt.Printf("the str len: %d",strLength)
	byteLen := len(str)
	fmt.Printf("the str bytes len: %d",byteLen)
}
