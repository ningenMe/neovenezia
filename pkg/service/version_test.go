package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetVersion(t *testing.T) {
	t.Run("", func(t *testing.T) {
		actual := getVersion()
		expect := "v0.3.0"
		assert.Equal(t,actual,expect)
	})
}

func ExampleExec() {
	ExecVersion()
	// Output: Neovenezia v0.3.0
}