package main

import (
	"testing"
)

func Test_Main(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func() {
			print(i)
		}()
	}
}
