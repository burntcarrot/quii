package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	dbProject "github.com/burntcarrot/quii/drivers/db/project"
	"github.com/burntcarrot/quii/entity/project"
	"github.com/burntcarrot/quii/errors"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

const MAX_FETCH_ROWS = 9 * 100000

type ProjectRepo struct {
	Conn   *redis.Client
	Logger *zap.SugaredLogger
}

func NewProjectRepo(conn *redis.Client, l *zap.SugaredLogger) project.DomainRepo {
	return &ProjectRepo{Conn: conn, Logger: l}
}

func (p *ProjectRepo) CreateProject(ctx context.Context, us project.Domain) (project.Domain, error) {
	projectsCounter := fmt.Sprintf("%s:projects:counter", strings.ToLower(us.Username))
	counterValue, projectCounterErr := p.Conn.Get(ctx, projectsCounter).Result()
	if projectCounterErr != nil {
		p.Logger.Error("[createproject] failed to get project counter")
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
		p.Logger.Errorf("[createproject] failed to marshal project: %v", err)
		return project.Domain{}, errors.ErrInternalServerError
	}

	key := fmt.Sprintf("%s:projects", us.Username)

	insertErr := p.Conn.RPush(ctx, key, raw).Err()
	if insertErr != nil {
		p.Logger.Errorf("[createproject] failed to push project in redis: %v", err)
		return project.Domain{}, errors.ErrInternalServerError
	}

	incrErr := p.Conn.Incr(ctx, projectsCounter).Err()
	if incrErr != nil {
		p.Logger.Errorf("[createproject] failed to increment project counter: %v", err)
		return project.Domain{}, errors.ErrInternalServerError
	}

	// set counter for task while creating project
	counter := fmt.Sprintf("%s:projects:%s:tasks:counter", strings.ToLower(us.Username), strings.ToLower(us.Name))
	counterErr := p.Conn.Set(ctx, counter, 1, 0).Err()
	if counterErr != nil {
		p.Logger.Errorf("[createproject] failed to set task counter: %v", err)
		return project.Domain{}, errors.ErrInternalServerError
	}

	return createdProject.ToDomain(), nil
}

func (p *ProjectRepo) GetProjects(ctx context.Context, username string) ([]project.Domain, error) {
	key := fmt.Sprintf("%s:projects", username)
	raw, err := p.Conn.LRange(ctx, key, 0, MAX_FETCH_ROWS).Result()
	if err != nil {
		p.Logger.Errorf("[getprojects] failed to get projects: %v", err)
		return []project.Domain{}, errors.ErrInternalServerError
	}

	pr := new(dbProject.Project)
	var projects []project.Domain

	for _, j := range raw {
		if err := json.Unmarshal([]byte(j), pr); err != nil {
			p.Logger.Errorf("[getprojects] failed to unmarshal project: %v", err)
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
		p.Logger.Errorf("[getprojectbyname] failed to get projects: %v", err)
		return []project.Domain{}, errors.ErrInternalServerError
	}

	pr := new(dbProject.Project)
	var projects []project.Domain

	for _, j := range raw {
		if err := json.Unmarshal([]byte(j), pr); err != nil {
			p.Logger.Errorf("[getprojectbyname] failed to unmarshal project: %v", err)
			return []project.Domain{}, errors.ErrInternalServerError
		}

		if strings.EqualFold(pr.Name, projectName) {
			projects = append(projects, pr.ToDomain())
			return projects, nil
		}
	}

	return []project.Domain{}, nil
}
