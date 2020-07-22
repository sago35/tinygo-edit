package main

import (
	"fmt"
	"io"

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

	k, err := app.Parse(args[1:])
	if err != nil {
		return err
	}

	switch k {
	default:
		err := edit(*target, *editor)
		if err != nil {
			return err
		}
	}

	return nil
}
