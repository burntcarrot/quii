package task

import (
	"net/http"

	"github.com/burntcarrot/quii/controllers"
	"github.com/burntcarrot/quii/entity/task"
	"github.com/burntcarrot/quii/errors"
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
			Deadline: task.Deadline,
			Status:   task.Status,
		}

		response = append(response, getResponse)
	}

	return controllers.Success(c, response)
}

func (t *TaskController) GetTaskByID(c echo.Context) error {
	username := c.Param("userName")
	projectName := c.Param("projectName")
	taskID := c.Param("taskID")

	ctx := c.Request().Context()

	// get task
	tasks, err := t.Usecase.GetTaskByID(ctx, username, projectName, taskID)
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
			Deadline: task.Deadline,
			Status:   task.Status,
		}

		response = append(response, getResponse)
	}

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

	ts, _ := t.Usecase.GetTaskByName(ctx, taskRequest.Username, taskRequest.ProjectName, taskRequest.Name)
	if len(ts) != 0 {
		return controllers.Error(c, http.StatusBadRequest, errors.ErrTaskAlreadyExists)
	}

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

	return controllers.Success(c, response)
}
