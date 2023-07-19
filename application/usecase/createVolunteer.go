package usecase

import (
	"context"
	"go-social-action/application/repository"
	"go-social-action/domain/entity"
	"go-social-action/idGenerator"
	"time"
)

type CreateVolunteerUseCase struct {
	volunteerRepository repository.VolunteerRepository
	idGenerator         idGenerator.IDGenerator
}

func NewCreateVolunteerUseCase(volunteerRepository repository.VolunteerRepository, idGenerator idGenerator.IDGenerator) *CreateVolunteerUseCase {
	return &CreateVolunteerUseCase{
		volunteerRepository: volunteerRepository,
		idGenerator:         idGenerator,
	}
}

func (uc *CreateVolunteerUseCase) Execute(ctx context.Context, input *CreateVolunteerInput) (*entity.Volunteer, error) {
	volunteer := entity.NewVolunteer(uc.idGenerator.Generate(), input.FirstName, input.LastName, input.Neighborhood, input.City, time.Now().UTC(), time.Now().UTC())
	if err := uc.volunteerRepository.Create(ctx, volunteer); err != nil {
		return nil, err
	}
	return volunteer, nil
}

type CreateVolunteerInput struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
}
