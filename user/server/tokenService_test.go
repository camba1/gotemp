package main

import (
	"github.com/golang/protobuf/ptypes"
	pb "goTemp/user/proto"
	"reflect"
	"strings"
	"testing"
)

//TestTokenService_Encode: Tests both encode and decode functions
func TestTokenService_Encode(t *testing.T) {
	type args struct {
		user *pb.User
	}

	users := getUsers()
	myUser := users.User[0]

	tests := []struct {
		name    string
		args    args
		want    args
		wantErr bool
	}{
		{name: "Test encode", args: args{user: myUser}, want: args{user: myUser}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &TokenService{}
			got, err := ts.Encode(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == "" {
				t.Errorf("Encode() got = %v, want %v", got, tt.want)
			}
			if len(strings.Split(got, ".")) != 3 {
				t.Errorf("Encode() got = %v, want %v", got, tt.want)
			}
			decoded, err := ts.Decode(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if decoded == nil {
				t.Errorf("Decode() got nil when trying to decode token %v", got)
			} else {
				if !reflect.DeepEqual(decoded.User, tt.want.user) {
					t.Errorf("Decode() got = %v, want %v", decoded, tt.want.user)
				}
			}
		})
	}
}

func Test_getKeyFromVault(t *testing.T) {
	tests := []struct {
		name string
		//want    int
		wantErr bool
	}{
		{name: "Get Key", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getKeyFromVault()
			if (err != nil) != tt.wantErr {
				t.Errorf("getKeyFromVault() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(got) == "" {
				t.Errorf("getKeyFromVault() got an empty string (%v), want non empty string", got)
			}
		})
	}
}

func getUsers() *pb.Users {
	currentTime := ptypes.TimestampNow()
	users := pb.Users{}
	userOne := pb.User{
		Id:         1234,
		Firstname:  "Super",
		Lastname:   "Duck",
		ValidFrom:  currentTime,
		ValidThru:  currentTime,
		Active:     true,
		Name:       "Super Duck",
		Email:      "duck@mymail.com",
		Company:    "Ducks Inc",
		Createdate: currentTime,
		Updatedate: currentTime,
	}
	users.User = append(users.User, &userOne)
	return &users
}
