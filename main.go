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
      harvest log <issue_ref> <hours>
      harvest log <hours>
      harvest view [--yesterday]
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

	if isLog, _ := opts.Bool("log"); isLog {
		cli.Log(config, opts)
	} else if isView, _ := opts.Bool("view"); isView {
		cli.View(config, opts)
	}

	return nil
}
