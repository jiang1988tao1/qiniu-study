package main

import (
	"fmt"
	"./reminder"
)
func main()  {
	if i,e := remiander.Reminder(2,2);e!=nil{
		fmt.Printf("wrong num")
	}else{
		fmt.Printf("reminder is %d",i)
	}
}
