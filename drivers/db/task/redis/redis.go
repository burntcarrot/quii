package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	dbTask "github.com/burntcarrot/pm/drivers/db/task"
	"github.com/burntcarrot/pm/entity/task"
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
	counter := fmt.Sprintf("%s:projects:%s:tasks:counter", us.Username, strings.ToLower(us.ProjectName))

	counterValue, counterErr := p.Conn.Get(ctx, counter).Result()
	if counterErr != nil {
		return task.Domain{}, counterErr
	}

	task_id := "task_" + counterValue

	// TODO: do NOT use UUID for tasks, instead use a incremental ID
	createdTask := dbTask.Task{
		ID:       task_id,
		Name:     us.Name,
		Type:     us.Type,
		Deadline: us.Deadline,
		Status:   us.Status,
	}

	raw, err := json.Marshal(createdTask)
	if err != nil {
		return task.Domain{}, err
	}

	// old: key hierarchy
	// key := fmt.Sprintf("%s:projects:%s", us.Username, createdTask.Name)

	// new: list-based
	key := fmt.Sprintf("%s:projects:%s:tasks", us.Username, strings.ToLower(us.ProjectName))

	insertErr := p.Conn.RPush(ctx, key, raw).Err()
	if insertErr != nil {
		return task.Domain{}, insertErr
	}

	// increment counter after creating task
	incrErr := p.Conn.Incr(ctx, counter).Err()
	if incrErr != nil {
		return task.Domain{}, incrErr
	}

	return createdTask.ToDomain(), nil
}

func (p *TaskRepo) GetTasks(ctx context.Context, username, projectName string) ([]task.Domain, error) {
	key := fmt.Sprintf("%s:projects:%s:tasks", username, projectName)
	raw, err := p.Conn.LRange(ctx, key, 0, MAX_FETCH_ROWS).Result()
	// TODO: remove print statements
	fmt.Println("Lrange raw:", raw, "\n")
	fmt.Println("Lrange error:", err)
	if err != nil {
		return []task.Domain{}, err
	}

	ts := new(dbTask.Task)
	var tasks []task.Domain

	for _, j := range raw {
		if err := json.Unmarshal([]byte(j), ts); err != nil {
			return []task.Domain{}, err
		}

		tasks = append(tasks, ts.ToDomain())
	}

	return tasks, nil
}

func (p *TaskRepo) GetTaskByName(ctx context.Context, username, projectName, taskName string) ([]task.Domain, error) {
	key := fmt.Sprintf("%s:projects:%s:tasks", username, projectName)
	raw, err := p.Conn.LRange(ctx, key, 0, MAX_FETCH_ROWS).Result()
	// TODO: remove print statements
	fmt.Println("Lrange raw:", raw, "\n")
	fmt.Println("Lrange error:", err)
	if err != nil {
		return []task.Domain{}, err
	}

	ts := new(dbTask.Task)
	var tasks []task.Domain

	for _, j := range raw {
		if err := json.Unmarshal([]byte(j), ts); err != nil {
			return []task.Domain{}, err
		}

		if strings.ToLower(ts.ID) == strings.ToLower(taskName) {
			tasks = append(tasks, ts.ToDomain())
			return tasks, nil
		}
	}

	return []task.Domain{}, nil
}
