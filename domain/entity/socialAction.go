package entity

import "time"

type SocialAction struct {
	ID                    string
	Name                  string
	Organizer             string
	Description           string
	Address               *Address
	SocialActionVolunteer []*SocialActionVolunteer
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

type Address struct {
	StreetLine   string
	StreetNumber string
	Neighborhood string
	City         string
}

type SocialActionVolunteer struct {
	ID             string
	SocialActionID string
	FirstName      string
	LastName       string
	Neighborhood   string
	City           string
}

func NewAddress(streetLine, streetNumber, neighborhood, city string) *Address {
	return &Address{
		StreetLine:   streetLine,
		StreetNumber: streetNumber,
		Neighborhood: neighborhood,
		City:         city,
	}
}

func NewSocialActionVolunteer(ID, socialActionID, firstName, lastName, neighborhood, city string) *SocialActionVolunteer {
	return &SocialActionVolunteer{
		ID:             ID,
		SocialActionID: socialActionID,
		FirstName:      firstName,
		LastName:       lastName,
		Neighborhood:   neighborhood,
		City:           city,
	}
}

func NewSocialAction(ID, name, organizer, description string, address *Address, createdAt, updatedAt time.Time) *SocialAction {
	return &SocialAction{
		ID:                    ID,
		Name:                  name,
		Organizer:             organizer,
		Description:           description,
		Address:               address,
		CreatedAt:             createdAt,
		UpdatedAt:             updatedAt,
		SocialActionVolunteer: []*SocialActionVolunteer{},
	}
}

func (s *SocialAction) AddSocialActionVolunteers(volunteers []*SocialActionVolunteer) {
	s.SocialActionVolunteer = volunteers
}

func (s *SocialAction) UpdateName(name string) {
	s.Name = name
	s.updated()
}

func (s *SocialAction) UpdateOrganizer(organizer string) {
	s.Organizer = organizer
	s.updated()
}

func (s *SocialAction) UpdateDescription(description string) {
	s.Description = description
	s.updated()
}

func (s *SocialAction) UpdateStreetLine(streetLine string) {
	s.Address.StreetLine = streetLine
	s.updated()
}

func (s *SocialAction) UpdateStreetNumber(streetNumber string) {
	s.Address.StreetNumber = streetNumber
	s.updated()
}

func (s *SocialAction) UpdateNeighborhood(neighborhood string) {
	s.Address.Neighborhood = neighborhood
	s.updated()
}

func (s *SocialAction) UpdateCity(city string) {
	s.Address.City = city
	s.updated()
}

func (s *SocialAction) updated() {
	s.UpdatedAt = time.Now().UTC()
}
