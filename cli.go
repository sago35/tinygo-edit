package main

import (
	"bufio"
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

	targets, err := getTargets(os.Getenv(`TINYGOPATH`))
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

func getTargets(tinygopath string) ([]string, error) {

	r, err := os.Open(filepath.Join(os.Getenv(`TINYGOPATH`), "Makefile"))
	if err != nil {
		return nil, err
	}
	defer r.Close()

	scanner := bufio.NewScanner(r)
	targets := []string{}
	exists := map[string]bool{}
	for scanner.Scan() {
		t := scanner.Text()
		if strings.Contains(t, `$(TINYGO)`) && strings.Contains(t, `-target=`) {
			fields := strings.Fields(t)
			for _, f := range fields {
				if strings.HasPrefix(f, `-target=`) {
					target := f[8:]
					if !exists[target] {
						exists[target] = true
						targets = append(targets, target)
					}
				}
			}
		}
	}

	return targets, nil
}
