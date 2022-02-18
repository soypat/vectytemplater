package main

import (
	"os"
	"testing"
)

const dirperm = 0755

func TestMain(t *testing.M) {
	os.Mkdir("testdata", 0777)
	err := os.Chdir("testdata")
	if err != nil {
		panic(err.Error())
	}
	defer func() {
		err := os.RemoveAll("testdata")
		if err != nil {
			panic(err.Error())
		}
	}()
	code := t.Run()
	if code != 0 {
		panic("non-zero exit code")
	}
}

func TestRun(t *testing.T) {
	err := run([]string{"vectytemplater"})
	if err != nil {
		t.Fatal(err)
	}
}