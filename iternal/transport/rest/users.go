package rest

import "github.com/gin-gonic/gin"

func (h *Handler) InitUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.POST("/", h.createUser)
		users.PUT("/:userid", h.updateUser)
	}
}

func (h *Handler) createUser(c *gin.Context) {
	
}

func (h *Handler) updateUser(c *gin.Context) {
	
}