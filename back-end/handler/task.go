package handler

import (
	"net/http"
	"strconv"

	"example.com/hello/service"
	"github.com/gin-gonic/gin"
)

type taskHandler struct {
	tsrv service.TaskService
}

func NewTaskHandler(tsrv service.TaskService) taskHandler {
	return taskHandler{tsrv: tsrv}
}

func (h taskHandler) GetTaskAllCos(c *gin.Context) {
	customers, err := h.tsrv.GetTaskAllCos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customers)
}

func (h taskHandler) GetTaskByIdCos(c *gin.Context) {
	taskID, _ := strconv.Atoi(c.Param("taskID"))

	customers, err := h.tsrv.GetTaskByIdCos(taskID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, customers)
}

func (h taskHandler) CreateTaskCos(c *gin.Context) {
	var task service.TaskRequest
	err := c.ShouldBindJSON(&task)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	tx, err := h.tsrv.CreateTaskCos(task)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, tx)
}

func (h taskHandler) UpdateTaskCos(c *gin.Context) {
	taskID, _ := strconv.Atoi(c.Param("taskID"))
	var task service.TaskRequest
	err := c.ShouldBindJSON(&task)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	tx, err := h.tsrv.UpdateTaskCos(taskID, task)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, tx)
}
func (h taskHandler) DeleteTaskCos(c *gin.Context) {
	taskID, _ := strconv.Atoi(c.Param("taskID"))
	err := h.tsrv.DeleteTaskCos(taskID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}
