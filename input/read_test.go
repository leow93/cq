package input

import (
	"bufio"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	t.Run("it returns bytes from a reader", func(t *testing.T) {
		reader := bufio.NewReader(strings.NewReader("hello"))
		err, result := Read(reader)
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
		if string(result) != "hello" {
			t.Errorf("Expected 'hello', got %s", result)
		}
	})

	t.Run("it can read data > 4KB", func(t *testing.T) {
		longString := strings.Repeat("a", (1024*4)+10)
		reader := bufio.NewReader(strings.NewReader(longString))
		err, result := Read(reader)
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
		if string(result) != longString {
			t.Errorf("Expected %s', got %s", longString, result)
		}
	})
}
