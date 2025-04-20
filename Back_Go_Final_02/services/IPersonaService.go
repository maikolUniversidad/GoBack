package services

import "go-backend/model/entities"

type IPersonaService interface {
    Create(p entities.PersonaEntity) error
    GetAll() []entities.PersonaEntity
}