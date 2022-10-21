package main

import (
	"FP1hacktiv8/docs"
	"FP1hacktiv8/routes"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title 			Hacktiv8 Final Project 1 - Group 1
// @version 		1.0
// @description 	This is a sample server celler server.
// @contact.name 	API Support
// @contact.url 	http://www.swagger.io/support
// @contact.email 	support@swagger.io

// @license.name 	Apache 2.0
// @license.url 	http://www.apache.org/licenses/LICENSE-2.0.html

// @host 			localhost:8080
// @BasePath 		/
// @query.collection.format multi
// @securityDefinitions.basic  BasicAuth

func main() {

	port := os.Getenv("PORT")
	host := os.Getenv("BASE_URL")
	router := gin.Default()
	routes.TodoRouters(router)

	docs.SwaggerInfo.BasePath = "/"
	url := ginSwagger.URL(host + ":" + port + "/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.Run(":" + port)

}
