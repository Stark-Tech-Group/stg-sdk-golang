package cmd

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}

func TestToDo(t *testing.T) {
	// to do

}
