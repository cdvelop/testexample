package testexample

import (
	"testing"
)

func Test_addNumbers(t *testing.T) {
	if addNumbers([]int{3, 7}) != 10 {
		t.FailNow()
	}
}
