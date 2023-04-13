package rest

import "github.com/gin-gonic/gin"

func (h *Handler) InitListsRoutes(api *gin.RouterGroup) {
	lists := api.Group("/lists")
	{
		lists.GET("/", h.getAllLists)
		lists.GET("/:id", h.getListById)
		lists.POST("/", h.createList)
		lists.PUT("/:id", h.updateList)
		lists.DELETE("/:id", h.deleteList)
	}
}

func (h *Handler) getAllLists(c *gin.Context) {
	
}

func (h *Handler) getListById(c *gin.Context) {
	
}

func (h *Handler) createList(c *gin.Context) {
	
}

func (h *Handler) updateList(c *gin.Context) {
	
}

func (h *Handler) deleteList(c *gin.Context) {
	
}