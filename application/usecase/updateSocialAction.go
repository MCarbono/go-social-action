package usecase

import (
	"context"
	"go-social-action/application/repository"
	"go-social-action/domain/entity"
)

type UpdateSocialActionUseCase struct {
	socialActionRepository repository.SocialActionRepository
	volunteerRepository    repository.VolunteerRepository
}

func NewUpdateSocialActionUseCase(
	socialActionRepository repository.SocialActionRepository,
	volunteerRepository repository.VolunteerRepository,
) *UpdateSocialActionUseCase {
	return &UpdateSocialActionUseCase{
		socialActionRepository: socialActionRepository,
		volunteerRepository:    volunteerRepository,
	}
}

func (uc *UpdateSocialActionUseCase) Execute(ctx context.Context, ID string, input *UpdateSocialActionInput) error {
	socialAction, err := uc.socialActionRepository.FindByID(ctx, ID)
	if err != nil {
		return err
	}
	if input.Name != "" {
		socialAction.UpdateName(input.Name)
	}
	if input.Organizer != "" {
		socialAction.UpdateOrganizer(input.Organizer)
	}
	if input.Description != "" {
		socialAction.UpdateDescription(input.Description)
	}
	if input.StreetLine != "" {
		socialAction.UpdateStreetLine(input.StreetLine)
	}
	if input.StreetNumber != "" {
		socialAction.UpdateStreetNumber(input.StreetNumber)
	}
	if input.Neighborhood != "" {
		socialAction.UpdateNeighborhood(input.Neighborhood)
	}
	if input.City != "" {
		socialAction.UpdateCity(input.City)
	}
	if len(input.SocialActionsVolunteers) > 0 {
		for i := 0; i < len(input.SocialActionsVolunteers); i++ {
			socialActionVolunteer := socialAction.FindVolunteer(input.SocialActionsVolunteers[i].ID)
			if socialActionVolunteer != nil {
				socialActionVolunteer.Update(
					input.SocialActionsVolunteers[i].FirstName, input.SocialActionsVolunteers[i].LastName,
					input.SocialActionsVolunteers[i].Neighborhood, input.SocialActionsVolunteers[i].City)
				continue
			}
			volunteer, err := uc.volunteerRepository.FindByID(ctx, input.SocialActionsVolunteers[i].ID)
			if err != nil {
				return err
			}
			socialActionVolunteer = entity.NewSocialActionVolunteer(
				volunteer.ID, socialAction.ID, volunteer.FirstName,
				volunteer.LastName, volunteer.Neighborhood, volunteer.City,
			)
			socialAction.AddSocialActionVolunteer(socialActionVolunteer)
		}
	}
	socialAction.Updated()
	err = uc.socialActionRepository.Update(ctx, socialAction)
	if err != nil {
		return err
	}
	return nil
}

type UpdateSocialActionInput struct {
	Name                    string                              `json:"name"`
	Organizer               string                              `json:"organizer"`
	Description             string                              `json:"description"`
	StreetLine              string                              `json:"street_line"`
	StreetNumber            string                              `json:"street_number"`
	Neighborhood            string                              `json:"neighborhood"`
	City                    string                              `json:"city"`
	SocialActionsVolunteers []UpdateSocialActionVolunteersInput `json:"social_action_volunteers"`
}

type UpdateSocialActionVolunteersInput struct {
	ID           string `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
}

// CREATE TABLE social_actions_volunteers (
//     id TEXT,
//     social_action_id TEXT REFERENCES social_actions (id) ON DELETE CASCADE,
//     first_name TEXT NOT NULL,
//     last_name TEXT NOT NULL,
//     neighborhood TEXT,
//     city TEXT
// );
