package project_test

import (
	"context"
	"testing"
	"time"

	"github.com/burntcarrot/pm/entity/project"
	"github.com/burntcarrot/pm/entity/project/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var projectRepo mocks.DomainRepo
var projectService project.DomainService
var projectDomain project.Domain

func setup() {
	projectService = project.NewUsecase(&projectRepo, time.Minute*15)
	projectDomain = project.Domain{
		ID:          "project_1",
		Username:    "burntcarrot",
		Name:        "PM",
		Description: "Project Management made easy",
		Github:      "github.com/burntcarrot/pm",
	}
}

func TestCreateProject(t *testing.T) {
	setup()
	projectRepo.On("CreateProject", mock.Anything, mock.AnythingOfType("Domain")).Return(projectDomain, nil).Once()
	t.Run("Valid Project Creation", func(t *testing.T) {
		project, err := projectService.CreateProject(context.Background(), projectDomain)
		assert.Nil(t, err)
		assert.Equal(t, projectDomain.Name, project.Name)
	})

	t.Run("Invalid Project Creation", func(t *testing.T) {
		_, err := projectService.CreateProject(context.Background(), project.Domain{
			Username:    "",
			Name:        "",
			Description: "",
			Github:      "",
		})

		assert.NotNil(t, err)
	})
}

func TestGetProjects(t *testing.T) {
	setup()
	projectRepo.On("GetProjects", mock.Anything, mock.AnythingOfType("string")).Return([]project.Domain{projectDomain}, nil).Once()

	t.Run("Get Projects", func(t *testing.T) {
		projects, err := projectService.GetProjects(context.Background(), "burntcarrot")
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if len(projects) == 0 {
			t.Error("No projects fetched.")
		}

		assert.Nil(t, err)
		assert.Equal(t, projectDomain, projects[0])
	})
}

func TestGetProjectByID(t *testing.T) {
	setup()
	projectRepo.On("GetProjectByName", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]project.Domain{projectDomain}, nil)

	t.Run("Valid Get Project by Name", func(t *testing.T) {
		projects, err := projectService.GetProjectByName(context.Background(), "burntcarrot", "project_1")
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if len(projects) == 0 {
			t.Error("No projects fetched.")
		}

		assert.Nil(t, err)
		assert.Equal(t, projectDomain, projects[0])
	})

	t.Run("Invalid Get Project by Name", func(t *testing.T) {
		project, err := projectService.GetProjectByName(context.Background(), "burntcarrot", "project_nonexistent")
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		assert.Nil(t, err)
		assert.NotEqual(t, project, projectDomain)
	})
}
