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

func TestCreateVolunteer(t *testing.T) {
	config := database.NewPostgresConfig("localhost", "5432", "go-social-action", "go", "go-social-action", "disable")
	db, err := database.Open(config)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	volunteerRepository := repository.NewVolunteerRepositoryPostgres(db)
	idGenerator := idGenerator.New()
	createVolunteerUseCase := usecase.NewCreateVolunteerUseCase(volunteerRepository, idGenerator)

	type args struct {
		ctx   context.Context
		input *usecase.CreateVolunteerInput
	}
	type test struct {
		name string
		args args
	}
	tests := []test{
		{
			name: "Should create a new volunteer",
			args: args{
				ctx: context.Background(),
				input: &usecase.CreateVolunteerInput{
					FirstName:    "Teste",
					LastName:     "da Silva",
					Neighborhood: "Bairro de teste",
					City:         "Testelandia",
				},
			},
		},
	}
	for _, scenario := range tests {
		t.Run(scenario.name, func(t *testing.T) {
			defer db.Exec("DELETE FROM volunteers")
			got, err := createVolunteerUseCase.Execute(scenario.args.ctx, scenario.args.input)
			if err != nil {
				t.Fatal(err)
			}
			want, err := volunteerRepository.FindByID(scenario.args.ctx, got.ID)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("Create volunteer mismatch (-want +got):\n%v", diff)
			}
		})
	}
}
