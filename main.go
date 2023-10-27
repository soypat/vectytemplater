package main

import (
	"embed"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/soypat/rebed"
)

var (
	//go:embed _templates/default
	defaultFS embed.FS

	//go:embed _templates/websocket-cli
	websocketFS embed.FS
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "error in run:", err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stdout, "finished succesfully")
}

func run(args []string) (err error) {
	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" || args[0] == "-help" {
		usage("")
		return errors.New("showing help for vecty templater")
	}
	template := flag.String("template", "default", "template name. \"default\" or \"websocket\"")
	err = flag.CommandLine.Parse(args[1:])
	if err != nil {
		return err
	}
	workDir, err := os.Getwd()
	if err != nil {
		return err
	}
	dir := args[0]
	targetDir := filepath.Join(workDir, dir)
	if _, err := os.Stat(targetDir); err == nil {
		usage("")
		return errors.New("cannot create directory \"" + dir + "\" as it already exists")
	}

	tempDir := os.TempDir()
	if tempDir == "" {
		return errors.New("unable to create temporary directory")
	}
	err = os.Chdir(tempDir)
	if err != nil {
		return err
	}

	const perm os.FileMode = 0777
	rebed.FolderMode = perm
	// removes existing template write if present.
	if err = os.RemoveAll(filepath.Join(tempDir, "_templates")); err != nil {
		return err
	}
	switch *template {
	case "default":
		err = rebed.Create(defaultFS, ".")
	case "websocket-cli":
		err = rebed.Create(websocketFS, ".")
	default:
		err = errors.New("template " + *template + " not found")
	}
	if err != nil {
		return errors.New("creating temporary project structure: " + err.Error())
	}
	tmpOutput := filepath.Join(tempDir, "_templates", *template)
	fp, err := os.Create(filepath.Join(tmpOutput, "go.mod"))
	if err != nil {
		return err
	}
	fp.WriteString(gomod)
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Env = append(os.Environ(), "GOOS=js", "GOARCH=wasm", "GO111MODULE=on")
	cmd.Dir = tmpOutput
	fmt.Fprintln(os.Stdout, "running `go mod tidy`")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New("project created but failed to initialize module: " + err.Error() + "\n" + string(out))
	}
	fmt.Fprint(os.Stdout, string(out))
	err = os.Rename(tmpOutput, targetDir)
	if err != nil {
		switch runtime.GOOS {
		case "windows":
			err = exec.Command("xcopy", tmpOutput, targetDir, "/E", "/I", "/Y").Run()
		case "linux", "darwin":
			err = exec.Command("cp", "-r", tmpOutput, targetDir).Run()
		}
		return err
	}
	return nil
}

func usage(command string) {
	fmt.Fprintf(os.Stderr, "vectytemplater usage: %s <output> [-template=<name>]\n", os.Args[0])
	fmt.Fprintln(os.Stderr, "flags:")
	fmt.Fprintln(os.Stderr, "    template  template name. Available: [\"default\", \"websocket\"]")
}

const gomod = `module vecty-templater-project
` // No version info so `go mod tidy` sets own version
