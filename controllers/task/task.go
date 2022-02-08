package task

import (
	"fmt"
	"net/http"

	"github.com/burntcarrot/pm/controllers"
	"github.com/burntcarrot/pm/entity/task"
	"github.com/burntcarrot/pm/errors"
	"github.com/labstack/echo/v4"
)

type TaskController struct {
	Usecase task.Usecase
}

func NewTaskController(u task.Usecase) *TaskController {
	return &TaskController{
		Usecase: u,
	}
}

func (t *TaskController) GetTasks(c echo.Context) error {
	username := c.Param("userName")
	projectName := c.Param("projectName")

	ctx := c.Request().Context()

	// get tasks
	tasks, err := t.Usecase.GetTasks(ctx, username, projectName)
	if err == errors.ErrValidationFailed {
		return controllers.Error(c, http.StatusBadRequest, errors.ErrValidationFailed)
	}
	if err != nil {
		return controllers.Error(c, http.StatusInternalServerError, errors.ErrInternalServerError)
	}

	var response []GetResponse

	for _, task := range tasks {
		getResponse := GetResponse{
			ID:       task.ID,
			Name:     task.Name,
			Type:     task.Type,
			Deadline: task.Type,
			Status:   task.Status,
		}

		response = append(response, getResponse)
	}

	fmt.Println("Woohoo fetched tasks!")

	return controllers.Success(c, response)
}

func (t *TaskController) GetTaskByName(c echo.Context) error {
	username := c.Param("userName")
	projectName := c.Param("projectName")
	taskName := c.Param("taskName")

	ctx := c.Request().Context()

	// get task
	tasks, err := t.Usecase.GetTaskByName(ctx, username, projectName, taskName)
	if err == errors.ErrValidationFailed {
		return controllers.Error(c, http.StatusBadRequest, errors.ErrValidationFailed)
	}
	if err != nil {
		return controllers.Error(c, http.StatusInternalServerError, errors.ErrInternalServerError)
	}

	var response []GetResponse

	for _, task := range tasks {
		getResponse := GetResponse{
			ID:       task.ID,
			Name:     task.Name,
			Type:     task.Type,
			Deadline: task.Type,
			Status:   task.Status,
		}

		response = append(response, getResponse)
	}

	fmt.Println("Woohoo fetched tasks!")

	return controllers.Success(c, response)
}

func (t *TaskController) CreateTask(c echo.Context) error {
	taskRequest := CreateRequest{}
	err := c.Bind(&taskRequest)
	if err != nil {
		return controllers.Error(c, http.StatusBadRequest, errors.ErrBadRequest)
	}

	// fetch context
	ctx := c.Request().Context()

	// TODO: check if task already exists

	taskDomain := task.Domain{
		ID:          taskRequest.ID,
		Username:    taskRequest.Username,
		ProjectName: taskRequest.ProjectName,
		Name:        taskRequest.Name,
		Type:        taskRequest.Type,
		Deadline:    taskRequest.Deadline,
		Status:      taskRequest.Status,
	}

	// create task
	task, err := t.Usecase.CreateTask(ctx, taskDomain)
	if err == errors.ErrValidationFailed {
		return controllers.Error(c, http.StatusBadRequest, errors.ErrValidationFailed)
	}
	if err != nil {
		return controllers.Error(c, http.StatusInternalServerError, errors.ErrInternalServerError)
	}

	response := CreateResponse{
		ID:       task.ID,
		Name:     task.Name,
		Type:     task.Type,
		Deadline: task.Deadline,
		Status:   task.Status,
	}

	fmt.Println("Woohoo task created!")

	return controllers.Success(c, response)
}
