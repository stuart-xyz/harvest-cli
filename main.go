package main

import (
	"log"
	"os"
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

	opts, _ := docopt.ParseArgs(usage, nil, "0.1")

	err := executeCommand(opts)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func executeCommand(opts docopt.Opts) (err error) {
	user, err := user.Current()
	if err != nil {
		return err
	}

	config, err := services.GetConfig(user.HomeDir)
	if err != nil {
		return err
	}

	if isViewCommand, _ := opts.Bool("view"); isViewCommand {
		cli.View(config, opts)
	} else if isLogCommand, _ := opts.Bool("log"); isLogCommand {
		cli.Log(config, opts)
	}

	return nil
}
