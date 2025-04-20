package services

import "go-backend/model/entities"

type IPropiedadService interface {
    Create(p entities.PropiedadEntity) error
    GetAll() []entities.PropiedadEntity
}