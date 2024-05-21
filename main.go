package main

import (
	
	"example.com/rest-api/db"
	_ "example.com/rest-api/docs"
	"example.com/rest-api/routes"
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

// Router Get the gin router with all the routes defined
//
//	@title					Event Management API
//	@version				1.0.0
//	@description				A simple event management API featuring functionalities such as event creation, updating, deletion, and retrieval. Users can register for events, cancel registrations, and utilize user registration and signup services. Documented using Swagger!
//
//	@contact.name				Divij Sharma
//	@contact.url				https://github.com/dvjsharma
//	@contact.email				divijs75@gmail.com
//
//	@host					localhost:3000
//	@BasePath				/api/v1
//
//	@securityDefinitions.apikey		ApiKeyAuth
//	@in					header
//	@name					Authorization
//	@description				Bearer Token from /login endpoint

func main() {
	db.InitDB()
	server := gin.Default()
	api := server.Group("/api/v1")

	api.GET("/events", routes.GetEvents)
	api.GET("/events/:id", routes.GetEvent)

	authenticated := api.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", routes.CreateEvent)
	authenticated.PUT("/events/:id", routes.UpdateEvent)
	authenticated.DELETE("/events/:id", routes.DeleteEvent)
	authenticated.POST("/events/:id/register", routes.RegisterForEvent)
	authenticated.DELETE("/events/:id/register", routes.CancelRegistration)

	api.POST("/signup", routes.Signup)
	api.POST("/login", routes.Login)

	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.Run(":3000")
}
