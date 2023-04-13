package rest

import "github.com/gin-gonic/gin"

func (h *Handler) InitTasksRoutes(api *gin.RouterGroup) {
	tasks := api.Group("list/:id/tasks")
	{
		tasks.GET("/", h.getAllTasks)
		tasks.POST("/", h.createTask)
		tasks.GET("/:id", h.getTaskById)
		tasks.PUT("/:id", h.updateTask)
		tasks.DELETE("/:id", h.deleteTask)
	}
}

func (h *Handler) getTaskById(c *gin.Context) {
	
}
func (h *Handler) updateTask(c *gin.Context) {
	
}

func (h *Handler) deleteTask(c *gin.Context) {
	
}

func (h *Handler) getAllTasks(c *gin.Context) {
	
}

func (h *Handler) createTask(c *gin.Context) {
	
}