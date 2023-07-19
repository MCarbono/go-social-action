package repository

import "time"

type VolunteerModel struct {
	ID           string
	FirstName    string
	LastName     string
	Neighborhood string
	City         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type SocialActionModel struct {
	ID           string
	Name         string
	Organizer    string
	Description  string
	StreetLine   string
	StreetNumber string
	Neighborhood string
	City         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type SocialActionVolunteerModel struct {
	ID             string
	SocialActionID string
	FirstName      string
	LastName       string
	Neighborhood   string
	City           string
}
