package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"harvest-cli/model"
	"io/ioutil"
	"net/http"
	"strconv"

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

func LogTime(config model.Config, task model.Task, timeBlock model.TimeBlock) (statusCode int, err error) {
	externalRef := model.ExternalReference{
		Id:        "37857",
		GroupId:   "11800",
		Permalink: timeBlock.Url,
	}
	logTimeRequest := model.LogTimeRequest{
		ProjectId:   task.ProjectId,
		TaskId:      task.TaskId,
		Date:        timeBlock.Date,
		Hours:       timeBlock.Hours,
		Note:        timeBlock.Note,
		ExternalRef: externalRef,
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

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))

	return resp.StatusCode, nil
}
