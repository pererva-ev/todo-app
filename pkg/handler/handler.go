package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		project := api.Group("/project")
		{
			project.POST("/", h.createProject)
			project.GET("/", h.getAllProjects)
			project.GET("/:id", h.getProjectById)
			project.PUT("/:id", h.updateProject)
			project.DELETE("/:id", h.deleteProject)
		}
		task := api.Group("/task")
		{
			task.POST("/", h.createTask)
			task.GET("/", h.getAllTasks)
			task.GET("/:task_id", h.getTaskById)
			task.PUT("/:task_id", h.updateTask)
			task.DELETE("/:comment_id", h.deleteTask)
		}
		comment := api.Group("/comment")
		{
			comment.POST("/", h.createTask)
			comment.GET("/", h.getAllTasks)
			comment.GET("/:comment_id", h.getTaskById)
			comment.PUT("/:comment_id", h.updateTask)
			comment.DELETE("/:comment_id", h.deleteTask)
		}
		column := api.Group("/colomn")
		{
			column.POST("/", h.createTask)
			column.GET("/", h.getAllTasks)
			column.GET("/:column_id", h.getTaskById)
			column.PUT("/:column_id", h.updateTask)
			column.DELETE("/:column_id", h.deleteTask)
		}

	}

	return router
}
