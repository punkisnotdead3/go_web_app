package models

import (
	"fmt"
	"math"
	"testing"
	"unsafe"
)

type P1 struct {
	a int8
	b string
	c int8
}

type P2 struct {
	a int8
	c int8
	b string
}

func TestInt64(t *testing.T) {
	fmt.Println("max:", math.MaxInt64)
}

func TestPP(t *testing.T) {
	p1 := P1{
		a: 1,
		b: "b",
		c: 2,
	}

	p2 := P2{
		a: 1,
		c: 2,
		b: "b",
	}

	fmt.Println("p1:", unsafe.Sizeof(p1))
	fmt.Println("p2:", unsafe.Sizeof(p2))

}
