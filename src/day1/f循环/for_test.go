package for_test

import "testing"

func TestFor(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(i)
	}
}
