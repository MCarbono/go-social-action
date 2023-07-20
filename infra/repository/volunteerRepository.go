package repository

import (
	"context"
	"database/sql"
	"fmt"
	"go-social-action/application/appError"
	"go-social-action/domain/entity"
	"strings"
)

type VolunteerRepositoryPostgres struct {
	DB *sql.DB
}

func NewVolunteerRepositoryPostgres(db *sql.DB) *VolunteerRepositoryPostgres {
	return &VolunteerRepositoryPostgres{
		DB: db,
	}
}

func (r *VolunteerRepositoryPostgres) Create(ctx context.Context, volunteer *entity.Volunteer) error {
	_, err := r.DB.Exec(`INSERT INTO volunteers (id, first_name, last_name, neighborhood, city, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		volunteer.ID, volunteer.FirstName, volunteer.LastName, volunteer.Neighborhood, volunteer.City, volunteer.CreatedAt, volunteer.UpdatedAt)
	return err
}

func (r *VolunteerRepositoryPostgres) FindByID(ctx context.Context, ID string) (*entity.Volunteer, error) {
	var model VolunteerModel
	row := r.DB.QueryRow(`SELECT * FROM volunteers WHERE id = $1`, ID)
	if err := row.Scan(&model.ID, &model.FirstName, &model.LastName, &model.Neighborhood, &model.City, &model.CreatedAt, &model.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, appError.NotFoundError{Message: "volunteer not found"}
		}
		return nil, err
	}
	volunteer := entity.NewVolunteer(model.ID, model.FirstName, model.LastName, model.Neighborhood, model.City, model.CreatedAt, model.UpdatedAt)
	return volunteer, nil
}

func (r *VolunteerRepositoryPostgres) Find(ctx context.Context, IDS []string) ([]*entity.Volunteer, error) {
	placeholders := make([]string, len(IDS))
	for i := range IDS {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}
	query := fmt.Sprintf("SELECT * FROM volunteers WHERE id IN (%s)", strings.Join(placeholders, ", "))
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	args := make([]interface{}, len(IDS))
	for i, id := range IDS {
		args[i] = id
	}
	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var volunteersModel []VolunteerModel
	for rows.Next() {
		var model VolunteerModel
		err := rows.Scan(&model.ID, &model.FirstName, &model.LastName, &model.Neighborhood, &model.City, &model.CreatedAt, &model.UpdatedAt)
		if err != nil {
			return nil, err
		}
		volunteersModel = append(volunteersModel, model)
	}
	if rows.Err() != nil {
		return nil, err
	}
	volunteers := make([]*entity.Volunteer, len(volunteersModel))
	for i := 0; i < len(volunteers); i++ {
		volunteers[i] = entity.NewVolunteer(volunteersModel[i].ID, volunteersModel[i].FirstName, volunteersModel[i].LastName, volunteersModel[i].Neighborhood, volunteersModel[i].City, volunteersModel[i].CreatedAt, volunteersModel[i].UpdatedAt)
	}
	return volunteers, nil
}
