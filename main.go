package main

import (
	"errors"
	"log"
	"os/user"

	"github.com/docopt/docopt-go"
	"github.com/stuart-xyz/harvest-cli/cli"
	"github.com/stuart-xyz/harvest-cli/services"
)

func main() {
	usage := `Harvest CLI.

    Usage:
      harvest view [--yesterday]
      harvest log <hours> [--yesterday]
      harvest log <issue_ref> <hours> [--yesterday]
      harvest -h | --help
      harvest --version

    Options:
      -h --help     Show this screen.
      --version     Show version.`

	opts, _ := docopt.ParseArgs(usage, nil, "1.0")

	err := executeCommand(opts)
	if err != nil {
		log.Fatal(err)
	}
}

func executeCommand(opts docopt.Opts) (err error) {
	currentUser, err := user.Current()
	if err != nil {
		return err
	}

	config, err := services.GetConfig(currentUser.HomeDir)
	if err != nil {
		return err
	}

	if isViewCommand, _ := opts.Bool("view"); isViewCommand {
		return cli.View(config, opts)
	} else if isLogCommand, _ := opts.Bool("log"); isLogCommand {
		return cli.Log(config, opts)
	}

	return errors.New("unrecognised command")
}
