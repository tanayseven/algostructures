package main

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello"
	actual := GetHello()
	if expected != actual {
		t.Errorf("Got wrong string from GetHello")
	}
}
