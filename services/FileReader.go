package services

import (
	"encoding/csv"
	"fmt"
	"harvest-cli/model/harvest"
	"io"
	"io/ioutil"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

func GetConfig(homeDir string) (config Config, err error) {
	configFile, err := ioutil.ReadFile(fmt.Sprintf("%s/.jira.d/config.yml", homeDir))
	if err != nil {
		return Config{}, err
	}

	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func BuildTaskIndex(homeDir string) (taskIndex map[string]harvest.Task, taskIndexKeys []string, err error) {
	csvfile, err := os.Open(fmt.Sprintf("%s/.jira.d/harvest-task-list.csv", homeDir))
	if err != nil {
		return nil, nil, err
	}

	taskIndex = make(map[string]harvest.Task)
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
		taskIndex[taskDescription] = harvest.Task{projectId, taskId, taskDescription}
	}

	return taskIndex, taskIndexKeys, nil
}
