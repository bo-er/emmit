package utils

import (
	"testing"
)

func Test_getNumber(t *testing.T) {

	name1 := "MIT6_042JF10_chap02.pdf"
	number1 := getNumber(name1)
	if number1 != 2 {
		t.Error("failed to get number")
	}
	name2 := "MIT6_042JF10_chap22.pdf"
	number2 := getNumber(name2)
	if number2 != 22 {
		t.Error("failed to get number")
	}
}
