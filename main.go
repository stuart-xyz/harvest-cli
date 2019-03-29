package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"
	"strconv"
	"strings"
	"time"

	"github.com/stuart-xyz/harvest-cli/services"

	"github.com/stuart-xyz/harvest-cli/model/harvest"

	jiramodel "github.com/stuart-xyz/harvest-cli/model/jira"

	"github.com/coryb/oreo"
	"github.com/docopt/docopt-go"
)

func main() {
	usage := `Harvest CLI.

    Usage:
      harvest log <issue_ref> <hours>
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
	if isLog, _ := opts.Bool("log"); isLog {
		issueReference, err := opts.String("<issue_ref>")
		if err != nil {
			return err
		}
		hours, err := opts.Float64("<hours>")
		if err != nil {
			return err
		}

		user, err := user.Current()
		if err != nil {
			return err
		}

		config, err := services.GetConfig(user.HomeDir)
		if err != nil {
			return err
		}

		httpClient := oreo.New().WithPreCallback(
			func(req *http.Request) (*http.Request, error) {
				// need to set basic auth header with user@domain:api-token
				authHeader := fmt.Sprintf(
					"Basic %s",
					base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", config.JiraEmail, config.JiraApiToken))),
				)
				req.Header.Add("Authorization", authHeader)
				return req, nil
			},
		)

		taskIndex, taskIndexKeys, err := services.BuildTaskIndex(user.HomeDir)
		if err != nil {
			return err
		}

		jiraIssue, err := services.GetIssue(httpClient, config, issueReference)
		if err != nil {
			return err
		}

		jiraIssueToFuzzyMatch := jiraIssue
		var selectedTask harvest.Task
		for {
			tasks, err := services.FuzzyMatchIssue(taskIndex, taskIndexKeys, jiraIssueToFuzzyMatch)
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
				jiraIssueToFuzzyMatch = jiramodel.Issue{
					Id:         "",
					ProjectId:  "",
					ProjectKey: jiraIssue.ProjectKey,
					Summary:    strippedInput,
					Labels:     "",
				}
			}
		}

		timeBlock := harvest.TimeBlock{
			Date:  time.Now().Format("2006-01-02"),
			Hours: hours,
			Note:  fmt.Sprintf("%s: %s", issueReference, jiraIssue.Summary),
			ExternalRef: harvest.ExternalReference{
				Id:      jiraIssue.Id,
				GroupId: jiraIssue.ProjectId,
				Permalink: fmt.Sprintf(
					"%s/secure/RapidBoard.jspa?rapidView=35&projectKey=%s&modal=detail&selectedIssue=%s",
					config.JiraEndpoint,
					jiraIssue.ProjectKey,
					issueReference,
				),
			},
		}

		statusCode, err := services.LogTime(config, selectedTask, timeBlock)
		if err != nil {
			return err
		}

		fmt.Printf("Response status code %d\n", statusCode)
	}

	return nil
}
