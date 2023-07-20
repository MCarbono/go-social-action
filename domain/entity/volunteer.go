package entity

import "time"

type Volunteer struct {
	ID           string    `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Neighborhood string    `json:"neighborhood"`
	City         string    `json:"city"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func NewVolunteer(ID, firstName, lastName, neighborhood, city string, createdAt, updatedAt time.Time) *Volunteer {
	return &Volunteer{
		ID:           ID,
		FirstName:    firstName,
		LastName:     lastName,
		Neighborhood: neighborhood,
		City:         city,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
}
