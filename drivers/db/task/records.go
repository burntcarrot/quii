package task

import "github.com/burntcarrot/pm/entity/task"

type Task struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Deadline string `json:"deadline"`
	Status   string `json:"status"`
}

func (t *Task) ToDomain() task.Domain {
	return task.Domain{
		ID:       t.ID,
		Name:     t.Name,
		Type:     t.Type,
		Deadline: t.Deadline,
		Status:   t.Status,
	}
}
