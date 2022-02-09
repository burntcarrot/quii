package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	dbProject "github.com/burntcarrot/pm/drivers/db/project"
	"github.com/burntcarrot/pm/entity/project"
	"github.com/burntcarrot/pm/errors"
	"github.com/go-redis/redis/v8"
)

const MAX_FETCH_ROWS = 9 * 100000

type ProjectRepo struct {
	Conn *redis.Client
}

func NewProjectRepo(conn *redis.Client) project.DomainRepo {
	return &ProjectRepo{Conn: conn}
}

func (p *ProjectRepo) CreateProject(ctx context.Context, us project.Domain) (project.Domain, error) {
	projectsCounter := fmt.Sprintf("%s:projects:counter", strings.ToLower(us.Username))
	counterValue, projectCounterErr := p.Conn.Get(ctx, projectsCounter).Result()
	if projectCounterErr != nil {
		return project.Domain{}, errors.ErrInternalServerError
	}

	project_id := "project_" + counterValue

	createdProject := dbProject.Project{
		ID:          project_id,
		Name:        us.Name,
		Description: us.Description,
		Github:      us.Github,
	}

	raw, err := json.Marshal(createdProject)
	if err != nil {
		return project.Domain{}, errors.ErrInternalServerError
	}

	key := fmt.Sprintf("%s:projects", us.Username)

	insertErr := p.Conn.RPush(ctx, key, raw).Err()
	if insertErr != nil {
		return project.Domain{}, errors.ErrInternalServerError
	}

	// set counter for tasks while creating project itself
	// counter := fmt.Sprintf("%s:projects:counter", us.Username, strings.ToLower(us.Name))
	// counterErr := p.Conn.Set(ctx, counter, 1, 0).Err()

	// if counterErr != nil {
	// 	return project.Domain{}, errors.ErrInternalServerError
	// }

	incrErr := p.Conn.Incr(ctx, projectsCounter).Err()
	if incrErr != nil {
		return project.Domain{}, errors.ErrInternalServerError
	}

	return createdProject.ToDomain(), nil
}

func (p *ProjectRepo) GetProjects(ctx context.Context, username string) ([]project.Domain, error) {
	key := fmt.Sprintf("%s:projects", username)
	raw, err := p.Conn.LRange(ctx, key, 0, MAX_FETCH_ROWS).Result()
	if err != nil {
		return []project.Domain{}, errors.ErrInternalServerError
	}

	pr := new(dbProject.Project)
	var projects []project.Domain

	for _, j := range raw {
		if err := json.Unmarshal([]byte(j), pr); err != nil {
			return []project.Domain{}, errors.ErrInternalServerError
		}

		projects = append(projects, pr.ToDomain())
	}

	return projects, nil
}

func (p *ProjectRepo) GetProjectByName(ctx context.Context, username, projectName string) ([]project.Domain, error) {
	key := fmt.Sprintf("%s:projects", username)
	raw, err := p.Conn.LRange(ctx, key, 0, MAX_FETCH_ROWS).Result()
	if err != nil {
		return []project.Domain{}, errors.ErrInternalServerError
	}

	fmt.Println("raw:", raw)

	pr := new(dbProject.Project)
	var projects []project.Domain

	for _, j := range raw {
		if err := json.Unmarshal([]byte(j), pr); err != nil {
			return []project.Domain{}, errors.ErrInternalServerError
		}

		fmt.Println("pr ID:", pr.Name)
		fmt.Println("project ID:", projectName)

		if strings.EqualFold(pr.Name, projectName) {
			projects = append(projects, pr.ToDomain())
			return projects, nil
		}
	}

	return []project.Domain{}, nil
}
