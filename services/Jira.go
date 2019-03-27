package services

import (
	"harvest-cli/model"

	"gopkg.in/yaml.v2"
)

func GetJiraTicket(reference string) (ticket model.JiraTicket, err error) {
	ticketSummary, err := runInSystem("jira", []string{"view", reference})
	if err != nil {
		return model.JiraTicket{}, err
	}

	err = yaml.Unmarshal(ticketSummary, &ticket)
	if err != nil {
		return model.JiraTicket{}, err
	}

	return ticket, nil
}
