package jira

type Response struct {
	Id     string         `json:"id"`
	Fields ResponseFields `json:"fields"`
}
