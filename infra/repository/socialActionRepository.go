package repository

import (
	"context"
	"database/sql"

	"go-social-action/application/appError"
	"go-social-action/domain/entity"
)

type SocialActionRepositoryPostgres struct {
	DB *sql.DB
}

func NewSocialActionRepositoryPostgres(db *sql.DB) *SocialActionRepositoryPostgres {
	return &SocialActionRepositoryPostgres{
		DB: db,
	}
}

func (r *SocialActionRepositoryPostgres) Create(ctx context.Context, socialAction *entity.SocialAction) error {
	_, err := r.DB.Exec(
		`INSERT INTO social_actions (id, name, organizer, description, street_line, street_number, neighborhood, city, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		socialAction.ID, socialAction.Name, socialAction.Organizer, socialAction.Description,
		socialAction.Address.StreetLine, socialAction.Address.StreetNumber, socialAction.Address.Neighborhood,
		socialAction.Address.City, socialAction.CreatedAt, socialAction.UpdatedAt,
	)
	if err != nil {
		return err
	}
	if len(socialAction.SocialActionVolunteer) > 0 {
		for _, v := range socialAction.SocialActionVolunteer {
			_, err := r.DB.Exec(
				`INSERT INTO social_actions_volunteers (id, social_action_id, first_name, last_name, neighborhood, city)
				VALUES ($1, $2, $3, $4, $5, $6)`, v.ID, socialAction.ID, v.FirstName, v.LastName, v.Neighborhood, v.City,
			)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *SocialActionRepositoryPostgres) Delete(ctx context.Context, ID string) error {
	_, err := r.DB.Exec("DELETE FROM social_actions WHERE id = $1", ID)
	if err == sql.ErrNoRows {
		return appError.NotFoundError{Message: "social action not found"}
	}
	return err
}

func (r *SocialActionRepositoryPostgres) Update(ctx context.Context, socialAction *entity.SocialAction) error {
	_, err := r.DB.Exec(
		`UPDATE social_actions SET name = $1, organizer = $2, description = $3, street_line = $4,
		street_number = $5, neighborhood = $6, city = $7, updated_at = $8 WHERE id = $9`,
		socialAction.Name, socialAction.Organizer, socialAction.Description, socialAction.Address.StreetLine, socialAction.Address.StreetNumber,
		socialAction.Address.Neighborhood, socialAction.Address.City, socialAction.UpdatedAt, socialAction.ID,
	)
	return err
}

func (r *SocialActionRepositoryPostgres) FindByID(ctx context.Context, ID string) (*entity.SocialAction, error) {
	var socialActionModel SocialActionModel
	row := r.DB.QueryRow(`SELECT * FROM social_actions WHERE id = $1`, ID)
	if err := row.Scan(&socialActionModel.ID, &socialActionModel.Name, &socialActionModel.Organizer,
		&socialActionModel.Description, &socialActionModel.StreetLine, &socialActionModel.StreetNumber,
		&socialActionModel.Neighborhood, &socialActionModel.City, &socialActionModel.CreatedAt, &socialActionModel.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, appError.NotFoundError{Message: "social action not found"}
		}
		return nil, err
	}
	rows, err := r.DB.Query("SELECT * FROM social_actions_volunteers WHERE social_action_id = $1", socialActionModel.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var socialActionVolunteersModel []SocialActionVolunteerModel
	for rows.Next() {
		var socialActionVolunteerModel SocialActionVolunteerModel
		err := rows.Scan(&socialActionVolunteerModel.ID, &socialActionVolunteerModel.SocialActionID, &socialActionVolunteerModel.FirstName,
			&socialActionVolunteerModel.LastName, &socialActionVolunteerModel.Neighborhood, &socialActionVolunteerModel.City)
		if err != nil {
			return nil, err
		}
		socialActionVolunteersModel = append(socialActionVolunteersModel, socialActionVolunteerModel)
	}
	if rows.Err() != nil {
		return nil, err
	}
	socialActionsVolunteers := make([]*entity.SocialActionVolunteer, len(socialActionVolunteersModel))
	for i := 0; i < len(socialActionsVolunteers); i++ {
		socialActionsVolunteers[i] = entity.NewSocialActionVolunteer(
			socialActionVolunteersModel[i].ID, socialActionVolunteersModel[i].SocialActionID,
			socialActionVolunteersModel[i].FirstName, socialActionVolunteersModel[i].LastName,
			socialActionVolunteersModel[i].Neighborhood, socialActionVolunteersModel[i].City,
		)
	}
	socialActionAddress := entity.NewAddress(socialActionModel.StreetLine, socialActionModel.StreetNumber, socialActionModel.Neighborhood, socialActionModel.City)
	socialAction := entity.NewSocialAction(
		socialActionModel.ID, socialActionModel.Name, socialActionModel.Organizer,
		socialActionModel.Description, socialActionAddress,
		socialActionModel.CreatedAt, socialActionModel.UpdatedAt,
	)
	socialAction.AddSocialActionVolunteers(socialActionsVolunteers)
	return socialAction, nil
}

func (r *SocialActionRepositoryPostgres) FindAll(ctx context.Context) ([]*entity.SocialAction, error) {
	rows, err := r.DB.Query("SELECT * FROM social_actions;")
	if err != nil {
		return nil, err
	}
	var socialActions = make([]*entity.SocialAction, 0)
	for rows.Next() {
		var socialActionModel SocialActionModel
		if err := rows.Scan(&socialActionModel.ID, &socialActionModel.Name, &socialActionModel.Organizer,
			&socialActionModel.Description, &socialActionModel.StreetLine, &socialActionModel.StreetNumber,
			&socialActionModel.Neighborhood, &socialActionModel.City, &socialActionModel.CreatedAt, &socialActionModel.UpdatedAt); err != nil {
			return nil, err
		}
		rowsSocialActionVolunteers, err := r.DB.Query("SELECT * FROM social_actions_volunteers WHERE social_action_id = $1", socialActionModel.ID)
		if err != nil {
			return nil, err
		}
		defer rowsSocialActionVolunteers.Close()
		var socialActionVolunteersModel []SocialActionVolunteerModel
		for rowsSocialActionVolunteers.Next() {
			var socialActionVolunteerModel SocialActionVolunteerModel
			err := rowsSocialActionVolunteers.Scan(&socialActionVolunteerModel.ID, &socialActionVolunteerModel.SocialActionID, &socialActionVolunteerModel.FirstName,
				&socialActionVolunteerModel.LastName, &socialActionVolunteerModel.Neighborhood, &socialActionVolunteerModel.City)
			if err != nil {
				return nil, err
			}
			socialActionVolunteersModel = append(socialActionVolunteersModel, socialActionVolunteerModel)
		}
		if rowsSocialActionVolunteers.Err() != nil {
			return nil, err
		}
		socialActionsVolunteers := make([]*entity.SocialActionVolunteer, len(socialActionVolunteersModel))
		for i := 0; i < len(socialActionsVolunteers); i++ {
			socialActionsVolunteers[i] = entity.NewSocialActionVolunteer(
				socialActionVolunteersModel[i].ID, socialActionVolunteersModel[i].SocialActionID,
				socialActionVolunteersModel[i].FirstName, socialActionVolunteersModel[i].LastName,
				socialActionVolunteersModel[i].Neighborhood, socialActionVolunteersModel[i].City,
			)
		}
		socialActionAddress := entity.NewAddress(socialActionModel.StreetLine, socialActionModel.StreetNumber, socialActionModel.Neighborhood, socialActionModel.City)
		socialAction := entity.NewSocialAction(
			socialActionModel.ID, socialActionModel.Name, socialActionModel.Organizer,
			socialActionModel.Description, socialActionAddress,
			socialActionModel.CreatedAt, socialActionModel.UpdatedAt,
		)
		socialAction.AddSocialActionVolunteers(socialActionsVolunteers)
		socialActions = append(socialActions, socialAction)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return socialActions, nil
}
