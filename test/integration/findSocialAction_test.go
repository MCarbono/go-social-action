package integration

import (
	"context"
	"go-social-action/application/appError"
	"go-social-action/application/usecase"
	"go-social-action/domain/entity"
	"go-social-action/infra/database"
	"go-social-action/infra/repository"
	"go-social-action/test/assets/fakes"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestFindSocialAction(t *testing.T) {
	config := database.NewPostgresConfig("localhost", "5432", "go-social-action", "go", "go-social-action", "disable")
	db, err := database.Open(config)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	volunteerRepository := repository.NewVolunteerRepositoryPostgres(db)
	idGenerator := fakes.NewIDGeneratorFake()
	socialActionRepository := repository.NewSocialActionRepositoryPostgres(db)
	createSocialActionUseCase := usecase.NewCreateSocialActionUseCase(volunteerRepository, socialActionRepository, idGenerator)
	createVolunteerUseCase := usecase.NewCreateVolunteerUseCase(volunteerRepository, idGenerator)
	findSocialActionUseCase := usecase.NewFindSocialActionUseCase(socialActionRepository)

	type args struct {
		ctx context.Context
		ID  string
	}
	type test struct {
		name   string
		args   args
		want   *entity.SocialAction
		assert func(want, got *entity.SocialAction, err error)
	}
	tests := []test{
		{
			name: "Should find a social action",
			args: args{
				ctx: context.Background(),
				ID:  idGenerator.Generate(),
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
				SocialActionVolunteer: []*entity.SocialActionVolunteer{},
			},
			assert: func(want, got *entity.SocialAction, err error) {
				if diff := cmp.Diff(want, got, cmpopts.IgnoreFields(entity.SocialAction{}, "CreatedAt", "UpdatedAt")); diff != "" {
					t.Errorf("Create social action mismatch (-want +got):\n%v", diff)
				}
			},
		},
		{
			name: "Should not find a social action",
			args: args{
				ctx: context.Background(),
				ID:  "notFoundID",
			},
			want: nil,
			assert: func(want, got *entity.SocialAction, err error) {
				if got != nil {
					t.Errorf("FindSocialAction failed. Should returned an error, but got: %v", got)
					return
				}
				wantErr := appError.NotFoundError{Message: "social action not found"}
				if diff := cmp.Diff(err, wantErr); diff != "" {
					t.Errorf("FindSocialAction error mismatch (-err +wantErr):\n%v", diff)
				}
			},
		},
	}
	for _, scenario := range tests {
		t.Run(scenario.name, func(t *testing.T) {
			defer db.Exec("DELETE FROM volunteers;")
			defer db.Exec("DELETE FROM social_actions_volunteers;")
			defer db.Exec("DELETE FROM social_actions;")
			_, err = createVolunteerUseCase.Execute(context.Background(), &usecase.CreateVolunteerInput{
				FirstName:    "fakeFirstName1",
				LastName:     "fakeLastName1",
				Neighborhood: "fakeNeighborhood1",
				City:         "fakeCity1",
			})
			if err != nil {
				t.Fatal(err)
			}
			_, err = createSocialActionUseCase.Execute(context.Background(), &usecase.CreateSocialActionInput{
				Name:         "fake social action name",
				Organizer:    "fake organizer",
				Description:  "fake description",
				StreetLine:   "fake street line",
				StreetNumber: "fake street number",
				Neighborhood: "fake neighborhood",
				City:         "fake city",
			})
			if err != nil {
				t.Fatal(err)
			}
			got, err := findSocialActionUseCase.Execute(scenario.args.ctx, scenario.args.ID)
			scenario.assert(scenario.want, got, err)
		})
	}
}
