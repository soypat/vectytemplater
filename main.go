package main

import (
	"embed"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/soypat/rebed"
)

var (
	//go:embed _templates/default
	defaultFS embed.FS
)

func main() {
	template := flag.String("t", "default", "template directory.")
	var err error
	switch *template {
	case "default":
		err = rebed.Create(defaultFS, ".")
	default:
		err = errors.New("template " + *template + " not found")
	}
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
	err = os.Rename("_templates/"+*template, *template)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
	err = os.RemoveAll("_templates")
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Fprintln(os.Stdout, "finished succesfully")
}
