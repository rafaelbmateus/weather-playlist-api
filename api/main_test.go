package main

import "testing"

func TestMain(t *testing.T) {
	if true == false {
		t.Errorf("true is equals false 😲")
	}
}
