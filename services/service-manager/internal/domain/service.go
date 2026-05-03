package domain

type ServiceStatus struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Status  string `json:"status"`
	Profile string `json:"profile"`
}
