package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Yusuf")
	if result != "Hello Yusuf" {
		// unit test failed
		panic("Result is not Hello Yusuf")
	}
}

func TestHelloSupriadi(t *testing.T) {
	result := HelloWorld("Supriadi")
	if result != "Hello Supriadi" {
		// unit test failed
		panic("Result is not Hello Supriadi")
	}
}
