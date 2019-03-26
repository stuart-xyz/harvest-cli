package main

import (
	"fmt"
	"harvest-cli/model"
	"os"
	"os/exec"

	"github.com/docopt/docopt-go"
	"gopkg.in/yaml.v2"
)

func main() {
	usage := `Harvest CLI.

    Usage:
      harvest log <ticket_ref> <category> <time>...
      harvest -h | --help
      harvest --version

    Options:
      -h --help     Show this screen.
      --version     Show version.`

	opts, _ := docopt.ParseArgs(usage, nil, "0.1")

	err := executeCommand(opts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func executeCommand(opts docopt.Opts) (err error) {
	if isLog, _ := opts.Bool("log"); isLog {
		ticketRef, _ := opts.String("<ticket_ref>")
		ticketSummary, err := runInSystem("jira", []string{"view", ticketRef})
		if err != nil {
			return err
		}
		var jiraTicket model.JiraTicket
		err = yaml.Unmarshal(ticketSummary, &jiraTicket)
		if err != nil {
			return err
		}
		fmt.Println(jiraTicket)
	}
	return
}

func runInSystem(script string, args []string) (output []byte, err error) {
	osCommand := exec.Command(script, args...)
	output, err = osCommand.Output()
	if err != nil {
		return nil, err
	}
	return output, nil
}
