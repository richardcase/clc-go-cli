# CenturyLink CLI

Command Line Interface for manipulating the CenturyLink IaaS.

## Getting Started

Download a binary compiled for your platform from the [releases page](https://github.com/CenturyLinkCloud/clc-go-cli/releases).

## Autocomplete

### Bash

Release tarballs for Linux/Unix/Darwin (starting from the release `2015-08-18`) contain 2 files for enabling autocomplete: `bash_autocomplete` and `install_autocompletion`. Execute `source bash_autocomplete` to turn on autocomplete for the current terminal session. `install_autocompletion` is provided for you to install autocomplete user-wide. The script, upon invoking,
copies the `bash_autocomplete` contents into `~/.bash_completion/clc` and updates `~/.bashrc` accordingly.

### PowerShell

Only v3 support is provided because previous versions do not support custom autocomplete handlers. PowerShell v3 is distributed as a part of Windows Management Framework 3.0, which can be downloaded [from here](http://www.microsoft.com/en-us/download/details.aspx?id=34595). You can check the version by typing `$PSVersionTable.PSVersion`.

To turn on autocomplete execute `.\powershell3_autocomplete.ps1`. You can find the file in the release tarball for Windows.

## Getting Help

Explore the available resources, commands, options and other useful guidance using the `--help` option:
`clc --help`, `clc <resource> --help` and `clc <resouce> <command> --help` are all at your service.

The documentation of the underlying HTTP API can be found [here](https://www.ctl.io/api-docs/v2/).

## The Development Process

* [Install Go](https://golang.org/).
* Install Godep: `go get github.com/tools/godep`.
* Clone this repo (you do not have to use `go get`).
* [Ensure your $GOPATH is set correctly](http://golang.org/cmd/go/#hdr-GOPATH_environment_variable).
* Install dependencies with Godep: enter the repo's root and `godep restore`.
* Use the dev script to run commands: `./dev <resource> <command>`.
* Install go vet: `go get code.google.com/p/go.tools/cmd/vet`.
* Before commit check that `gofmt -d=true ./..` and `go vet ./...` do not produce any output (except for that coming from `Godeps/_workspace` - ignore it) and check that all tests pass via `./run_tests`.

If you want to make an executable, simply run `./build`. The binary will appear in the `./out` folder.

If you want bash autocomplete to work, you have to use the binary generated by the `./build` command. To turn autocomplete on either source `autocomplete/bash_autocomplete` or install it by running `./install_autocompletion`.
