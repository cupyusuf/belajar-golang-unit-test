package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Yusuf")
	require.Equal(t, "Hello Yusuf", result, "Result must be Hello Yusuf")
	fmt.Println("TestHelloWorldRequire Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Yusuf")
	assert.Equal(t, "Hello Yusuf", result, "Result must be Hello Yusuf")
	fmt.Println("TestHelloWorldAssert Done")
}

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
