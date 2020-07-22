# tinygo-edit

Add an environment variable for tinygo and open the editor.

## Description

* https://dev.to/sago35/tinygo-vim-gopls-48h1

### Bash/ZSH Shell Completion

By default, all flags and commands/subcommands generate completions internally.  
You can enable autocompletion by setting the following to `~/.bashrc` etc.  

```
$ eval "$(your-cli-tool --completion-script-bash)"
```

Or for ZSH

```
$ eval "$(your-cli-tool --completion-script-zsh)"
```

* https://github.com/alecthomas/kingpin#bashzsh-shell-completion

## Usage

```
usage: tinygo-edit [<flags>]

Flags:
  -h, --help             Show context-sensitive help (also try --help-long and
                         --help-man).
      --target=pyportal  target name
      --editor="vim"     editor path
      --wait             wait for the editor to close
      --version          Show application version.
```

This program uses $TINYGOPATH, so set it up.  

```
# bash
$ export TINYGOPATH=/path/to/your/tinygo

$ windows cmd.exe
$ set TINYGOPATH=C:\path\to\your\tinygo
```

## Installation

```
$ go get github.com/sago35/tinygo-edit
```

## Build

```
$ go build
```

### Environment

* go
* kingpin.v2

## Notice

## Author

sago35 - <sago35@gmail.com>
