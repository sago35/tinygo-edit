package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	appName        = "tinygo-edit"
	appDescription = `This program uses $TINYGOPATH
  export TINYGOPATH=/path/to/tinygo-root-directory

You can use the following environment variables
  To get a list of targets from the result of 'tinygo targets':
    export TINYGO_EDIT_WITH_GOROOT=1

  Using the GOROOT environment variable to link with gopls:
    export TINYGO_EDIT_WITH_TINYGO_TARGETS=1
`
)

type cli struct {
	outStream io.Writer
	errStream io.Writer
}

var (
	app    = kingpin.New(appName, appDescription)
	target string
	editor = app.Flag("editor", "editor path").Default("vim").String()
	wait   = app.Flag("wait", "wait for the editor to close").Bool()
	goroot = app.Flag("with-goroot", "use proper GOROOT").Envar("TINYGO_EDIT_WITH_GOROOT").Default("0").Bool()
)

// Run ...
func (c *cli) Run(args []string) error {
	app.UsageWriter(c.errStream)

	targets := []string{}
	var err error

	if _, exists := os.LookupEnv("TINYGO_EDIT_WITH_TINYGO_TARGETS"); exists {
		targets, err = getTargetsFromTinygoTargets()
	} else {
		if os.Getenv(`TINYGOPATH`) == "" {
			return fmt.Errorf("$TINYGOPATH is not set. ex: export TINYGOPATH=/path/to/your/tinygo/")
		}
		targets, err = getTargets(os.Getenv(`TINYGOPATH`))
	}
	if err != nil {
		return err
	}
	app.Flag("target", "target name").EnumVar(&target, targets...)

	if VERSION != "" {
		app.Version(fmt.Sprintf("%s version %s build %s", appName, VERSION, BUILDDATE))
	} else {
		app.Version(fmt.Sprintf("%s version - build -", appName))
	}
	app.HelpFlag.Short('h')

	k, err := app.Parse(args[1:])
	if err != nil {
		return err
	}

	if target == "" {
		app.Usage(args[1:])
		return nil
	}

	switch k {
	default:
		if *editor == `vim` {
			*wait = true
		}

		if *goroot {
			err := editWithGOROOT(target, *editor, *wait)
			if err != nil {
				return err
			}
			return nil
		}

		err := edit(target, *editor, *wait)
		if err != nil {
			return err
		}
	}

	return nil
}

func getTargets(tinygopath string) ([]string, error) {
	return getTargetsFromJson(tinygopath)
}

func getTargetsFromJson(tinygopath string) ([]string, error) {
	// read from $TINYGOPATH/targets/*.json
	matches, err := filepath.Glob(filepath.Join(os.Getenv(`TINYGOPATH`), `targets`, `*.json`))
	if err != nil {
		return nil, err
	}
	for i := range matches {
		matches[i] = strings.TrimSuffix(filepath.Base(matches[i]), filepath.Ext(matches[i]))
	}

	return matches, err
}

func getTargetsFromTinygoTargets() ([]string, error) {
	buf := new(bytes.Buffer)
	cmd := exec.Command("tinygo", "targets")
	cmd.Stdout = buf
	cmd.Stderr = buf

	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	targets := []string{}
	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		targets = append(targets, scanner.Text())
	}

	return targets, nil
}
