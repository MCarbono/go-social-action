package usecase

import (
	"context"
	"go-social-action/application/repository"
	"go-social-action/domain/entity"
)

type FindVolunteerUseCase struct {
	volunteerRepository repository.VolunteerRepository
}

func NewFindVolunteerUseCase(
	volunteerRepository repository.VolunteerRepository,
) *FindVolunteerUseCase {
	return &FindVolunteerUseCase{
		volunteerRepository: volunteerRepository,
	}
}

func (uc *FindVolunteerUseCase) Execute(ctx context.Context, ID string) (*entity.Volunteer, error) {
	return uc.volunteerRepository.FindByID(ctx, ID)
}
