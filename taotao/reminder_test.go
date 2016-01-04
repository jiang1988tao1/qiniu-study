package remiander

import (
	"fmt"
	"testing"
)

func Test_Reminder1(t *testing.T) {
	if i, e := Reminder(2, 2); i != 0 || e != nil {
		fmt.Printf("something wrong\n")
		t.Error("something wrong\n")
	} else {
		fmt.Printf("the first go is ok\n")
		t.Log("the first go is ok\n")
	}

}

func Test_Reminder2(t *testing.T) {
	if _, e := Reminder(2, 0); e == nil {
		fmt.Printf("something  wrong\n")
		t.Error("something  wrong\n")
	} else {
		fmt.Printf("tips: %s \n", e)
		fmt.Printf("the second go is ok\n")

		t.Log("the second go is ok\n")
	}
}

func Test_Reminder3(t *testing.T) {
	if i, e := Reminder(3, 2); i != 1 || e != nil {
		fmt.Printf("something wrong\n")
		t.Error("something wrong\n")
	} else {
		fmt.Printf("the third go is ok\n")
		t.Log("the third go is ok\n")
	}
}
func Test_Reminder4(t *testing.T) {
	if _, e := Reminder(1, 2); e == nil {
		fmt.Printf("something  wrong\n")
		t.Error("something  wrong\n")
	} else {
		fmt.Printf("tips: %s \n", e)
		fmt.Printf("the second go is ok\n")

		t.Log("the second go is ok\n")
	}
}

