package harvest

type LogTimeRequest struct {
	ProjectId   int               `json:"project_id"`
	TaskId      int               `json:"task_id"`
	Date        string            `json:"spent_date"`
	Hours       float64           `json:"hours"`
	Note        string            `json:"notes"`
	ExternalRef ExternalReference `json:"external_reference"`
}
