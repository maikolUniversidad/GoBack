package entities

type PropiedadEntity struct {
    ID      string  `json:"id"`
    Name    string  `json:"name"`
    Price   float64 `json:"price"`
    OwnerID string  `json:"ownerId"`
}