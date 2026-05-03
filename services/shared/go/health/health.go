package health

type Snapshot struct {
	Service string            `json:"service"`
	Status  string            `json:"status"`
	Checks  map[string]string `json:"checks"`
}
