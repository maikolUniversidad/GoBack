package routes

import (
    "github.com/gin-gonic/gin"
    "go-backend/controller"
    "go-backend/services/implementation"
)

func RegisterRoutes(r *gin.Engine) {
    personaService := &implementation.PersonaService{}
    propiedadService := &implementation.PropiedadService{}

    r.POST("/persona", controller.CreatePersona(personaService))
    r.GET("/personas", controller.GetPersonas(personaService))
    r.POST("/propiedad", controller.CreatePropiedad(propiedadService))
    r.GET("/propiedades", controller.GetPropiedades(propiedadService))
}