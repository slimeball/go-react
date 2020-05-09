package main

import (
	"bookstore/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	gin := gin.Default()
	// gin.Use(gin.Logger())
	gin.Use(cors.Default())

	router.Routes(gin)
	gin.Run(":5050")
}
