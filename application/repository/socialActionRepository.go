package repository

import (
	"context"
	"go-social-action/domain/entity"
)

type SocialActionRepository interface {
	Create(ctx context.Context, socialAction *entity.SocialAction) error
	Update(ctx context.Context, socialAction *entity.SocialAction) error
	FindByID(ctx context.Context, ID string) (*entity.SocialAction, error)
	Delete(ctx context.Context, ID string) error
}
