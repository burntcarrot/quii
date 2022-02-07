package task

type CreateRequest struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	ProjectName string `json:"project_name"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Deadline    string `json:"deadline"`
	Status      string `json:"status"`
}
