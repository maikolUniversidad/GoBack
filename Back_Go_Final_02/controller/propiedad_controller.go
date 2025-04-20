package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "go-backend/model/entities"
    "go-backend/services"
)

func CreatePropiedad(service services.IPropiedadService) gin.HandlerFunc {
    return func(c *gin.Context) {
        var propiedad entities.PropiedadEntity
        if err := c.ShouldBindJSON(&propiedad); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        service.Create(propiedad)
        c.JSON(http.StatusOK, entities.StatusResponse{Status: "propiedad creada"})
    }
}

func GetPropiedades(service services.IPropiedadService) gin.HandlerFunc {
    return func(c *gin.Context) {
        c.JSON(http.StatusOK, service.GetAll())
    }
}