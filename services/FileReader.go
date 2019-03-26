package services

import (
	"encoding/csv"
	"harvest-cli/model"
	"io"
	"os"
	"strconv"
)

func buildTaskIndex() (taskIndex map[string]model.Task, taskIndexKeys []string, err error) {
	csvfile, err := os.Open("/Users/stuart/.jira.d/harvest-task-list.csv")
	if err != nil {
		return nil, nil, err
	}

	taskIndex = make(map[string]model.Task)
	reader := csv.NewReader(csvfile)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, err
		}

		projectId, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, nil, err
		}

		taskId, err := strconv.Atoi(record[1])
		if err != nil {
			return nil, nil, err
		}

		taskDescription := record[2]
		taskIndexKeys = append(taskIndexKeys, taskDescription)
		taskIndex[taskDescription] = model.Task{projectId, taskId, taskDescription}
	}

	return taskIndex, taskIndexKeys, nil
}
