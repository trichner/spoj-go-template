package main

import (
	"os"
	"testing"
)

type test func(*testing.T)

type mTest testing.T

func (t *mTest) MockStdin(infile string, fn test) {
	old := os.Stdin
	f, err := os.Open(infile)
	if err != nil {
		t.Fatal(err)
	}
	os.Stdin = f
	defer func() {
		f.Close()
		os.Stdin = old
	}()

	fn(testingTPtr(t))
}

func (t *mTest) MockStdout(outfile string, fn test) {
	old := os.Stdout
	w, err := os.Create(outfile)
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w

	fn(testingTPtr(t))

	w.Close()
	os.Stdout = old
}

func (mt *mTest) MockStdio(infile string, outfile string, fn test) {
	mt.MockStdout(outfile, func(t *testing.T) {
		mt := mTestPtr(t)
		mt.MockStdin(infile, fn)
	})
}

func mTestPtr(t *testing.T) *mTest {
	return (*mTest)(t)
}

func testingTPtr(t *mTest) *testing.T {
	return (*testing.T)(t)
}
