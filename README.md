# tinygo-edit

Add an environment variable for tinygo and open the editor.  
Using tinygo-edit, you can easily integrate with gopls.  

## Description

### Vim

![tinygo-edit-with-vim](tinygo-edit-with-vim.gif)

If you are using Vim, you had better read the following.

* https://github.com/sago35/tinygo.vim

### VSCode

![tinygo-edit-with-code](tinygo-edit-with-code.gif)

If you are using VSCode, you had better read the following.

* https://marketplace.visualstudio.com/items?itemName=tinygo.vscode-tinygo

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
  -h, --help            Show context-sensitive help (also try --help-long and
                        --help-man).
      --editor="vim"    editor path
      --wait            wait for the editor to close
      --without-goroot  don't use proper GOROOT
      --target=TARGET   target name
      --version         Show application version.
```

Now you can use tinygo-edit.  
It works with or without go.mod, so you can work with gopls very simply.  

```
# Vim
$ tinygo-edit --target xiao --editor vim --wait

# gVim
$ tinygo-edit --target xiao --editor gvim

# VSCode
$ tinygo-edit --target xiao --editor code
```

## Usage (with TinyGo older than 0.15)

*deprecated : To be removed in 0.3.0*

If you want to use TinyGo older than 0.15, you can disable it with the following  

```
$ tinygo-edit --without-goroot --target xiao --editor code
```

If it doesn't work, please try the following  

1. Remove go.mod in the current dir
2. If $TINYGOPATH/go.mod exists, delete it.
3. Restart tinygo-edit.

If you don't want to remove the go.mod, try the following page  

* https://github.com/tinygo-org/tinygo-site/pull/107
  * https://deploy-preview-107--tinygo.netlify.app/ide-integration/

## Installation

To install, run:
```
go install github.com/sago35/tinygo-edit@latest
```
Be sure that you have added your GOBIN to the PATH.
You can find your GOBIN by running ```go env```.

### If GOBIN is empty
The command below should be added to your ```.bashrc``` or ```.zshrc```.
```
export PATH="$GOPATH/bin/:$PATH"
```
### If GOBIN is not empty
The command below should be added to your ```.bashrc``` or ```.zshrc```.
```
export PATH="$GOBIN:$PATH"
```
## Build

```
go build
```

### Environment

* go
* kingpin.v2
* tinygo 0.15

## FAQ

### I can't "tinygo build" in a vim opened by tinygo-edit.

You can use the `unset GOROOT` command, which will allow you to build.

## Author

sago35 - <sago35@gmail.com>
