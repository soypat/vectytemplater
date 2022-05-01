package main

import (
	"flag"
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

func TestDefault(t *testing.T) {
	resetFlags()
	err := run([]string{"default"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestWebsocket(t *testing.T) {
	resetFlags()
	err := run([]string{"ws", "-template=websocket-cli"})
	if err != nil {
		t.Fatal(err)
	}
}

var blankCMDLine = *flag.CommandLine

func resetFlags() {
	cmdLine := blankCMDLine
	flag.CommandLine = &cmdLine
}
