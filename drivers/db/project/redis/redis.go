package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	dbProject "github.com/burntcarrot/pm/drivers/db/project"
	"github.com/burntcarrot/pm/entity/project"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

const MAX_FETCH_ROWS = 9 * 100000

type ProjectRepo struct {
	Conn *redis.Client
}

func NewProjectRepo(conn *redis.Client) project.DomainRepo {
	return &ProjectRepo{Conn: conn}
}

func (p *ProjectRepo) CreateProject(ctx context.Context, us project.Domain) (project.Domain, error) {
	createdProject := dbProject.Project{
		ID:          uuid.New().String(),
		Name:        us.Name,
		Description: us.Description,
		Github:      us.Github,
	}

	raw, err := json.Marshal(createdProject)
	if err != nil {
		return project.Domain{}, err
	}

	// old: key hierarchy
	// key := fmt.Sprintf("%s:projects:%s", us.Username, createdProject.Name)

	// new: list-based
	key := fmt.Sprintf("%s:projects", us.Username)

	insertErr := p.Conn.RPush(ctx, key, raw).Err()
	if insertErr != nil {
		return project.Domain{}, insertErr
	}

	return createdProject.ToDomain(), nil
}

func (p *ProjectRepo) GetProjects(ctx context.Context, username string) ([]project.Domain, error) {
	key := fmt.Sprintf("%s:projects", username)
	raw, err := p.Conn.LRange(ctx, key, 0, MAX_FETCH_ROWS).Result()
	// TODO: remove print statements
	fmt.Println("Lrange raw:", raw, "\n")
	fmt.Println("Lrange error:", err)
	if err != nil {
		return []project.Domain{}, err
	}

	pr := new(dbProject.Project)
	var projects []project.Domain

	for _, j := range raw {
		if err := json.Unmarshal([]byte(j), pr); err != nil {
			return []project.Domain{}, err
		}

		projects = append(projects, pr.ToDomain())
	}

	return projects, nil
}

func (p *ProjectRepo) GetProjectByName(ctx context.Context, username, projectName string) ([]project.Domain, error) {
	key := fmt.Sprintf("%s:projects", username)
	raw, err := p.Conn.LRange(ctx, key, 0, MAX_FETCH_ROWS).Result()
	// TODO: remove print statements
	fmt.Println("Lrange raw:", raw, "\n")
	fmt.Println("Lrange error:", err)
	if err != nil {
		return []project.Domain{}, err
	}

	pr := new(dbProject.Project)
	var projects []project.Domain

	for _, j := range raw {
		if err := json.Unmarshal([]byte(j), pr); err != nil {
			return []project.Domain{}, err
		}

		if strings.ToLower(pr.Name) == strings.ToLower(projectName) {
			projects = append(projects, pr.ToDomain())
			return projects, nil
		}
	}

	return []project.Domain{}, nil
}
