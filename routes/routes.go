package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents) // this is how we register a handler for http methods
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)
}
