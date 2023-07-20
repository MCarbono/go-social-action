package integration

import (
	"context"
	"go-social-action/application/usecase"
	"go-social-action/domain/entity"
	"go-social-action/infra/database"
	"go-social-action/infra/repository"
	"go-social-action/test/assets/fakes"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestUpdateSocialAction(t *testing.T) {
	config := database.NewPostgresConfig("localhost", "5432", "go-social-action", "go", "go-social-action", "disable")
	db, err := database.Open(config)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	volunteerRepository := repository.NewVolunteerRepositoryPostgres(db)
	socialActionRepository := repository.NewSocialActionRepositoryPostgres(db)
	idGenerator := fakes.NewIDGeneratorFake()
	createVolunteerUseCase := usecase.NewCreateVolunteerUseCase(volunteerRepository, idGenerator)
	createSocialActionUseCase := usecase.NewCreateSocialActionUseCase(volunteerRepository, socialActionRepository, idGenerator)
	updateSocialActionUseCase := usecase.NewUpdateSocialActionUseCase(socialActionRepository, volunteerRepository)
	findSocialActionUseCase := usecase.NewFindSocialActionUseCase(socialActionRepository)
	volunteer, err := createVolunteerUseCase.Execute(context.Background(), &usecase.CreateVolunteerInput{
		FirstName:    "Teste",
		LastName:     "da Silva",
		Neighborhood: "Bairro de teste",
		City:         "Testelandia",
	})
	if err != nil {
		t.Fatal(err)
	}
	defer db.Exec("DELETE FROM volunteers;")
	type args struct {
		ctx   context.Context
		input *usecase.UpdateSocialActionInput
		*usecase.CreateSocialActionInput
	}
	type test struct {
		name string
		args args
		want *entity.SocialAction
	}
	tests := []test{
		{
			name: "Should update name, organizer, description, streetLine, streetNumber, neighborhood and city from a social action",
			args: args{
				ctx: context.Background(),
				input: &usecase.UpdateSocialActionInput{
					Name:         "updated name",
					Organizer:    "updated organizer",
					Description:  "updated description",
					StreetLine:   "updated street line",
					StreetNumber: "updated street number",
					Neighborhood: "updated neighborhood",
					City:         "updated city",
				},
				CreateSocialActionInput: &usecase.CreateSocialActionInput{
					Name:         "fake social action name",
					Organizer:    "fake organizer",
					Description:  "fake description",
					StreetLine:   "fake street line",
					StreetNumber: "fake street number",
					Neighborhood: "fake neighborhood",
					City:         "fake city",
				},
			},
			want: &entity.SocialAction{
				ID:          "fakeUUID",
				Name:        "updated name",
				Organizer:   "updated organizer",
				Description: "updated description",
				Address: &entity.Address{
					StreetLine:   "updated street line",
					StreetNumber: "updated street number",
					Neighborhood: "updated neighborhood",
					City:         "updated city",
				},
				SocialActionVolunteer: []*entity.SocialActionVolunteer{},
			},
		},
		{
			name: "Should update firstName, lastName, neighborhood and city of a social action volunteer",
			args: args{
				ctx: context.Background(),
				input: &usecase.UpdateSocialActionInput{
					SocialActionsVolunteers: []usecase.UpdateSocialActionVolunteersInput{
						{
							ID:           volunteer.ID,
							FirstName:    "updated volunteer first name",
							LastName:     "updated volunteer last name",
							Neighborhood: "updated volunteer neighborhood",
							City:         "updated volunteer city",
						},
					},
				},
				CreateSocialActionInput: &usecase.CreateSocialActionInput{
					Name:                    "fake social action name",
					Organizer:               "fake organizer",
					Description:             "fake description",
					StreetLine:              "fake street line",
					StreetNumber:            "fake street number",
					Neighborhood:            "fake neighborhood",
					City:                    "fake city",
					SocialActionsVolunteers: []string{volunteer.ID},
				},
			},
			want: &entity.SocialAction{
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
						ID:             "fakeUUID",
						SocialActionID: "fakeUUID",
						FirstName:      "updated volunteer first name",
						LastName:       "updated volunteer last name",
						Neighborhood:   "updated volunteer neighborhood",
						City:           "updated volunteer city",
					},
				},
			},
		},
		{
			name: "Should create a new volunteer to a social action already created",
			args: args{
				ctx: context.Background(),
				input: &usecase.UpdateSocialActionInput{
					SocialActionsVolunteers: []usecase.UpdateSocialActionVolunteersInput{
						{
							ID: volunteer.ID,
						},
					},
				},
				CreateSocialActionInput: &usecase.CreateSocialActionInput{
					Name:         "fake social action name",
					Organizer:    "fake organizer",
					Description:  "fake description",
					StreetLine:   "fake street line",
					StreetNumber: "fake street number",
					Neighborhood: "fake neighborhood",
					City:         "fake city",
				},
			},
			want: &entity.SocialAction{
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
						ID:             "fakeUUID",
						SocialActionID: "fakeUUID",
						FirstName:      "Teste",
						LastName:       "da Silva",
						Neighborhood:   "Bairro de teste",
						City:           "Testelandia",
					},
				},
			},
		},
	}
	for _, scenario := range tests {
		t.Run(scenario.name, func(t *testing.T) {
			defer db.Exec("DELETE FROM social_actions_volunteers;")
			defer db.Exec("DELETE FROM social_actions;")
			socialActionCreated, err := createSocialActionUseCase.Execute(scenario.args.ctx, scenario.args.CreateSocialActionInput)
			if err != nil {
				t.Fatal(err)
			}
			err = updateSocialActionUseCase.Execute(scenario.args.ctx, socialActionCreated.ID, scenario.args.input)
			if err != nil {
				t.Fatal(err)
			}
			got, err := findSocialActionUseCase.Execute(scenario.args.ctx, socialActionCreated.ID)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(scenario.want, got, cmpopts.IgnoreFields(entity.SocialAction{}, "CreatedAt", "UpdatedAt")); diff != "" {
				t.Errorf("Update social action mismatch (-want +got):\n%v", diff)
			}
		})
	}
}
