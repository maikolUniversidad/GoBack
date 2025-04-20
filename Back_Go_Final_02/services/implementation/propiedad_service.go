package implementation

import "go-backend/model/entities"

var propiedades []entities.PropiedadEntity

type PropiedadService struct{}

func (s *PropiedadService) Create(p entities.PropiedadEntity) error {
    propiedades = append(propiedades, p)
    return nil
}

func (s *PropiedadService) GetAll() []entities.PropiedadEntity {
    return propiedades
}