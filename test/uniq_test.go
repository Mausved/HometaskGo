package test

import "testing"

func TestCheckSuccess(t *testing.T) {
	if 2 == 3 {
		t.Fatalf("Err")
	}
}
