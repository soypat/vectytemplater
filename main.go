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
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "error in run: ", err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stdout, "finished succesfully")
}

func run(args []string) (err error) {
	template := flag.String("t", "default", "template directory.")
	err = flag.CommandLine.Parse(args)
	if err != nil {
		return err
	}

	const perm os.FileMode = 0777
	rebed.FolderMode = perm

	switch *template {
	case "default":
		err = rebed.Create(defaultFS, ".")
	default:
		err = errors.New("template " + *template + " not found")
	}
	if err != nil {
		return err
	}
	err = os.Rename("_templates/"+*template, *template)
	if err != nil {
		return err
	}
	err = os.RemoveAll("_templates")
	if err != nil {
		return err
	}

	return nil
}
