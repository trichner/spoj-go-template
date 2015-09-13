package main

import (
	"testing"
)

func TestStdioMain(t *testing.T) {
	//=== Mock stdout
	mt := mTestPtr(t)
	mt.MockStdio("stdin.txt", "stdout.txt", func(t *testing.T) {
		main()
	})
}
