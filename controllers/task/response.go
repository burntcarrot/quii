package task

type GetResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Deadline string `json:"deadline"`
	Status   string `json:"status"`
}

type CreateResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Deadline string `json:"deadline"`
	Status   string `json:"status"`
}
