package task

import (
	"fmt"

	"github.com/burntcarrot/pm/controllers"
	"github.com/burntcarrot/pm/entity/task"
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
	username := c.Param("userID")
	projectName := c.Param("projectName")

	ctx := c.Request().Context()

	// get task
	u, err := t.Usecase.GetTasks(ctx, username, projectName)
	if err != nil {
		return err
	}

	var ps []GetResponse

	for _, j := range u {
		getResponse := GetResponse{
			ID:       j.ID,
			Name:     j.Name,
			Type:     j.Type,
			Deadline: j.Type,
			Status:   j.Status,
		}

		ps = append(ps, getResponse)
	}

	fmt.Println("Woohoo fetched tasks!")

	return controllers.Success(c, ps)
}

func (t *TaskController) GetTaskByName(c echo.Context) error {
	username := c.Param("userID")
	projectName := c.Param("projectName")
	taskName := c.Param("taskName")

	ctx := c.Request().Context()

	// get task
	u, err := t.Usecase.GetTaskByName(ctx, username, projectName, taskName)
	if err != nil {
		return err
	}

	var ps []GetResponse

	for _, j := range u {
		getResponse := GetResponse{
			ID:       j.ID,
			Name:     j.Name,
			Type:     j.Type,
			Deadline: j.Type,
			Status:   j.Status,
		}

		ps = append(ps, getResponse)
	}

	fmt.Println("Woohoo fetched tasks!")

	return controllers.Success(c, ps)
}

func (t *TaskController) CreateTask(c echo.Context) error {
	taskCreate := CreateRequest{}
	err := c.Bind(&taskCreate)
	if err != nil {
		return err
	}

	// fetch context
	ctx := c.Request().Context()

	// TODO: check if task already exists

	// map task
	taskDomain := task.Domain{
		ID:          taskCreate.ID,
		Username:    taskCreate.Username,
		ProjectName: taskCreate.ProjectName,
		Name:        taskCreate.Name,
		Type:        taskCreate.Type,
		Deadline:    taskCreate.Deadline,
		Status:      taskCreate.Status,
	}

	// create task
	u, err := t.Usecase.CreateTask(ctx, taskDomain)
	if err != nil {
		return err
	}

	createResponse := CreateResponse{
		ID:       u.ID,
		Name:     u.Name,
		Type:     u.Type,
		Deadline: u.Deadline,
		Status:   u.Status,
	}

	fmt.Println("Woohoo task created!")

	return controllers.Success(c, createResponse)
}
