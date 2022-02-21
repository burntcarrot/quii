package project

import (
	"net/http"

	"github.com/burntcarrot/quii/controllers"
	"github.com/burntcarrot/quii/entity/project"
	"github.com/burntcarrot/quii/errors"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type ProjectController struct {
	Usecase project.Usecase
	Logger  *zap.SugaredLogger
}

func NewProjectController(u project.Usecase, l *zap.SugaredLogger) *ProjectController {
	return &ProjectController{
		Usecase: u,
		Logger:  l,
	}
}

func (p *ProjectController) GetProjects(c echo.Context) error {
	username := c.Param("userName")

	ctx := c.Request().Context()

	// get projects
	projects, err := p.Usecase.GetProjects(ctx, username)
	if err == errors.ErrValidationFailed {
		p.Logger.Error("[getprojects] validation failed")
		return controllers.Error(c, http.StatusBadRequest, errors.ErrValidationFailed)
	}
	if err != nil {
		p.Logger.Errorf("[getprojects] failed to get projects: %v", err)
		return controllers.Error(c, http.StatusInternalServerError, errors.ErrInternalServerError)
	}

	var response []GetResponse

	for _, project := range projects {
		getResponse := GetResponse{
			ID:          project.ID,
			Name:        project.Name,
			Description: project.Description,
			Github:      project.Github,
		}

		response = append(response, getResponse)
	}

	p.Logger.Infof("[getprojects] fetched %d projects", len(response))

	return controllers.Success(c, response)
}

func (p *ProjectController) GetProjectByName(c echo.Context) error {
	username := c.Param("userName")
	projectName := c.Param("projectName")

	ctx := c.Request().Context()

	// get project
	projects, err := p.Usecase.GetProjectByName(ctx, username, projectName)
	if err == errors.ErrValidationFailed {
		p.Logger.Error("[getprojectbyname] validation failed")
		return controllers.Error(c, http.StatusBadRequest, errors.ErrValidationFailed)
	}
	if err != nil {
		p.Logger.Errorf("[getprojectbyname] failed to get projects: %v", err)
		return controllers.Error(c, http.StatusInternalServerError, errors.ErrInternalServerError)
	}

	var response []GetResponse

	for _, project := range projects {
		getResponse := GetResponse{
			ID:          project.ID,
			Name:        project.Name,
			Description: project.Description,
			Github:      project.Github,
		}

		response = append(response, getResponse)
	}

	return controllers.Success(c, response)
}

func (p *ProjectController) CreateProject(c echo.Context) error {
	projectRequest := CreateRequest{}
	err := c.Bind(&projectRequest)
	if err != nil {
		p.Logger.Errorf("[createproject] bad project creation request: %v", err)
		return controllers.Error(c, http.StatusInternalServerError, errors.ErrInternalServerError)
	}

	// fetch context
	ctx := c.Request().Context()

	pr, err := p.Usecase.GetProjectByName(ctx, projectRequest.Username, projectRequest.Name)
	if err != nil {
		p.Logger.Errorf("[createproject] failed to get projects: %v", err)
		return controllers.Error(c, http.StatusInternalServerError, errors.ErrInternalServerError)
	}
	if len(pr) != 0 {
		p.Logger.Error("[createproject] project already exists")
		return controllers.Error(c, http.StatusBadRequest, errors.ErrProjectAlreadyExists)
	}

	projectDomain := project.Domain{
		ID:          projectRequest.ID,
		Username:    projectRequest.Username,
		Name:        projectRequest.Name,
		Description: projectRequest.Description,
		Github:      projectRequest.Github,
	}

	// create project
	project, err := p.Usecase.CreateProject(ctx, projectDomain)
	if err == errors.ErrValidationFailed {
		p.Logger.Error("[createproject] validation failed")
		return controllers.Error(c, http.StatusBadRequest, errors.ErrValidationFailed)
	}
	if err != nil {
		p.Logger.Errorf("[createproject] failed to create project: %v", err)
		return controllers.Error(c, http.StatusInternalServerError, errors.ErrInternalServerError)
	}

	response := CreateResponse{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		Github:      project.Github,
	}

	p.Logger.Info("[createproject] created project: %s", project.ID)

	return controllers.Success(c, response)
}
