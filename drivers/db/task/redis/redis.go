package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	dbTask "github.com/burntcarrot/quii/drivers/db/task"
	"github.com/burntcarrot/quii/entity/task"
	"github.com/burntcarrot/quii/errors"
	"github.com/go-redis/redis/v8"
)

const MAX_FETCH_ROWS = 9 * 100000

type TaskRepo struct {
	Conn *redis.Client
}

func NewTaskRepo(conn *redis.Client) task.DomainRepo {
	return &TaskRepo{Conn: conn}
}

func (p *TaskRepo) CreateTask(ctx context.Context, us task.Domain) (task.Domain, error) {
	// get ID from task counter
	counter := fmt.Sprintf("%s:projects:%s:tasks:counter", us.Username, strings.ToLower(us.ProjectName))
	counterValue, counterErr := p.Conn.Get(ctx, counter).Result()
	if counterErr != nil {
		return task.Domain{}, errors.ErrInternalServerError
	}

	// generate task ID
	task_id := "task_" + counterValue

	createdTask := dbTask.Task{
		ID:       task_id,
		Name:     us.Name,
		Type:     us.Type,
		Deadline: us.Deadline,
		Status:   us.Status,
	}

	raw, err := json.Marshal(createdTask)
	if err != nil {
		return task.Domain{}, errors.ErrInternalServerError
	}

	// new: list-based
	key := fmt.Sprintf("%s:projects:%s:tasks", us.Username, strings.ToLower(us.ProjectName))

	insertErr := p.Conn.RPush(ctx, key, raw).Err()
	if insertErr != nil {
		return task.Domain{}, errors.ErrInternalServerError
	}

	// increment counter after creating task
	incrErr := p.Conn.Incr(ctx, counter).Err()
	if incrErr != nil {
		return task.Domain{}, errors.ErrInternalServerError
	}

	return createdTask.ToDomain(), nil
}

func (p *TaskRepo) GetTasks(ctx context.Context, username, projectName string) ([]task.Domain, error) {
	key := fmt.Sprintf("%s:projects:%s:tasks", username, strings.ToLower(projectName))
	raw, err := p.Conn.LRange(ctx, key, 0, MAX_FETCH_ROWS).Result()
	if err != nil {
		return []task.Domain{}, errors.ErrInternalServerError
	}

	ts := new(dbTask.Task)
	var tasks []task.Domain

	for _, j := range raw {
		if err := json.Unmarshal([]byte(j), ts); err != nil {
			return []task.Domain{}, errors.ErrInternalServerError
		}

		tasks = append(tasks, ts.ToDomain())
	}

	return tasks, nil
}

func (p *TaskRepo) GetTaskByID(ctx context.Context, username, projectName, taskID string) ([]task.Domain, error) {
	key := fmt.Sprintf("%s:projects:%s:tasks", username, strings.ToLower(projectName))
	raw, err := p.Conn.LRange(ctx, key, 0, MAX_FETCH_ROWS).Result()
	if err != nil {
		return []task.Domain{}, errors.ErrInternalServerError
	}

	ts := new(dbTask.Task)
	var tasks []task.Domain

	for _, j := range raw {
		if err := json.Unmarshal([]byte(j), ts); err != nil {
			return []task.Domain{}, errors.ErrInternalServerError
		}

		if strings.EqualFold(ts.ID, taskID) {
			tasks = append(tasks, ts.ToDomain())
			return tasks, nil
		}
	}

	return []task.Domain{}, nil
}

func (p *TaskRepo) GetTaskByName(ctx context.Context, username, projectName, taskName string) ([]task.Domain, error) {
	key := fmt.Sprintf("%s:projects:%s:tasks", username, strings.ToLower(projectName))
	raw, err := p.Conn.LRange(ctx, key, 0, MAX_FETCH_ROWS).Result()
	if err != nil {
		return []task.Domain{}, errors.ErrInternalServerError
	}

	ts := new(dbTask.Task)
	var tasks []task.Domain

	for _, j := range raw {
		if err := json.Unmarshal([]byte(j), ts); err != nil {
			return []task.Domain{}, errors.ErrInternalServerError
		}

		if strings.EqualFold(ts.Name, taskName) {
			tasks = append(tasks, ts.ToDomain())
			return tasks, nil
		}
	}

	return []task.Domain{}, nil
}
