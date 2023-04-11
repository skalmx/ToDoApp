package rest

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine{
	r := gin.Default()

	app := r.Group("/app")
	{
		users := app.Group("/user")
		{
			users.POST("/create", h.createUser)
			users.PUT("/update/:id", h.updateUser)
		}
		
		lists := app.Group("/lists")
		{
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.POST("/", h.createList)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			tasks := lists.Group(":id/tasks")
			{
				tasks.GET("/", h.getAllTasksInList)
				tasks.POST("/", h.createTaskInList)
			}
		}

		tasks := app.Group("/tasks")
		{
			tasks.GET("/:id", h.getTaskById)
			tasks.PUT("/:id", h.updateTask)
			tasks.DELETE("/:id", h.deleteTask)
		}
	}
	
	return r
}
