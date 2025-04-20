package implementation

import "go-backend/model/entities"

var personas []entities.PersonaEntity

type PersonaService struct{}

func (s *PersonaService) Create(p entities.PersonaEntity) error {
    personas = append(personas, p)
    return nil
}

func (s *PersonaService) GetAll() []entities.PersonaEntity {
    return personas
}