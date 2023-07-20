package entity

import (
	"time"
)

type SocialAction struct {
	ID                    string                   `json:"id"`
	Name                  string                   `json:"name"`
	Organizer             string                   `json:"organizer"`
	Description           string                   `json:"description"`
	Address               *Address                 `json:"address"`
	SocialActionVolunteer []*SocialActionVolunteer `json:"social_action_volunteers"`
	CreatedAt             time.Time                `json:"created_at"`
	UpdatedAt             time.Time                `json:"updated_at"`
}

type Address struct {
	StreetLine   string `json:"street_line"`
	StreetNumber string `json:"street_number"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
}

type SocialActionVolunteer struct {
	ID             string `json:"id"`
	SocialActionID string `json:"social_action_id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Neighborhood   string `json:"neighborhoor"`
	City           string `json:"city"`
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

func (s *SocialAction) AddSocialActionVolunteer(volunteer *SocialActionVolunteer) {
	s.SocialActionVolunteer = append(s.SocialActionVolunteer, volunteer)
}

func (s *SocialAction) UpdateName(name string) {
	s.Name = name
}

func (s *SocialAction) UpdateOrganizer(organizer string) {
	s.Organizer = organizer
}

func (s *SocialAction) UpdateDescription(description string) {
	s.Description = description
}

func (s *SocialAction) UpdateStreetLine(streetLine string) {
	s.Address.StreetLine = streetLine
}

func (s *SocialAction) UpdateStreetNumber(streetNumber string) {
	s.Address.StreetNumber = streetNumber
}

func (s *SocialAction) UpdateNeighborhood(neighborhood string) {
	s.Address.Neighborhood = neighborhood
}

func (s *SocialAction) UpdateCity(city string) {
	s.Address.City = city
}

func (s *SocialAction) Updated() {
	s.UpdatedAt = time.Now().UTC()
}

func (s *SocialAction) FindVolunteer(ID string) *SocialActionVolunteer {
	for _, v := range s.SocialActionVolunteer {
		if v.ID == ID {
			return v
		}
	}
	return nil
}

func (s *SocialActionVolunteer) Update(firstName, lastName, neighborhood, city string) {
	if firstName != "" {
		s.FirstName = firstName
	}
	if lastName != "" {
		s.LastName = lastName
	}
	if neighborhood != "" {
		s.Neighborhood = neighborhood
	}
	if city != "" {
		s.City = city
	}
}
