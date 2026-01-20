package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/higorworkout/todo-api/internal/application"
	"github.com/higorworkout/todo-api/internal/domain"
)

type TaskHandler struct {
	uc *application.TaskUseCase
}

func NewTaskHandler(r *gin.Engine, uc *application.TaskUseCase) {
	handler := &TaskHandler{uc: uc}

	r.POST("/tasks", handler.create)
	r.GET("/tasks", handler.list)
	r.GET("/tasks/:id", handler.get)
	r.PUT("/tasks/:id", handler.update)
	r.DELETE("/tasks/:id", handler.delete)
}

func (h *TaskHandler) create(c *gin.Context) {
	var body struct {
		Title string `json:"title"`
	}
	c.BindJSON(&body)

	task, _ := h.uc.CreateTask(c.Request.Context(), body.Title)

	c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) list(c *gin.Context) {
	tasks, _ := h.uc.ListTasks(c.Request.Context())
	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) get(c *gin.Context) {
	task, _ := h.uc.GetTask(c.Request.Context(), c.Param("id"))
	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) update(c *gin.Context) {
	var task domain.Task
	c.BindJSON(&task)
	task.ID = c.Param("id")

	h.uc.UpdateTask(c.Request.Context(), &task)
	c.Status(http.StatusNoContent)
}

func (h *TaskHandler) delete(c *gin.Context) {
	h.uc.DeleteTask(c.Request.Context(), c.Param("id"))
	c.Status(http.StatusNoContent)
}