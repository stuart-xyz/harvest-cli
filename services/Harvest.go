package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	jiramodel "github.com/stuart-xyz/harvest-cli/model/jira"

	"github.com/stuart-xyz/harvest-cli/model/harvest"

	"github.com/schollz/closestmatch"
)

func FuzzyMatchIssue(taskIndex map[string]harvest.Task, taskIndexKeys []string, issue jiramodel.Issue) (tasks []harvest.Task, err error) {
	bagSizes := []int{2}
	closestMatchModel := closestmatch.New(taskIndexKeys, bagSizes)
	closestMatches := closestMatchModel.ClosestN(fmt.Sprintf("%s %s %s", issue.ProjectKey, issue.Summary, issue.Labels), 3)

	closestMatchingTasks := []harvest.Task{}
	for _, key := range closestMatches {
		closestMatchingTasks = append(closestMatchingTasks, taskIndex[key])
	}

	return closestMatchingTasks, nil
}

func LogTime(config Config, task harvest.Task, timeBlock harvest.TimeBlock) (statusCode int, err error) {
	logTimeRequest := harvest.LogTimeRequest{
		ProjectId:   task.ProjectId,
		TaskId:      task.TaskId,
		Date:        timeBlock.Date,
		Hours:       timeBlock.Hours,
		Note:        timeBlock.Note,
		ExternalRef: timeBlock.ExternalRef,
	}
	json, err := json.Marshal(logTimeRequest)
	if err != nil {
		return -1, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.harvestapp.com/v2/time_entries", bytes.NewBuffer(json))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.HarvestApiToken))
	req.Header.Add("Harvest-Account-Id", strconv.Itoa(config.HarvestAccountId))

	resp, err := client.Do(req)
	if err != nil {
		return -1, err
	}

	return resp.StatusCode, nil
}
