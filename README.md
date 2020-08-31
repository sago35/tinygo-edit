# tinygo-edit

Add an environment variable for tinygo and open the editor.  
Using tinygo-edit, you can easily integrate with gopls.  

## Description

### Vim

![tinygo-edit-with-vim](tinygo-edit-with-vim.gif)

### VSCode

![tinygo-edit-with-code](tinygo-edit-with-code.gif)

* https://dev.to/sago35/tinygo-vim-gopls-48h1

### Bash/ZSH Shell Completion

By default, all flags and commands/subcommands generate completions internally.  
You can enable autocompletion by setting the following to `~/.bashrc` etc.  

```
$ eval "$(tinygo-edit --completion-script-bash)"
```

Or for ZSH

```
$ eval "$(tinygo-edit --completion-script-zsh)"
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

Now you can use tinygo-edit.

```
# Vim
$ tinygo-edit --target xiao --editor vim --wait

# gVim
$ tinygo-edit --target xiao --editor gvim

# VSCode
$ tinygo-edit --target xiao --editor code
```

If it doesn't work, please try the following  

1. Remove go.mod in the current dir
2. If $TINYGOPATH/go.mod exists, delete it.
3. Restart tinygo-edit.

If you don't want to remove the go.mod, try the following page  

* https://github.com/tinygo-org/tinygo-site/pull/107
  * https://deploy-preview-107--tinygo.netlify.app/ide-integration/


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
