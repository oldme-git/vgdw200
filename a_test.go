package main

import (
	"fmt"
	"github.com/oldme-git/vgdw200/internal/logic"
	"testing"
)

func TestA(t *testing.T) {
	f, err := logic.NewResource("./1.zip")
	if err != nil {
		panic(err)
	}
	//f.Transmit(func(i uint, bytes []byte, err error) {
	//	if err != nil {
	//		return
	//	}
	//	fmt.Printf("%d\n", i)
	//})
	a := f.GetPkgNum()
	fmt.Println(a)
}
