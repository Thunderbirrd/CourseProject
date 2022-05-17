package handler

import (
	"github.com/Thunderbirrd/CourseProject/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		api.GET("/all", h.getAllUsers)
		api.GET("/user/:id", h.getUserById)
		api.GET("/by-service/:service", h.getUsersByServiceType)
		api.GET("/by-location/:location", h.getUsersByLocation)
	}

	user := router.Group("/user", h.userIdentity)
	{
		user.POST("/update", h.updateUser)
	}

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders(authorizationHeader)
	router.Use(cors.New(config))
	return router
}
