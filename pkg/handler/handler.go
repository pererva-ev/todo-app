package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pererva-ev/todo-app/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		project := api.Group("/project")
		{
			project.POST("/", h.createProject)
			project.GET("/", h.getAllProjects)
			project.GET("/:id", h.getProjectByID)
			project.PUT("/:id", h.updateProject)
			project.DELETE("/:id", h.deleteProject)
		}
		task := api.Group("/task")
		{
			task.POST("/", h.createTask)
			task.GET("/", h.getAllTasks)
			task.GET("/:task_id", h.getTaskByID)
			task.PUT("/:task_id", h.updateTask)
			task.DELETE("/:task_id", h.deleteTask)
		}
		comment := api.Group("/comment")
		{
			comment.POST("/", h.createComment)
			comment.GET("/", h.getAllComments)
			comment.GET("/:comment_id", h.getCommentByID)
			comment.PUT("/:comment_id", h.updateComment)
			comment.DELETE("/:comment_id", h.deleteComment)
		}
		column := api.Group("/colomn")
		{
			column.POST("/", h.createColumn)
			column.GET("/", h.getAllColumns)
			column.GET("/:column_id", h.getColumnByID)
			column.PUT("/:column_id", h.updateColumn)
			column.DELETE("/:column_id", h.deleteColumn)
		}

	}

	return router
}
