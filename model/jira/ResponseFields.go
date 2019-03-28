package jira

type ResponseFields struct {
	Project Project  `json:"project"`
	Summary string   `json:"summary"`
	Labels  []string `json:"labels"`
}
