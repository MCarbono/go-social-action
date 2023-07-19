package usecase

import (
	"context"
	"go-social-action/application/repository"
	"go-social-action/domain/entity"
)

type FindSocialActionsUseCase struct {
	socialActionRepository repository.SocialActionRepository
}

func NewFindSocialActionsUseCase(
	socialActionRepository repository.SocialActionRepository,
) *FindSocialActionsUseCase {
	return &FindSocialActionsUseCase{
		socialActionRepository: socialActionRepository,
	}
}

func (uc *FindSocialActionsUseCase) Execute(ctx context.Context) ([]*entity.SocialAction, error) {
	socialAction, err := uc.socialActionRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return socialAction, err
}
