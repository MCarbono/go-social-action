package integration

import (
	"context"
	"go-social-action/application/usecase"
	"go-social-action/idGenerator"
	"go-social-action/infra/database"
	"go-social-action/infra/repository"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreateSocialAction(t *testing.T) {
	config := database.NewPostgresConfig("localhost", "5432", "go-social-action", "go", "go-social-action", "disable")
	db, err := database.Open(config)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	volunteerRepository := repository.NewVolunteerRepositoryPostgres(db)
	idGenerator := idGenerator.New()
	createVolunteerUseCase := usecase.NewCreateVolunteerUseCase(volunteerRepository, idGenerator)
	firstVolunteer, err := createVolunteerUseCase.Execute(context.Background(), &usecase.CreateVolunteerInput{
		FirstName:    "fakeFirstName1",
		LastName:     "fakeLastName1",
		Neighborhood: "fakeNeighborhood1",
		City:         "fakeCity1",
	})
	if err != nil {
		t.Fatal(err)
	}
	secondVolunteer, err := createVolunteerUseCase.Execute(context.Background(), &usecase.CreateVolunteerInput{
		FirstName:    "fakeFirstName2",
		LastName:     "fakeLastName2",
		Neighborhood: "fakeNeighborhood2",
		City:         "fakeCity2",
	})
	if err != nil {
		t.Fatal(err)
	}
	defer db.Exec("DELETE FROM volunteers;")
	socialActionRepository := repository.NewSocialActionRepositoryPostgres(db)
	createSocialActionUseCase := usecase.NewCreateSocialActionUseCase(volunteerRepository, socialActionRepository, idGenerator)
	findSocialActionUseCase := usecase.NewFindSocialActionUseCase(socialActionRepository)
	type args struct {
		ctx   context.Context
		input *usecase.CreateSocialActionInput
	}
	type test struct {
		name string
		args args
	}
	tests := []test{
		{
			name: "Should create a new social action with volunteers",
			args: args{
				ctx: context.Background(),
				input: &usecase.CreateSocialActionInput{
					Name:                    "fake social action name",
					Organizer:               "fake organizer",
					Description:             "fake description",
					StreetLine:              "fake street line",
					StreetNumber:            "fake street number",
					Neighborhood:            "fake neighborhood",
					City:                    "fake city",
					SocialActionsVolunteers: []string{firstVolunteer.ID, secondVolunteer.ID},
				},
			},
		},
		{
			name: "Should create a new social action without volunteers",
			args: args{
				ctx: context.Background(),
				input: &usecase.CreateSocialActionInput{
					Name:         "fake social action name",
					Organizer:    "fake organizer",
					Description:  "fake description",
					StreetLine:   "fake street line",
					StreetNumber: "fake street number",
					Neighborhood: "fake neighborhood",
					City:         "fake city",
				},
			},
		},
	}
	for _, scenario := range tests {
		t.Run(scenario.name, func(t *testing.T) {
			defer db.Exec("DELETE FROM social_actions_volunteers;")
			defer db.Exec("DELETE FROM social_actions;")
			got, err := createSocialActionUseCase.Execute(scenario.args.ctx, scenario.args.input)
			if err != nil {
				t.Fatal(err)
			}
			want, err := findSocialActionUseCase.Execute(scenario.args.ctx, got.ID)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("Create social action mismatch (-want +got):\n%v", diff)
			}
		})
	}
}
