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
	socialAction, err := uc.socialActionRepository.FindByID(ctx, ID)
	if err != nil {
		return nil, err
	}
	return socialAction, nil
}
