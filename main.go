package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"
)

const (
	DefaultPackage string = "main"
	DefaultOutput  string = "version.go"
)

var (
	output      string
	pckg        string
	showVersion bool
)

//go:embed version.template
var tmplt string

type verinfo struct {
	TS      string
	Package string
	Version string
}

func init() {
	flag.StringVar(&output, "o", DefaultOutput, "output file name")
	flag.StringVar(&pckg, "p", DefaultPackage, "package name")
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.Parse()
	if showVersion {
		fmt.Println("\nversion:", version)
		os.Exit(0)
	}
}

func main() {
	var sb strings.Builder
	cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")
	cmd.Stderr = os.Stderr
	cmd.Stdout = &sb
	if err := cmd.Run(); err != nil {
		log.Fatalf("error: %v", err)
	}

	vi := verinfo{
		TS:      time.Now().Format("2006-01-02 15:04:05 MST"),
		Package: pckg,
		Version: strings.TrimSuffix(sb.String(), "\n"),
	}

	t, err := template.New("version").Parse(tmplt)
	if err != nil {
		log.Fatalf("error parsing template: %v", err)
	}
	f, err := os.Create(output)
	if err != nil {
		log.Fatalf("error creating %s: %v", output, err)
	}
	defer f.Close()

	if err := t.Execute(f, vi); err != nil {
		log.Fatalf("error writing version info: %v", err)
	}
}
