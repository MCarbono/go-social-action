package entity

import "time"

type Volunteer struct {
	ID           string
	FirstName    string
	LastName     string
	Neighborhood string
	City         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
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
