package cli

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/coryb/oreo"
	"github.com/docopt/docopt-go"
	"github.com/stuart-xyz/harvest-cli/model/harvest"
	jiramodel "github.com/stuart-xyz/harvest-cli/model/jira"
	"github.com/stuart-xyz/harvest-cli/services"
)

func Log(config services.Config, opts docopt.Opts) (err error) {
	issueReference, err := opts.String("<issue_ref>")
	if err != nil {
		issueReference = ""
	}
	hours, err := opts.Float64("<hours>")
	if err != nil {
		return err
	}
	logForYesterday, err := opts.Bool("--yesterday")
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

	taskIndex, taskIndexKeys, err := services.BuildTaskIndex(*config.HomeDir)
	if err != nil {
		return err
	}

	var jiraIssue *jiramodel.Issue

	if issueReference != "" {
		jiraIssueResponse, err := services.GetIssue(httpClient, config, issueReference)
		jiraIssue = &jiraIssueResponse
		if err != nil {
			return err
		}
	}

	var jiraIssueToFuzzyMatch jiramodel.Issue
	if jiraIssue == nil {
		jiraIssueToFuzzyMatch = jiramodel.Issue{}
	} else {
		jiraIssueToFuzzyMatch = *jiraIssue
	}

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
			var projectKey string
			if jiraIssue != nil {
				projectKey = jiraIssue.ProjectKey
			}
			jiraIssueToFuzzyMatch = jiramodel.Issue{
				Id:         "",
				ProjectId:  "",
				ProjectKey: projectKey,
				Summary:    strippedInput,
				Labels:     "",
			}
		}
	}

	var note *string
	var externalRef *harvest.ExternalReference

	if jiraIssue != nil {
		noteString := fmt.Sprintf("%s: %s", issueReference, jiraIssue.Summary)
		note = &noteString
		externalRef = &harvest.ExternalReference{
			Id:      jiraIssue.Id,
			GroupId: jiraIssue.ProjectId,
			Permalink: fmt.Sprintf(
				"%s/secure/RapidBoard.jspa?rapidView=35&projectKey=%s&modal=detail&selectedIssue=%s",
				config.JiraEndpoint,
				jiraIssue.ProjectKey,
				issueReference,
			),
		}
	}

	var date string
	if logForYesterday {
		date = time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	} else {
		date = time.Now().Format("2006-01-02")
	}

	timeBlock := harvest.TimeBlock{
		Date:        date,
		Hours:       hours,
		Note:        note,
		ExternalRef: externalRef,
	}

	statusCode, err := services.LogTime(config, selectedTask, timeBlock)
	if err != nil {
		return err
	}

	fmt.Printf("Response status code %d\n", statusCode)

	return nil
}
