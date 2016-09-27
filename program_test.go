package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestRunMainNoHost(t *testing.T) {
	os.Setenv("HOST", "")
	os.Setenv("PORT", "1234")
	assert.Panics(t, func() {
		main()
	})
}

func TestRunMainNoPort(t *testing.T) {
	os.Setenv("HOST", "localhost")
	os.Setenv("PORT", "")
	assert.Panics(t, func() {
		main()
	})
}
