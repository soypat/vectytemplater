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
		fmt.Fprintln(os.Stderr, "error in run:", err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stdout, "finished succesfully")
}

func run(args []string) (err error) {
	if len(args) < 1 {
		usage("")
		return errors.New("no output directory specified")
	}
	template := flag.String("template", "default", "template name.")
	err = flag.CommandLine.Parse(args[1:])
	if err != nil {
		return err
	}
	dir := args[0]
	if _, err := os.Stat(dir); err == nil {
		usage("")
		return errors.New("cannot create directory \"" + dir + "\" as it already exists")
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
	err = os.Rename("_templates/"+*template, dir)
	if err != nil {
		return err
	}
	err = os.RemoveAll("_templates")
	if err != nil {
		return err
	}

	return nil
}

func usage(command string) {
	fmt.Fprintf(os.Stderr, "vectytemplater usage: %s output [-template=<name>]\n", os.Args[0])
	fmt.Fprintln(os.Stderr, "flags:")
	fmt.Fprintln(os.Stderr, "   template  template name. Available: [\"default\"]")
}
