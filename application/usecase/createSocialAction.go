package usecase

import (
	"context"
	"go-social-action/application/repository"
	"go-social-action/domain/entity"
	"go-social-action/idGenerator"
	"time"
)

type CreateSocialActionUseCase struct {
	volunteerRepository    repository.VolunteerRepository
	socialActionRepository repository.SocialActionRepository
	idGenerator            idGenerator.IDGenerator
}

func NewCreateSocialActionUseCase(
	volunteerRepository repository.VolunteerRepository,
	socialActionRepository repository.SocialActionRepository,
	idGenerator idGenerator.IDGenerator,
) *CreateSocialActionUseCase {
	return &CreateSocialActionUseCase{
		volunteerRepository:    volunteerRepository,
		socialActionRepository: socialActionRepository,
		idGenerator:            idGenerator,
	}
}

func (uc *CreateSocialActionUseCase) Execute(ctx context.Context, input *CreateSocialActionInput) (*entity.SocialAction, error) {
	socialActionAddress := entity.NewAddress(input.StreetLine, input.StreetNumber, input.Neighborhood, input.City)
	socialAction := entity.NewSocialAction(
		uc.idGenerator.Generate(), input.Name, input.Organizer,
		input.Description, socialActionAddress, time.Now().UTC(), time.Now().UTC(),
	)
	if len(input.SocialActionsVolunteers) > 0 {
		volunteers, err := uc.volunteerRepository.Find(ctx, input.SocialActionsVolunteers)
		if err != nil {
			return nil, err
		}
		socialActionVolunteers := make([]*entity.SocialActionVolunteer, len(volunteers))
		for i := range socialActionVolunteers {
			socialActionVolunteers[i] = entity.NewSocialActionVolunteer(
				volunteers[i].ID, socialAction.ID, volunteers[i].FirstName,
				volunteers[i].LastName, volunteers[i].Neighborhood, volunteers[i].City,
			)
		}
		socialAction.AddSocialActionVolunteers(socialActionVolunteers)
	}
	if err := uc.socialActionRepository.Create(ctx, socialAction); err != nil {
		return nil, err
	}
	return socialAction, nil
}

type CreateSocialActionInput struct {
	Name                    string   `json:"name"`
	Organizer               string   `json:"organizer"`
	Description             string   `json:"description"`
	StreetLine              string   `json:"street_line"`
	StreetNumber            string   `json:"street_number"`
	Neighborhood            string   `json:"neighborhood"`
	City                    string   `json:"city"`
	SocialActionsVolunteers []string `json:"social_action_volunteers"`
}
