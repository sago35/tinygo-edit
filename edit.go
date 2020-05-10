package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
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

	env := []string{}
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
	cmd := exec.Command(editor)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), env...)
	return cmd.Start()
}
