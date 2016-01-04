package remiander

import (
	"errors"
)

func Reminder(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("y can not be zero")
	}else if x<y{
		return 0, errors.New("x can not < y can not be zero")
	}
	return x % y, nil
}
