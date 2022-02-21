package project

import "github.com/burntcarrot/quii/entity/project"

type Project struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Github      string `json:"github_url"`
}

func (p *Project) ToDomain() project.Domain {
	return project.Domain{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Github:      p.Github,
	}
}
