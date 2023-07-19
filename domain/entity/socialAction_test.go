package entity

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewSocialAction(t *testing.T) {
	type args struct {
		ID          string
		name        string
		organizer   string
		description string
		address     *Address
		createdAt   time.Time
		updatedAt   time.Time
	}
	type test struct {
		name string
		args args
		want *SocialAction
	}
	tests := []test{
		{
			name: "Should create a new social action",
			args: args{
				ID:          "fakeID",
				name:        "social action test",
				organizer:   "social action organizer test",
				description: "description social action",
				address:     NewAddress("rua de teste", "10", "bairro de teste", "cidade de teste"),
				createdAt:   time.Now().UTC(),
				updatedAt:   time.Now().UTC(),
			},
			want: &SocialAction{
				ID:                    "fakeID",
				Name:                  "social action test",
				Organizer:             "social action organizer test",
				Description:           "description social action",
				Address:               NewAddress("rua de teste", "10", "bairro de teste", "cidade de teste"),
				SocialActionVolunteer: []*SocialActionVolunteer{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSocialAction(tt.args.ID, tt.args.name, tt.args.organizer, tt.args.description, tt.args.address, tt.args.createdAt, tt.args.updatedAt)
			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(SocialAction{}, "CreatedAt", "UpdatedAt")); diff != "" {
				t.Errorf("Create Social Action mismatch (-want +got):\n%v", diff)
			}
		})
	}
}

func TestSocialAction_UpdateName(t *testing.T) {
	newName := "updatedName"
	want := NewSocialAction(
		"fakeID",
		newName,
		"social action organizer test",
		"description social action",
		NewAddress("rua de teste", "10", "bairro de teste", "cidade de teste"),
		time.Now().UTC(),
		time.Now().UTC(),
	)
	got := NewSocialAction(
		"fakeID",
		"social action test",
		"social action organizer test",
		"description social action",
		NewAddress("rua de teste", "10", "bairro de teste", "cidade de teste"),
		time.Now().UTC(),
		time.Now().UTC(),
	)
	got.UpdateName(newName)
	if diff := cmp.Diff(want, got, cmpopts.IgnoreFields(SocialAction{}, "CreatedAt", "UpdatedAt")); diff != "" {
		t.Errorf("Updated Social Action name mismatch (-want +got):\n%v", diff)
	}
}

func TestSocialAction_UpdateOrganizer(t *testing.T) {
	newOrganizer := "organizerUpdated"
	want := NewSocialAction(
		"fakeID",
		"social action test",
		newOrganizer,
		"description social action",
		NewAddress("rua de teste", "10", "bairro de teste", "cidade de teste"),
		time.Now().UTC(),
		time.Now().UTC(),
	)
	got := NewSocialAction(
		"fakeID",
		"social action test",
		"social action organizer test",
		"description social action",
		NewAddress("rua de teste", "10", "bairro de teste", "cidade de teste"),
		time.Now().UTC(),
		time.Now().UTC(),
	)
	got.UpdateOrganizer(newOrganizer)
	if diff := cmp.Diff(want, got, cmpopts.IgnoreFields(SocialAction{}, "CreatedAt", "UpdatedAt")); diff != "" {
		t.Errorf("Updated Social Action name mismatch (-want +got):\n%v", diff)
	}
}

func TestSocialAction_UpdateDescription(t *testing.T) {
	newDescription := "descriptionUpdated"
	want := NewSocialAction(
		"fakeID",
		"social action test",
		"social action organizer test",
		newDescription,
		NewAddress("rua de teste", "10", "bairro de teste", "cidade de teste"),
		time.Now().UTC(),
		time.Now().UTC(),
	)
	got := NewSocialAction(
		"fakeID",
		"social action test",
		"social action organizer test",
		"description social action",
		NewAddress("rua de teste", "10", "bairro de teste", "cidade de teste"),
		time.Now().UTC(),
		time.Now().UTC(),
	)
	got.UpdateDescription(newDescription)
	if diff := cmp.Diff(want, got, cmpopts.IgnoreFields(SocialAction{}, "CreatedAt", "UpdatedAt")); diff != "" {
		t.Errorf("Updated Social Action name mismatch (-want +got):\n%v", diff)
	}
}

func TestSocialAction_UpdateStreetLine(t *testing.T) {
	newStreetLine := "streetLine updated"
	want := NewSocialAction(
		"fakeID",
		"social action test",
		"social action organizer test",
		"description social action",
		NewAddress(newStreetLine, "10", "bairro de teste", "cidade de teste"),
		time.Now().UTC(),
		time.Now().UTC(),
	)
	got := NewSocialAction(
		"fakeID",
		"social action test",
		"social action organizer test",
		"description social action",
		NewAddress("rua de teste", "10", "bairro de teste", "cidade de teste"),
		time.Now().UTC(),
		time.Now().UTC(),
	)
	got.UpdateStreetLine(newStreetLine)
	if diff := cmp.Diff(want, got, cmpopts.IgnoreFields(SocialAction{}, "CreatedAt", "UpdatedAt")); diff != "" {
		t.Errorf("Updated Social Action name mismatch (-want +got):\n%v", diff)
	}
}

func TestSocialAction_UpdateStreetNumber(t *testing.T) {
	newStreetNumber := "1"
	want := NewSocialAction(
		"fakeID",
		"social action test",
		"social action organizer test",
		"description social action",
		NewAddress("rua de teste", newStreetNumber, "bairro de teste", "cidade de teste"),
		time.Now().UTC(),
		time.Now().UTC(),
	)
	got := NewSocialAction(
		"fakeID",
		"social action test",
		"social action organizer test",
		"description social action",
		NewAddress("rua de teste", "10", "bairro de teste", "cidade de teste"),
		time.Now().UTC(),
		time.Now().UTC(),
	)
	got.UpdateStreetNumber(newStreetNumber)
	if diff := cmp.Diff(want, got, cmpopts.IgnoreFields(SocialAction{}, "CreatedAt", "UpdatedAt")); diff != "" {
		t.Errorf("Updated Social Action name mismatch (-want +got):\n%v", diff)
	}
}

func TestSocialAction_UpdateNeighborhood(t *testing.T) {
	newNeighborhood := "neighborhood updated"
	want := NewSocialAction(
		"fakeID",
		"social action test",
		"social action organizer test",
		"description social action",
		NewAddress("rua de teste", "10", newNeighborhood, "cidade de teste"),
		time.Now().UTC(),
		time.Now().UTC(),
	)
	got := NewSocialAction(
		"fakeID",
		"social action test",
		"social action organizer test",
		"description social action",
		NewAddress("rua de teste", "10", "bairro de teste", "cidade de teste"),
		time.Now().UTC(),
		time.Now().UTC(),
	)
	got.UpdateNeighborhood(newNeighborhood)
	if diff := cmp.Diff(want, got, cmpopts.IgnoreFields(SocialAction{}, "CreatedAt", "UpdatedAt")); diff != "" {
		t.Errorf("Updated Social Action name mismatch (-want +got):\n%v", diff)
	}
}

func TestSocialAction_UpdateCity(t *testing.T) {
	newCity := "city updated"
	want := NewSocialAction(
		"fakeID",
		"social action test",
		"social action organizer test",
		"description social action",
		NewAddress("rua de teste", "10", "bairro de teste", newCity),
		time.Now().UTC(),
		time.Now().UTC(),
	)
	got := NewSocialAction(
		"fakeID",
		"social action test",
		"social action organizer test",
		"description social action",
		NewAddress("rua de teste", "10", "bairro de teste", "cidade de teste"),
		time.Now().UTC(),
		time.Now().UTC(),
	)
	got.UpdateCity(newCity)
	if diff := cmp.Diff(want, got, cmpopts.IgnoreFields(SocialAction{}, "CreatedAt", "UpdatedAt")); diff != "" {
		t.Errorf("Updated Social Action name mismatch (-want +got):\n%v", diff)
	}
}
