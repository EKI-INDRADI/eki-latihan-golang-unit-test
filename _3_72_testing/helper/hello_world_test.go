package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("eki")
	if result != "Hello eki" {
		//  t.Error("Expected 'Hello eki' but got ", result)
		panic("Result is not 'Hello eki'")
	}
}

func TestHelloWorldIndradi(t *testing.T) {
	result := HelloWorld("indradi")
	if result != "Hello indradi" {
		panic("Result is not 'Hello indradi'")
	}
}
