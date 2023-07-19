package usecase

import (
	"context"
	"go-social-action/application/repository"
	"go-social-action/idGenerator"
)

type UpdateSocialActionUseCase struct {
	socialActionRepository repository.SocialActionRepository
	idGenerator            idGenerator.IDGenerator
}

func NewUpdateSocialActionUseCase(
	socialActionRepository repository.SocialActionRepository,
	idGenerator idGenerator.IDGenerator,
) *UpdateSocialActionUseCase {
	return &UpdateSocialActionUseCase{
		socialActionRepository: socialActionRepository,
		idGenerator:            idGenerator,
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
	err = uc.socialActionRepository.Update(ctx, socialAction)
	if err != nil {
		return err
	}
	return nil
}

type UpdateSocialActionInput struct {
	Name                    string   `json:"name"`
	Organizer               string   `json:"organizer"`
	Description             string   `json:"description"`
	StreetLine              string   `json:"street_line"`
	StreetNumber            string   `json:"street_number"`
	Neighborhood            string   `json:"neighborhood"`
	City                    string   `json:"city"`
	SocialActionsVolunteers []string `json:"social_action_volunteers"`
}
