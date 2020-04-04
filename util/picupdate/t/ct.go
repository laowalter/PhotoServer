package main

import (
	"fmt"
	"time"
)

func createDate(ct string) time.Time {
	_time, err := time.Parse("2006:01:02 15:04:05", ct)
	if err != nil {
		_time, err = time.Parse("2006:01:02 15:04:05-07:00", ct)
		if err != nil {
			_ct := []rune(ct)
			if string(_ct[len(_ct)-2:]) == "下午" {
				ct = string(_ct[:len(_ct)-2]) + "PM"
				fmt.Println(ct)
			} else {
				ct = string(_ct[:len(_ct)-2]) + "AM"
				fmt.Println(ct)
			}
			_time, err = time.Parse("2006:01:02 15:04:05PM", ct)
			if err != nil {
				return time.Time{}

			}
		}
	}

	return _time
}
