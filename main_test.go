package main

import (
	"os"
	"testing"
)

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
	err := run([]string{"thedir"})
	if err != nil {
		t.Fatal(err)
	}
}
