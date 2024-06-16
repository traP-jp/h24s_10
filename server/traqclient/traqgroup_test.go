package traqclient

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/traP-jp/h24s_10/model"
	"github.com/traPtitech/go-traq"
)

func TestClient_CreateUserGroup(t *testing.T) {
	type fields struct {
		apiClient *traq.APIClient
	}
	type args struct {
		ctx          context.Context
		name         string
		description  string
		groupType    string
		adminID      string
		participants []model.Participant
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Group
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test Case 1",
			fields: fields{
				apiClient: traq.NewAPIClient(traq.NewConfiguration()),
			},
			args: args{
				ctx:          context.TODO(),
				name:         "Group 1",
				description:  "Test group 1",
				groupType:    "public",
				adminID:      "Luftalian",
				participants: []model.Participant{},
			},
			want:    Group{},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			fields: fields{
				apiClient: traq.NewAPIClient(traq.NewConfiguration()),
			},
			args: args{
				ctx:         context.TODO(),
				name:        "Group 2",
				description: "Test group 2",
				groupType:   "private",
				adminID:     "Luftalian",
				participants: []model.Participant{
					{
						ID:      uuid.UUID{},
						EventID: uuid.UUID{},
						TraQID:  "Luftalian",
						TeamID:  0,
					},
					{
						ID:      uuid.UUID{},
						EventID: uuid.UUID{},
						TraQID:  "pirosiki",
						TeamID:  0,
					},
				},
			},
			want: Group{
				Id:          "group2",
				Name:        "Group 2",
				Description: "Test group 2",
				Type:        "private",
				Admins:      []string{"Luftalian"},
				Members: []UserGroupMember{
					{
						Id:   "pirosiki",
						Role: "admin",
					},
					{
						Id:   "Luftalian",
						Role: "member",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				apiClient: tt.fields.apiClient,
			}
			got, err := c.CreateUserGroup(tt.args.ctx, tt.args.name, tt.args.description, tt.args.groupType, tt.args.adminID, tt.args.participants)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateUserGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.CreateUserGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
