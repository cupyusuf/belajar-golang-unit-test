package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Yusuf")
	if result != "Hello Yusuf" {
		// error
		t.Error("Result must be Hello Yusuf")
	}

	fmt.Println("TestHelloWorld Done")
}

func TestHelloSupriadi(t *testing.T) {
	result := HelloWorld("Supriadi")
	if result != "Hello Supriadi" {
		// unit test failed
		t.Fatal("Result must be Hello Supriadi")
	}

	fmt.Println("TestHelloSupriadi Done")
}
