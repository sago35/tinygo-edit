package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	appName        = "tinygo-edit"
	appDescription = ""
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
)

// Run ...
func (c *cli) Run(args []string) error {
	app.UsageWriter(c.errStream)

	if os.Getenv(`TINYGOPATH`) == "" {
		return fmt.Errorf("$TINYGOPATH is not set. ex: export TINYGOPATH=/path/to/your/tinygo/")
	}

	targets, err := filepath.Glob(filepath.Join(os.Getenv(`TINYGOPATH`), "targets", "*.json"))
	if err != nil {
		return err
	}
	for i := range targets {
		targets[i] = strings.TrimSuffix(filepath.Base(targets[i]), filepath.Ext(targets[i]))
	}
	targets = append(targets, "pyportal")
	app.Flag("target", "target name").Default("pyportal").EnumVar(&target, targets...)

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

	switch k {
	default:
		if *editor == `vim` {
			*wait = true
		}

		err := edit(target, *editor, *wait)
		if err != nil {
			return err
		}
	}

	return nil
}
