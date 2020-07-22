package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/mattn/go-tty"
)

func edit(target, editor string) error {
	buf := bytes.Buffer{}
	cmd := exec.Command(`tinygo`, `info`, `-target`, target)
	cmd.Stdout = &buf
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}

	env := []string{`GOPATH=` + strings.Join([]string{os.Getenv(`TINYGOPATH`), os.Getenv(`GOPATH`)}, string(os.PathListSeparator))}
	scanner := bufio.NewScanner(&buf)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.SplitN(line, `:`, 2)
		if len(s) != 2 {
			continue
		}
		s[1] = strings.TrimLeft(s[1], ` `)

		switch s[0] {
		case `GOOS`, `GOARCH`:
			env = append(env, fmt.Sprintf(`%s=%s`, s[0], s[1]))
		case `build tags`:
			env = append(env, fmt.Sprintf(`GOFLAGS=-tags=%s`, strings.Join(strings.Split(s[1], ` `), `,`)))
		}
	}

	err = startEditor(editor, env)
	if err != nil {
		return err
	}

	return nil
}

func startEditor(editor string, env []string) error {
	tty, err := tty.Open()
	if err != nil {
		return err
	}
	defer tty.Close()
	cmd := exec.Command(editor)
	cmd.Stdin = tty.Input()
	cmd.Stdout = tty.Output()
	cmd.Stderr = tty.Output()
	cmd.Env = append(os.Environ(), env...)
	//return cmd.Start()
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("abort renames: %s", err)
	}
	return nil
}
