package harvest

type ViewLogResponse struct {
	Entries []LogEntry `json:"time_entries"`
}
