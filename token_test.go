package main

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	var bShow []byte
	for b := 65; b <= 122; b++ {
		bShow = append(bShow, byte(b))
	}
	fmt.Println(string(bShow))

	var bShow2 []byte
	for b := 48; b <= 57; b++ {
		bShow2 = append(bShow2, byte(b))
	}
	fmt.Println(string(bShow2))

}
