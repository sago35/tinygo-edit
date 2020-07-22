package main

import (
	"fmt"
	"io"
	"os"

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
	target = app.Flag("target", "target name").Default("pyportal").Enum("pyportal", "feather-m4", "wioterminal", "xiao", "itsybitsy-nrf52840")
	editor = app.Flag("editor", "editor path").Default("vim").String()
	wait   = app.Flag("wait", "wait for the editor to close").Bool()
)

// Run ...
func (c *cli) Run(args []string) error {
	app.UsageWriter(c.errStream)

	if VERSION != "" {
		app.Version(fmt.Sprintf("%s version %s build %s", appName, VERSION, BUILDDATE))
	} else {
		app.Version(fmt.Sprintf("%s version - build -", appName))
	}
	app.HelpFlag.Short('h')

	if os.Getenv(`TINYGOPATH`) == "" {
		return fmt.Errorf("$TINYGOPATH is not set. ex: export TINYGOPATH=/path/to/your/tinygo/")
	}

	k, err := app.Parse(args[1:])
	if err != nil {
		return err
	}

	switch k {
	default:
		err := edit(*target, *editor, *wait)
		if err != nil {
			return err
		}
	}

	return nil
}
