package project

import (
	"fmt"
	"net/http"

	"github.com/burntcarrot/pm/controllers"
	"github.com/burntcarrot/pm/entity/project"
	"github.com/burntcarrot/pm/errors"
	"github.com/labstack/echo/v4"
)

type ProjectController struct {
	Usecase project.Usecase
}

func NewProjectController(u project.Usecase) *ProjectController {
	return &ProjectController{
		Usecase: u,
	}
}

func (p *ProjectController) GetProjects(c echo.Context) error {
	username := c.Param("userName")

	ctx := c.Request().Context()

	// get projects
	projects, err := p.Usecase.GetProjects(ctx, username)
	if err == errors.ErrValidationFailed {
		return controllers.Error(c, http.StatusBadRequest, errors.ErrValidationFailed)
	}
	if err != nil {
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

	fmt.Println("Woohoo fetched projects!")

	return controllers.Success(c, response)
}

func (p *ProjectController) GetProjectByName(c echo.Context) error {
	username := c.Param("userName")
	projectName := c.Param("projectName")

	ctx := c.Request().Context()

	// get project
	projects, err := p.Usecase.GetProjectByName(ctx, username, projectName)
	if err == errors.ErrValidationFailed {
		return controllers.Error(c, http.StatusBadRequest, errors.ErrValidationFailed)
	}
	if err != nil {
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

	fmt.Println("Woohoo fetched projects!")

	return controllers.Success(c, response)
}

func (p *ProjectController) CreateProject(c echo.Context) error {
	projectRequest := CreateRequest{}
	err := c.Bind(&projectRequest)
	if err != nil {
		return controllers.Error(c, http.StatusInternalServerError, errors.ErrInternalServerError)
	}

	// fetch context
	ctx := c.Request().Context()

	pr, err := p.Usecase.GetProjectByName(ctx, projectRequest.Username, projectRequest.Name)
	if err != nil {
		return controllers.Error(c, http.StatusInternalServerError, errors.ErrInternalServerError)
	}
	fmt.Println(pr)
	if len(pr) != 0 {
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
		return controllers.Error(c, http.StatusBadRequest, errors.ErrValidationFailed)
	}
	if err != nil {
		return controllers.Error(c, http.StatusInternalServerError, errors.ErrInternalServerError)
	}

	response := CreateResponse{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		Github:      project.Github,
	}

	fmt.Println("Woohoo project created!")

	return controllers.Success(c, response)
}
