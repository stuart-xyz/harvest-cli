package harvest

type LogEntry struct {
	Project ProjectResponse `json:"project"`
	Task    TaskResponse    `json:"task"`
	Hours   float64         `json:"hours"`
	Note    string          `json:"notes"`
}
