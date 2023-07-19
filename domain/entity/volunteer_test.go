package entity

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewVolunteer(t *testing.T) {
	type args struct {
		ID           string
		firstName    string
		lastName     string
		neighborhood string
		city         string
		createdAt    time.Time
		updatedAt    time.Time
	}
	type test struct {
		name string
		args args
		want *Volunteer
	}
	tests := []test{
		{
			name: "Should create a new volunteer",
			args: args{
				ID:           "fakeID",
				firstName:    "Teste",
				lastName:     "da Silva",
				neighborhood: "bairro teste",
				city:         "Testelandia",
				createdAt:    time.Now().UTC(),
				updatedAt:    time.Now().UTC(),
			},
			want: &Volunteer{
				ID:           "fakeID",
				FirstName:    "Teste",
				LastName:     "da Silva",
				Neighborhood: "bairro teste",
				City:         "Testelandia",
				CreatedAt:    time.Now().UTC(),
				UpdatedAt:    time.Now().UTC(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewVolunteer(tt.args.ID, tt.args.firstName, tt.args.lastName, tt.args.neighborhood, tt.args.city, tt.args.createdAt, tt.args.updatedAt)
			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(Volunteer{}, "CreatedAt", "UpdatedAt")); diff != "" {
				t.Errorf("Create Volunteer mismatch (-want +got):\n%v", diff)
			}
		})
	}
}
