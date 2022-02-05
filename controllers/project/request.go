package project

type CreateRequest struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Github      string `json:"github_url"`
}
