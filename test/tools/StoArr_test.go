package tools

import (
	"laboratory/pkg/utils"
	"testing"
)

type Test struct {
	A string `excel:"h"`
	B string
	C string `excel:"h"`
}

func NewTest() *Test {
	return &Test{
		A: "6",
		B: "7",
		C: "8",
	}
}

func TestStoArr(t *testing.T) {
	res := utils.StoArr(*NewTest())
	t.Error(res)
}
