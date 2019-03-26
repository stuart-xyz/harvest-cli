package services

import (
	"fmt"
	"harvest-cli/model"

	"github.com/schollz/closestmatch"
)

func FuzzyMatchTicketKeywords(ticket model.JiraTicket) (tasks []model.Task, err error) {
	taskIndex, taskIndexKeys, err := buildTaskIndex()
	if err != nil {
		return []model.Task{}, err
	}

	bagSizes := []int{2}
	closestMatchModel := closestmatch.New(taskIndexKeys, bagSizes)
	closestMatches := closestMatchModel.ClosestN(fmt.Sprintf("%s %s %s", ticket.Project, ticket.Summary, ticket.Labels), 3)

	closestMatchingTasks := []model.Task{}
	for _, key := range closestMatches {
		closestMatchingTasks = append(closestMatchingTasks, taskIndex[key])
	}

	return closestMatchingTasks, nil
}

// func logTime(apiToken string, userId int, timeBlock model.TimeBlock) (err error) {
// 	json, err := json.Marshal(timeBlock)
// 	if err != nil {
// 		return err
// 	}

// 	client := &http.Client{}
// 	req, err := http.NewRequest("POST", "https://api.harvestapp.com/v2/time_entries", bytes.NewBuffer(json))
// 	req.Header.Add("Authorization", fmt.Sprintf("Bearer: %s", apiToken))
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
