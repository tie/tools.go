package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"

	"go/parser"
	"go/token"
)

var (
	filePath = flag.String("tools", "tools.go", "path to tools import file")
	verbose  = flag.Bool("verbose", true, "verbose output")
)

func main() {
	log.SetFlags(0)
	flag.Parse()

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, *filePath, nil, parser.ImportsOnly)
	if err != nil {
		log.Fatalf("parse %q: %s", *filePath, err)
		return
	}

	packages := []string{}
	for _, spec := range f.Imports {
		if spec.Path.Kind != token.STRING {
			continue
		}
		var name string
		if spec.Name != nil {
			name = spec.Name.Name
		}
		if name != "_" {
			continue
		}
		path, err := strconv.Unquote(spec.Path.Value)
		if err != nil {
			continue
		}
		packages = append(packages, path)
	}

	if len(packages) <= 0 {
		log.Fatalf("no tools found in %q", *filePath)
		return
	}

	flagArgs := flag.Args()

	prog := "go"
	args := []string{"install"}
	args = append(args, flagArgs...)
	args = append(args, packages...)

	if *verbose {
		shell := make([]interface{}, len(args)+1)
		shell[0] = prog
		for i := range args {
			shell[i+1] = args[i]
		}
		fmt.Fprintln(os.Stderr, shell...)
	}

	cmd := exec.Command(prog, args...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	if *verbose {
		cmd.Stdout = os.Stdout
	}
	if err := cmd.Run(); err != nil {
		log.Fatalf("%s: %s", prog, err)
		return
	}
}
