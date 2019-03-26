package model

type TimeBlock struct {
	ProjectId   int    `json:"project_id"`
	TaskId      int    `json:"task_id"`
	Date        string `json:"spent_date"`
	Hours       int    `json:"hours"`
	Note        string `json:"note"`
	ExternalRef string `json:"external_reference"`
}
