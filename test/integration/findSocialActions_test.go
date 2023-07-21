package integration

import (
	"context"
	"go-social-action/application/usecase"
	"go-social-action/configs"
	"go-social-action/domain/entity"
	"go-social-action/idGenerator"
	"go-social-action/infra/database"
	"go-social-action/infra/repository"
	"go-social-action/test/assets/fakes"
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestFindSocialActions(t *testing.T) {
	ctx := context.Background()
	cfg, err := configs.LoadEnvConfig("../../.env")
	if err != nil {
		log.Fatal(err)
	}
	db, err := database.Open(cfg.PSQL)
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
	idGeneratorFake := fakes.NewIDGeneratorFake()
	createVolunteerUseCase := usecase.NewCreateVolunteerUseCase(volunteerRepository, idGenerator)
	createSocialActionUseCase := usecase.NewCreateSocialActionUseCase(volunteerRepository, socialActionRepository, idGeneratorFake)
	findSocialActionsUseCase := usecase.NewFindSocialActionsUseCase(socialActionRepository)
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
	_, err = createSocialActionUseCase.Execute(ctx, &usecase.CreateSocialActionInput{
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
	want := []*entity.SocialAction{
		{
			ID:          "fakeUUID",
			Name:        "fake social action name",
			Organizer:   "fake organizer",
			Description: "fake description",
			Address: &entity.Address{
				StreetLine:   "fake street line",
				StreetNumber: "fake street number",
				Neighborhood: "fake neighborhood",
				City:         "fake city",
			},
			SocialActionVolunteer: []*entity.SocialActionVolunteer{
				{
					SocialActionID: "fakeUUID",
					FirstName:      "fakeFirstName1",
					LastName:       "fakeLastName1",
					Neighborhood:   "fakeNeighborhood1",
					City:           "fakeCity1",
				},
				{
					SocialActionID: "fakeUUID",
					FirstName:      "fakeFirstName2",
					LastName:       "fakeLastName2",
					Neighborhood:   "fakeNeighborhood2",
					City:           "fakeCity2",
				},
			},
		},
	}
	got, err := findSocialActionsUseCase.Execute(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got,
		cmpopts.IgnoreFields(entity.SocialAction{}, "CreatedAt", "UpdatedAt"),
		cmpopts.IgnoreFields(entity.SocialActionVolunteer{}, "ID"),
	); diff != "" {
		t.Errorf("Find social actions mismatch (-want +got):\n%v", diff)
		return
	}
}
