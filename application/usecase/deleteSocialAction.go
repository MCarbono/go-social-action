package usecase

import (
	"context"
	"go-social-action/application/repository"
)

type DeleteSocialActionUseCase struct {
	socialActionRepository repository.SocialActionRepository
}

func NewDeleteSocialActionUseCase(
	socialActionRepository repository.SocialActionRepository,
) *DeleteSocialActionUseCase {
	return &DeleteSocialActionUseCase{
		socialActionRepository: socialActionRepository,
	}
}

func (uc *DeleteSocialActionUseCase) Execute(ctx context.Context, ID string) error {
	return uc.socialActionRepository.Delete(ctx, ID)
}
