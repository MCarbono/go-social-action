package repository

import (
	"context"
	"go-social-action/domain/entity"
)

type VolunteerRepository interface {
	Create(ctx context.Context, volunteer *entity.Volunteer) error
	FindByID(ctx context.Context, ID string) (*entity.Volunteer, error)
	Find(ctx context.Context, IDS []string) ([]*entity.Volunteer, error)
}
