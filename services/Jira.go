package services

import (
	"encoding/json"
	"harvest-cli/model/jira"
	"strings"
)

func GetJiraTicket(reference string) (ticket jira.Ticket, err error) {
	jsonResponse, err := runInSystem("jira", []string{"view", reference, "-t json"})
	if err != nil {
		return jira.Ticket{}, err
	}

	var response jira.Response
	err = json.Unmarshal(jsonResponse, &response)
	if err != nil {
		return jira.Ticket{}, err
	}

	ticket = jira.Ticket{
		Project: response.Fields.Project.Key,
		Summary: response.Fields.Summary,
		Labels:  strings.Join(response.Fields.Labels, ","),
	}

	return ticket, nil
}
