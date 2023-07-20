package usecase

import (
	"context"
	"go-social-action/application/repository"
	"go-social-action/domain/entity"
)

type FindSocialActionUseCase struct {
	socialActionRepository repository.SocialActionRepository
}

func NewFindSocialActionUseCase(
	socialActionRepository repository.SocialActionRepository,
) *FindSocialActionUseCase {
	return &FindSocialActionUseCase{
		socialActionRepository: socialActionRepository,
	}
}

func (uc *FindSocialActionUseCase) Execute(ctx context.Context, ID string) (*entity.SocialAction, error) {
	return uc.socialActionRepository.FindByID(ctx, ID)
}
