package project

import (
	"fmt"

	"github.com/burntcarrot/pm/controllers"
	"github.com/burntcarrot/pm/entity/project"
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
	username := c.Param("userID")

	ctx := c.Request().Context()

	// get project
	u, err := p.Usecase.GetProjects(ctx, username)
	if err != nil {
		return err
	}

	var ps []GetResponse

	for _, j := range u {
		getResponse := GetResponse{
			ID:          j.ID,
			Name:        j.Name,
			Description: j.Description,
			Github:      j.Github,
		}

		ps = append(ps, getResponse)
	}

	fmt.Println("Woohoo fetched projects!")

	return controllers.Success(c, ps)
}

func (p *ProjectController) GetProjectByID(c echo.Context) error {
	username := c.Param("userID")
	projectID := c.Param("projectID")

	ctx := c.Request().Context()

	// get project
	u, err := p.Usecase.GetProjectByID(ctx, username, projectID)
	if err != nil {
		return err
	}

	var ps []GetResponse

	for _, j := range u {
		getResponse := GetResponse{
			ID:          j.ID,
			Name:        j.Name,
			Description: j.Description,
			Github:      j.Github,
		}

		ps = append(ps, getResponse)
	}

	fmt.Println("Woohoo fetched projects!")

	return controllers.Success(c, ps)
}

func (p *ProjectController) CreateProject(c echo.Context) error {
	projectCreate := CreateRequest{}
	err := c.Bind(&projectCreate)
	if err != nil {
		return err
	}

	// fetch context
	ctx := c.Request().Context()

	// TODO: check if project already exists

	// map project
	projectDomain := project.Domain{
		ID:          projectCreate.ID,
		Username:    projectCreate.Username,
		Name:        projectCreate.Name,
		Description: projectCreate.Description,
		Github:      projectCreate.Github,
	}

	// create project
	u, err := p.Usecase.CreateProject(ctx, projectDomain)
	if err != nil {
		return err
	}

	createResponse := CreateResponse{
		ID:          u.ID,
		Name:        u.Name,
		Description: u.Description,
		Github:      u.Github,
	}

	fmt.Println("Woohoo project created!")

	return controllers.Success(c, createResponse)
}
