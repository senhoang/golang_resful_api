package main

import (
	"web_kenda_api/pkg/apis/auth"
	"web_kenda_api/pkg/apis/userapi"
	"web_kenda_api/pkg/database"
	"web_kenda_api/pkg/middlewares"
	"web_kenda_api/pkg/printcolor"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// "github.com/joho/godotenv"
)

func main() {
	printcolor.PrintlnY("Server starting ...")
	godotenv.Load()        // load .env file
	database.InitPostgre() // Tạo kết nối postgre database

	r := gin.Default()
	r.Use(middlewares.CORS())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/login", auth.Login)
	r.GET("/info", auth.Info)

	user := r.Group("/api")
	{
		user.GET("/user", userapi.GetUsers)
		user.GET("/user/:deptid", userapi.GetUserByDept)
		user.POST("/user", userapi.CreateUser)
		user.PUT("/user", userapi.UpdateUser) //Update all fields
		user.DELETE("/user/:id", userapi.DeleteUser)
	}

	err := r.Run(":1112")
	if err != nil {
		panic(err)
	}
}
