package rest

import "github.com/gin-gonic/gin"

func (h *Handler) InitTasksRoutes(api *gin.RouterGroup) {
	tasks := api.Group("list/:listid/tasks")
	{
		tasks.GET("/", h.getAllTasks)
		tasks.POST("/", h.createTask)
		tasks.GET("/:taskid", h.getTaskById)
		tasks.PUT("/:taskid", h.updateTask)
		tasks.DELETE("/:taskid", h.deleteTask)
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