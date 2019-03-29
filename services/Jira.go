package services

import (
	"strings"

	jiramodel "github.com/stuart-xyz/harvest-cli/model/jira"

	"gopkg.in/Netflix-Skunkworks/go-jira.v1"
)

func GetIssue(httpClient jira.HttpClient, config Config, reference string) (issue jiramodel.Issue, err error) {
	jsonIssue, err := jira.GetIssue(httpClient, config.JiraEndpoint, reference, nil)
	if err != nil {
		return jiramodel.Issue{}, err
	}

	var labels []string
	rawLabels := jsonIssue.Fields["labels"].([]interface{})
	for _, rawLabel := range rawLabels {
		label := rawLabel.(string)
		labels = append(labels, label)
	}

	issue = jiramodel.Issue{
		Id:         jsonIssue.ID,
		ProjectId:  jsonIssue.Fields["project"].(map[string]interface{})["id"].(string),
		ProjectKey: jsonIssue.Fields["project"].(map[string]interface{})["key"].(string),
		Summary:    jsonIssue.Fields["summary"].(string),
		Labels:     strings.Join(labels, " "),
	}

	return issue, nil
}
