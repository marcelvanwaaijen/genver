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
	Commit  string
}

func init() {
	if len(os.Getenv("GOPACKAGE")) == 0 {
		pckg = DefaultPackage
	} else {
		pckg = os.Getenv("GOPACKAGE")
	}
	flag.StringVar(&output, "o", DefaultOutput, "output file name")
	flag.StringVar(&pckg, "p", pckg, "package name")
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.Parse()
	if showVersion {
		fmt.Println("\nversion:", version)
		os.Exit(0)
	}
}

func main() {
	var sb strings.Builder
	cmd := exec.Command("git", "describe", "--tags")
	cmd.Stderr = os.Stderr
	cmd.Stdout = &sb
	if err := cmd.Run(); err != nil {
		log.Fatalf("error: %v", err)
	}

	v, c := splitver(strings.TrimSuffix(sb.String(), "\n"))
	vi := verinfo{
		TS:      time.Now().Format("2006-01-02 15:04:05 MST"),
		Package: pckg,
		Version: v,
		Commit:  c,
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

func splitver(v string) (string, string) {
	parts := strings.Split(v, "-")
	if len(parts) > 2 {
		nv := strings.Join(parts[0:len(parts)-2], "-")
		c := v
		return nv, c
	}
	return v, v
}
