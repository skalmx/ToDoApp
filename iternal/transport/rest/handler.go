package rest

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine{
	r := gin.Default()

	api := r.Group("/api")
	{
		h.InitUsersRoutes(api)	
		h.InitListsRoutes(api)
		h.InitTasksRoutes(api)
	}
	return r
}
