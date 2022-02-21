package task_test

import (
	"context"
	"testing"
	"time"

	"github.com/burntcarrot/quii/entity/task"
	"github.com/burntcarrot/quii/entity/task/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var taskRepo mocks.DomainRepo
var taskService task.DomainService
var taskDomain task.Domain

func setup() {
	taskService = task.NewUsecase(&taskRepo, time.Minute*15)
	taskDomain = task.Domain{
		ID:          "task_1",
		Username:    "burntcarrot",
		ProjectName: "quii",
		Type:        "feature",
		Name:        "Add Login endpoint",
		Status:      "doing",
	}
}

func TestCreateTask(t *testing.T) {
	setup()
	taskRepo.On("CreateTask", mock.Anything, mock.AnythingOfType("Domain")).Return(taskDomain, nil).Once()
	t.Run("Valid Task Creation", func(t *testing.T) {
		task, err := taskService.CreateTask(context.Background(), taskDomain)
		assert.Nil(t, err)
		assert.Equal(t, taskDomain.Name, task.Name)
	})

	t.Run("Invalid Task Creation", func(t *testing.T) {
		_, err := taskService.CreateTask(context.Background(), task.Domain{
			Username: "",
			Name:     "",
			Type:     "",
			Status:   "",
		})

		assert.NotNil(t, err)
	})
}

func TestGetTasks(t *testing.T) {
	setup()
	taskRepo.On("GetTasks", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]task.Domain{taskDomain}, nil).Once()

	t.Run("Get Tasks", func(t *testing.T) {
		tasks, err := taskService.GetTasks(context.Background(), "burntcarrot", "quii")
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if len(tasks) == 0 {
			t.Error("No tasks fetched.")
		}

		assert.Nil(t, err)
		assert.Equal(t, taskDomain, tasks[0])
	})
}

func TestGetTaskByName(t *testing.T) {
	setup()
	taskRepo.On("GetTaskByName", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]task.Domain{taskDomain}, nil)

	t.Run("Valid Get Task by Name", func(t *testing.T) {
		tasks, err := taskService.GetTaskByName(context.Background(), "burntcarrot", "quii", "task_1")
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if len(tasks) == 0 {
			t.Error("No tasks fetched.")
		}

		assert.Nil(t, err)
		assert.Equal(t, taskDomain, tasks[0])
	})

	t.Run("Invalid Get Task by Name", func(t *testing.T) {
		task, err := taskService.GetTaskByName(context.Background(), "burntcarrot", "quii", "task_nonexistent")
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		assert.Nil(t, err)
		assert.NotEqual(t, task, taskDomain)
	})
}
