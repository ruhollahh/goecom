package checkhandler

type DlvReadinessResp struct {
	Status string `json:"status"`
}

type DlvLivenessResp struct {
	Status     string `json:"status,omitempty"`
	Build      string `json:"build,omitempty"`
	Host       string `json:"host,omitempty"`
	GOMAXPROCS int    `json:"GOMAXPROCS,omitempty"`
}
