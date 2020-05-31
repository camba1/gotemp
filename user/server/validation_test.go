package main

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"goTemp/globalUtils"
	"goTemp/globalerrors"
	pb "goTemp/user/proto"
	"testing"
)

func TestUser_BeforeCreateUser(t *testing.T) {
	type args struct {
		ctx           context.Context
		user          *pb.User
		validationErr *pb.ValidationErr
	}

	ctx := context.Background()
	testUsers := getTestUsers()

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Empty User", args: args{ctx: ctx, user: testUsers["emptyUser"], validationErr: &pb.ValidationErr{FailureDesc: []string{}}}, wantErr: true},
		{name: "Good User", args: args{ctx: ctx, user: testUsers["goodUser"], validationErr: &pb.ValidationErr{FailureDesc: []string{}}}, wantErr: false},
		{name: "No names User", args: args{ctx: ctx, user: testUsers["noNamesUser"], validationErr: &pb.ValidationErr{FailureDesc: []string{}}}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{}
			err := u.BeforeCreateUser(tt.args.ctx, tt.args.user, tt.args.validationErr)
			if (err != nil) != tt.wantErr {
				t.Errorf("BeforeCreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err, ok := err.(*globalerrors.ValidationError); err != nil && tt.wantErr && !ok {
				t.Errorf("BeforeCreateUser() expected ValidationError but got error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_BeforeDeleteUser(t *testing.T) {
	type args struct {
		ctx           context.Context
		user          *pb.User
		validationErr *pb.ValidationErr
	}
	ctx := context.Background()
	activeUser := pb.User{
		Active: true,
	}
	inactiveUser := pb.User{
		Active: false,
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Active User", args: args{ctx: ctx, user: &activeUser, validationErr: &pb.ValidationErr{FailureDesc: []string{}}}, wantErr: true},
		{name: "Inactive User", args: args{ctx: ctx, user: &inactiveUser, validationErr: &pb.ValidationErr{FailureDesc: []string{}}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{}
			err := u.BeforeDeleteUser(tt.args.ctx, tt.args.user, tt.args.validationErr)
			if (err != nil) != tt.wantErr {
				t.Errorf("BeforeDeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err, ok := err.(*globalerrors.ValidationError); err != nil && tt.wantErr && !ok {
				t.Errorf("BeforeDeleteUser() expected ValidationError but got error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_BeforeUpdateUser(t *testing.T) {
	type args struct {
		ctx           context.Context
		user          *pb.User
		validationErr *pb.ValidationErr
	}
	ctx := context.Background()
	testUsers := getTestUsers()
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{

		{name: "Empty User", args: args{ctx: ctx, user: testUsers["emptyUser"], validationErr: &pb.ValidationErr{FailureDesc: []string{}}}, wantErr: true},
		{name: "Good User", args: args{ctx: ctx, user: testUsers["goodUser"], validationErr: &pb.ValidationErr{FailureDesc: []string{}}}, wantErr: false},
		{name: "No names User", args: args{ctx: ctx, user: testUsers["noNamesUser"], validationErr: &pb.ValidationErr{FailureDesc: []string{}}}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{}
			err := u.BeforeUpdateUser(tt.args.ctx, tt.args.user, tt.args.validationErr)
			if (err != nil) != tt.wantErr {
				t.Errorf("BeforeUpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err, ok := err.(*globalerrors.ValidationError); err != nil && tt.wantErr && !ok {
				t.Errorf("BeforeUpdateUser() expected ValidationError but got error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_checkMandatoryFields(t *testing.T) {
	type args struct {
		user *pb.User
	}

	testUsers := getTestUsers()

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "Empty User", args: args{testUsers["emptyUser"]}, want: 6, wantErr: false},
		{name: "Good User", args: args{testUsers["goodUser"]}, want: 0, wantErr: false},
		{name: "No Names", args: args{testUsers["noNamesUser"]}, want: 2, wantErr: false},
		{name: "Bad dates ", args: args{testUsers["badDatesUser"]}, want: 1, wantErr: false},
		{name: "No Password", args: args{testUsers["noPwdUser"]}, want: 1, wantErr: false},
		{name: "No Email", args: args{testUsers["noEmailUser"]}, want: 1, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkMandatoryFields(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkMandatoryFields() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("checkMandatoryFields() got %v failed validations, want %v", got, tt.want)
			}
		})
	}
}

func getTestUsers() map[string]*pb.User {
	nextYear := globalUtils.GetNextYearTimeStamp()
	thisYear := ptypes.TimestampNow()

	emptyUser := pb.User{}
	goodUser := pb.User{
		Id:        12344,
		Firstname: "Super",
		Lastname:  "Duck",
		ValidFrom: thisYear,
		ValidThru: nextYear,
		Active:    false,
		Pwd:       "xxxx",
		Email:     "duck@duck.com",
		Company:   "Duck R Us",
		Name:      "Super Duck",
	}
	noNamesUser := pb.User{
		Id:        12345,
		ValidFrom: thisYear,
		ValidThru: nextYear,
		Active:    false,
		Pwd:       "xxxx",
		Email:     "duck@duck.com",
		Company:   "Duck R Us",
		Name:      "Super Duck",
	}
	badDatesUser := pb.User{
		Id:        12344,
		Firstname: "Super",
		Lastname:  "Duck",
		ValidFrom: nextYear,
		ValidThru: thisYear,
		Active:    false,
		Pwd:       "xxxx",
		Email:     "duck@duck.com",
		Company:   "Duck R Us",
		Name:      "Super Duck",
	}
	noPwdUser := pb.User{
		Id:        12344,
		Firstname: "Super",
		Lastname:  "Duck",
		ValidFrom: thisYear,
		ValidThru: nextYear,
		Active:    false,
		Email:     "duck@duck.com",
		Company:   "Duck R Us",
		Name:      "Super Duck",
	}
	noEmailUser := pb.User{
		Id:        12344,
		Firstname: "Super",
		Lastname:  "Duck",
		ValidFrom: thisYear,
		ValidThru: nextYear,
		Active:    false,
		Pwd:       "xxxx",
		Company:   "Duck R Us",
		Name:      "Super Duck",
	}

	return map[string]*pb.User{
		"emptyUser":    &emptyUser,
		"goodUser":     &goodUser,
		"noNamesUser":  &noNamesUser,
		"badDatesUser": &badDatesUser,
		"noPwdUser":    &noPwdUser,
		"noEmailUser":  &noEmailUser,
	}
}

func Test_checkEmailUnique(t *testing.T) {
	type args struct {
		user     *pb.User
		users    *pb.Users
		isInsert bool
	}

	testUsers := getTestUsers()

	insertErrTxt := userErr.InsertDupEmail()
	updateErrTxt := userErr.UpdateDupEmail()

	user := testUsers["goodUser"]
	emptyUsers := &pb.Users{User: nil}
	oneUser := &pb.Users{User: []*pb.User{
		testUsers["goodUser"],
	},
	}
	oneUserDiffId := &pb.Users{User: []*pb.User{
		testUsers["noNamesUser"],
	},
	}
	twoUsers := &pb.Users{User: []*pb.User{
		testUsers["goodUser"],
		testUsers["noNamesUser"],
	},
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Nil user on insert", args: args{user: user, users: nil, isInsert: true}, want: ""},
		{name: "No user on insert", args: args{user: user, users: emptyUsers, isInsert: true}, want: ""},
		{name: "One user on insert", args: args{user: user, users: oneUser, isInsert: true}, want: insertErrTxt},
		{name: "Nil user on update", args: args{user: user, users: nil, isInsert: false}, want: ""},
		{name: "No user on update", args: args{user: user, users: emptyUsers, isInsert: false}, want: ""},
		{name: "One user on update", args: args{user: user, users: oneUser, isInsert: false}, want: ""},
		{name: "One user, incorrect id, on update", args: args{user: user, users: oneUserDiffId, isInsert: false}, want: updateErrTxt},
		{name: "Two users on update", args: args{user: user, users: twoUsers, isInsert: false}, want: updateErrTxt},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkEmailUnique(tt.args.user, tt.args.users, tt.args.isInsert)
			if got != tt.want {
				t.Errorf("checkEmailUnique() got = %v, want %v", got, tt.want)
			}
		})
	}
}
