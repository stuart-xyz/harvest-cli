package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	var closestMatchingTasks []harvest.Task
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
	requestJson, err := json.Marshal(logTimeRequest)
	if err != nil {
		return -1, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.harvestapp.com/v2/time_entries", bytes.NewBuffer(requestJson))
	if err != nil {
		return -1, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.HarvestApiToken))
	req.Header.Add("Harvest-Account-Id", strconv.Itoa(config.HarvestAccountId))

	resp, err := client.Do(req)
	if err != nil {
		return -1, err
	}

	return resp.StatusCode, nil
}

func ViewLog(config Config, date string) (entries []harvest.LogEntry, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.harvestapp.com/v2/time_entries?from=%s&to=%s", date, date), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.HarvestApiToken))
	req.Header.Add("Harvest-Account-Id", strconv.Itoa(config.HarvestAccountId))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var viewLogResponse harvest.ViewLogResponse
	err = json.Unmarshal(body, &viewLogResponse)
	if err != nil {
		return nil, err
	}

	return viewLogResponse.Entries, nil
}
