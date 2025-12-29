package task

type Task struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Payload    string `json:"payload"`
	Status     string `json:"status"`
	Retries    int    `json:"retries"`
	MaxRetries int    `json:"max_retries"`
}
