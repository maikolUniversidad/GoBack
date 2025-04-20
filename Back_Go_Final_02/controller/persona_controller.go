package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "go-backend/model/entities"
    "go-backend/services"
)

func CreatePersona(service services.IPersonaService) gin.HandlerFunc {
    return func(c *gin.Context) {
        var persona entities.PersonaEntity
        if err := c.ShouldBindJSON(&persona); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        service.Create(persona)
        c.JSON(http.StatusOK, entities.StatusResponse{Status: "persona creada"})
    }
}

func GetPersonas(service services.IPersonaService) gin.HandlerFunc {
    return func(c *gin.Context) {
        c.JSON(http.StatusOK, service.GetAll())
    }
}