package main

import (
	"fmt"
	"github.com/changhong-qiniu/qiniu-study/taotao/reminder"
)

func main() {
	if i, e := remiander.Reminder(2, 2); e != nil {
		fmt.Printf("wrong num")
	} else {
		fmt.Printf("reminder is %d \n", i)
	}
}
