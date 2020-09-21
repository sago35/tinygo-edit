package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	appName        = "tinygo-edit"
	appDescription = ``
)

type cli struct {
	outStream io.Writer
	errStream io.Writer
}

var (
	app      = kingpin.New(appName, appDescription)
	target   string
	editor   = app.Flag("editor", "editor path").Default("vim").String()
	wait     = app.Flag("wait", "wait for the editor to close").Bool()
	woGoroot = app.Flag("without-goroot", "don't use proper GOROOT").Bool()
)

// Run ...
func (c *cli) Run(args []string) error {
	app.UsageWriter(c.errStream)

	targets, err := getTargetsFromTinygoTargets()
	if err != nil {
		tinygopath, err := getTinygoPath()
		if err != nil {
			return err
		}
		targets, err = getTargets(tinygopath)
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

		if *woGoroot {
			// deprecated : To be removed in 0.3.0
			err := edit(target, *editor, *wait)
			if err != nil {
				return err
			}
			return nil
		}

		err := editWithGOROOT(target, *editor, *wait)
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
	matches, err := filepath.Glob(filepath.Join(tinygopath, `targets`, `*.json`))
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

func getTinygoPath() (string, error) {
	buf := new(bytes.Buffer)
	cmd := exec.Command("tinygo", "env", "TINYGOROOT")
	cmd.Stdout = buf
	cmd.Stderr = buf

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(buf.String()), nil
}
