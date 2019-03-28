package main

import (
	"bufio"
	"fmt"
	"harvest-cli/model"
	"harvest-cli/services"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/docopt/docopt-go"
)

func main() {
	usage := `Harvest CLI.

    Usage:
      harvest log <ticket_ref> <hours>
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
		ticketReference, err := opts.String("<ticket_ref>")
		if err != nil {
			return err
		}

		hours, err := opts.Float64("<hours>")
		if err != nil {
			return err
		}

		jiraTicket, err := services.GetJiraTicket(ticketReference)
		if err != nil {
			return err
		}

		jiraTicketToFuzzyMatch := jiraTicket
		var selectedTask model.Task
		for {
			tasks, err := services.FuzzyMatchTicket(jiraTicketToFuzzyMatch)
			if err != nil {
				return err
			}

			fmt.Println("Enter the number of the correct task, or enter a string to search for other tasks")
			for index, task := range tasks {
				fmt.Printf("[%d] %s\n", index, task.Description)
			}

			consoleReader := bufio.NewReader(os.Stdin)
			fmt.Println()
			input, _ := consoleReader.ReadString('\n')
			strippedInput := strings.TrimSuffix(input, "\n")

			if index, err := strconv.Atoi(strippedInput); err == nil {
				selectedTask = tasks[index]
				break
			} else {
				jiraTicketToFuzzyMatch = model.JiraTicket{
					Project: jiraTicket.Project,
					Summary: strippedInput,
					Labels:  "",
				}
			}
		}

		config, err := services.GetConfig()
		if err != nil {
			return err
		}

		timeBlock := model.TimeBlock{
			Date:  time.Now().Format("2006-01-02"),
			Hours: hours,
			Note:  fmt.Sprintf("%s: %s", ticketReference, jiraTicket.Summary),
			Url:   fmt.Sprintf("%s/secure/RapidBoard.jspa?rapidView=35&projectKey=DEV&modal=detail&selectedIssue=%s", config.JiraEndpoint, ticketReference),
		}

		statusCode, err := services.LogTime(config, selectedTask, timeBlock)
		if err != nil {
			return err
		}

		fmt.Printf("Response status code %d\n", statusCode)
	}

	return
}
