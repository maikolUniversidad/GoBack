package main

import (
    "github.com/gin-gonic/gin"
    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"

    _ "go-backend/docs"
    "go-backend/routes"
)

// @title API - Persona y Propiedad
// @version 1.0
// @description Backend para EDA demo
// @host localhost:8080
// @BasePath /
func main() {
    r := gin.Default()
    routes.RegisterRoutes(r)
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    r.Run(":8080")
}