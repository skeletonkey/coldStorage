package queue

type queue struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Queues   topics `json:"queues"`
}

type topics map[string]topic
type topic struct {
	Durable          bool                   `json:"durable"`
	DeleteWhenUnused bool                   `json:"delete_when_unused"`
	Exclusive        bool                   `json:"exclusive"`
	NoWait           bool                   `json:"no_wait"`
	Arguments        map[string]interface{} `json:"arguments"`
}
