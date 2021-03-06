package main

import (
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	sleepTime = 0 * time.Second
	eatTime = 0 * time.Second
	thinkTime = 0 * time.Second

	for i := 0; i < 100; i++ {
		main()
		if len(orderFinished) != len(philosophers) {
			t.Error("wrong number of entries in slice")
		}
		orderFinished = []string{}
	}

}
