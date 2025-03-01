package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("Yusuf")
	}
}

func BenchmarkHelloWorldSupriadi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("Supriadi")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Yusuf",
			request:  "Yusuf",
			expected: "Hello Yusuf",
		},
		{
			name:     "Supriadi",
			request:  "Supriadi",
			expected: "Hello Supriadi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result, fmt.Sprintf("Name %s must be %s", test.request, test.expected))
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Yusuf", func(t *testing.T) {
		result := HelloWorld("Yusuf")
		assert.Equal(t, "Hello Yusuf", result, "Result must be Hello Yusuf")
	})

	t.Run("Supriadi", func(t *testing.T) {
		result := HelloWorld("Supriadi")
		assert.Equal(t, "Hello Supriadi", result, "Result must be Hello Supriadi")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("Before Unit Test")
	m.Run()
	// after
	fmt.Println("After Unit Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Can't run on Linux")
	}

	result := HelloWorld("Yusuf")
	assert.Equal(t, "Hello Yusuf", result, "Result must be Hello Yusuf")
}
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
