package integration

import (
	"context"
	"go-social-action/application/appError"
	"go-social-action/application/usecase"
	"go-social-action/idGenerator"
	"go-social-action/infra/database"
	"go-social-action/infra/repository"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDeleteSocialAction(t *testing.T) {
	ctx := context.Background()
	config := database.NewPostgresConfig("localhost", "5432", "go-social-action", "go", "go-social-action", "disable")
	db, err := database.Open(config)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	defer db.Exec("DELETE FROM volunteers;")
	defer db.Exec("DELETE FROM social_actions_volunteers;")
	defer db.Exec("DELETE FROM social_actions;")
	volunteerRepository := repository.NewVolunteerRepositoryPostgres(db)
	socialActionRepository := repository.NewSocialActionRepositoryPostgres(db)
	idGenerator := idGenerator.New()
	createVolunteerUseCase := usecase.NewCreateVolunteerUseCase(volunteerRepository, idGenerator)
	createSocialActionUseCase := usecase.NewCreateSocialActionUseCase(volunteerRepository, socialActionRepository, idGenerator)
	deleteSocialActionUseCase := usecase.NewDeleteSocialActionUseCase(socialActionRepository)
	findSocialActionUseCase := usecase.NewFindSocialActionUseCase(socialActionRepository)
	firstVolunteer, err := createVolunteerUseCase.Execute(ctx, &usecase.CreateVolunteerInput{
		FirstName:    "fakeFirstName1",
		LastName:     "fakeLastName1",
		Neighborhood: "fakeNeighborhood1",
		City:         "fakeCity1",
	})
	if err != nil {
		t.Fatal(err)
	}
	secondVolunteer, err := createVolunteerUseCase.Execute(ctx, &usecase.CreateVolunteerInput{
		FirstName:    "fakeFirstName2",
		LastName:     "fakeLastName2",
		Neighborhood: "fakeNeighborhood2",
		City:         "fakeCity2",
	})
	if err != nil {
		t.Fatal(err)
	}
	socialAction, err := createSocialActionUseCase.Execute(ctx, &usecase.CreateSocialActionInput{
		Name:                    "fake social action name",
		Organizer:               "fake organizer",
		Description:             "fake description",
		StreetLine:              "fake street line",
		StreetNumber:            "fake street number",
		Neighborhood:            "fake neighborhood",
		City:                    "fake city",
		SocialActionsVolunteers: []string{firstVolunteer.ID, secondVolunteer.ID},
	})
	if err != nil {
		t.Fatal(err)
	}
	err = deleteSocialActionUseCase.Execute(ctx, socialAction.ID)
	if err != nil {
		t.Fatal(err)
	}
	wantErr := appError.NotFoundError{Message: "social action not found"}
	socialActionDeleted, err := findSocialActionUseCase.Execute(ctx, socialAction.ID)
	if socialActionDeleted != nil {
		t.Errorf("TestDeleteSocialAction failed. Should not returned a social action, but got: %v", socialActionDeleted)
		return
	}
	if diff := cmp.Diff(err, wantErr); diff != "" {
		t.Errorf("Delete social action mismatch (-err +wantErr):\n%v", diff)
		return
	}
	rows, err := db.Query("SELECT * FROM social_actions_volunteers WHERE id IN ($1, $2)", firstVolunteer.ID, secondVolunteer.ID)
	if err != nil {
		t.Fatal(err)
	}
	var socialActionVolunteers []repository.SocialActionVolunteerModel
	for rows.Next() {
		var model repository.SocialActionVolunteerModel
		err := rows.Scan(&model.ID, &model.SocialActionID, &model.FirstName, &model.LastName, &model.Neighborhood, &model.City)
		if err != nil {
			t.Fatal(err)
		}
		socialActionVolunteers = append(socialActionVolunteers, model)
	}
	if rows.Err() != nil {
		t.Fatal(err)
	}
	if len(socialActionVolunteers) > 0 {
		t.Errorf("TestDeleteSocialAction failed. Should not return any volunteer from the deleted social action, but got %v", socialActionVolunteers)
	}
}
