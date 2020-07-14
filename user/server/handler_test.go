package main

import (
	"context"
	"goTemp/globalUtils"
	pb "goTemp/user/proto"
	"reflect"
	"testing"
)

func TestUser_buildSearchWhereClause(t *testing.T) {
	type args struct {
		searchParms *pb.SearchParams
	}

	_, dtForSearch := globalUtils.TimeStringToTimestamp("2021-06-26")
	convertedDates, _ := globalUtils.TimeStampPPBToTime(dtForSearch)

	sqlEmptyWhereClause := " where 1=1"
	sqlFullSearch := sqlEmptyWhereClause + " AND appuser.id = $1 AND appuser.firstname = $2 AND appuser.lastname = $3 AND appuser.validfrom <= $4 AND appuser.validthru >= $4"
	sqlOnlyDateSearch := sqlEmptyWhereClause + " AND appuser.validfrom <= $1 AND appuser.validthru >= $1"
	sqlOnlyFirstNameSearch := sqlEmptyWhereClause + " AND appuser.firstname = $1"
	sqlTestSearch := sqlEmptyWhereClause + " AND appuser.id = $1 AND appuser.firstname = $2 AND appuser.lastname = $3 AND appuser.email = $4 AND appuser.company = $5 AND appuser.validfrom <= $6 AND appuser.validthru >= $6"

	var intEmptySearch []interface{}
	intFullSearch := []interface{}{int64(1), "Super", "Duck", convertedDates[0]}
	intOnlyDateSearch := []interface{}{convertedDates[0]}
	intOnlyFirstNameSearch := []interface{}{"Super"}
	intTestSearch := []interface{}{int64(1234), "Incredible", "Green Guy", "igg@mymail.com", "igg & Associates", convertedDates[0]}

	emptySearch := pb.SearchParams{}
	fullSearch := pb.SearchParams{
		Id:        1,
		Fisrtname: "Super",
		Lastname:  "Duck",
		ValidDate: dtForSearch,
	}
	onlyDateSearch := pb.SearchParams{
		ValidDate: dtForSearch,
	}
	onlyFirstNameSearch := pb.SearchParams{Fisrtname: "Super"}
	testSearch := pb.SearchParams{
		Id:        1234,
		Fisrtname: "Incredible",
		Lastname:  "Green Guy",
		ValidDate: dtForSearch,
		Email:     "igg@mymail.com",
		Company:   "igg & Associates",
	}

	tests := []struct {
		name       string
		args       args
		want       string
		wantValues []interface{}
		wantErr    bool
	}{
		{name: "Empty search", args: args{&emptySearch}, want: sqlEmptyWhereClause, wantValues: intEmptySearch, wantErr: false},
		{name: "Full search", args: args{&fullSearch}, want: sqlFullSearch, wantValues: intFullSearch, wantErr: false},
		{name: "Only date search", args: args{&onlyDateSearch}, want: sqlOnlyDateSearch, wantValues: intOnlyDateSearch, wantErr: false},
		{name: "Only first name search", args: args{&onlyFirstNameSearch}, want: sqlOnlyFirstNameSearch, wantValues: intOnlyFirstNameSearch, wantErr: false},
		{name: "Test search", args: args{&testSearch}, want: sqlTestSearch, wantValues: intTestSearch, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{}
			got, got1, err := u.buildSearchWhereClause(tt.args.searchParms)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildSearchWhereClause() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("buildSearchWhereClause() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.wantValues) {
				t.Errorf("buildSearchWhereClause() got1 = %v, want %v", got1, tt.wantValues)
			}
		})
	}
}

func TestUser_ValidateToken(t *testing.T) {
	type args struct {
		ctx      context.Context
		inToken  *pb.Token
		outToken *pb.Token
	}
	invalidToken := args{
		ctx: context.Background(),
		inToken: &pb.Token{
			Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7InB3ZCI6IjEyMzQiLCJlbWFpbCI6ImR1Y2tAbXltYWlsLmNvbSJ9LCJleHAiOjE1OTEyMTc2NDAsImlhdCI6MTU5MTEzMTI0MCwiaXNzIjoiZ29UZW1wLnVzZXJzcnYifQ.hJ7dxx3oIronb0tPMMWX8AaW4vWPP9Mu6PlWpVKHKHk",
			Valid: false,
		},
		outToken: &pb.Token{},
	}
	validToken := args{
		ctx: context.Background(),
		inToken: &pb.Token{
			Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoyMzIyODUzODcyOTQ0NTUwOTEzLCJjb21wYW55IjoiRHVjayBJbmMuIn0sImV4cCI6MTU5MTIyMTc0MiwiaWF0IjoxNTkxMTM1MzQyLCJpc3MiOiJnb1RlbXAudXNlcnNydiJ9.Cvy4PYUd8mxub2PBggCDyYwwYsa-rx_JlHSMIdQYvKk",
			Valid: false,
		},
		outToken: &pb.Token{},
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Validate token", args: invalidToken, wantErr: true},
		{name: "Validate token", args: validToken, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{}
			err := u.ValidateToken(tt.args.ctx, tt.args.inToken, tt.args.outToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_Auth(t *testing.T) {
	type args struct {
		ctx   context.Context
		user  *pb.User
		token *pb.Token
	}

	loadConfig()
	conn = connectToDB()
	ctx := context.Background()
	goodUser := &pb.User{
		Pwd:   "1234",
		Email: "duck@mymail.com"}
	badUser := &pb.User{
		Pwd:   "0000",
		Email: "XYZ@mybadmail.com"}

	tests := []struct {
		name      string
		args      args
		wantErr   bool
		wantToken bool
	}{
		{name: "Find customer", args: args{ctx: ctx, user: goodUser, token: &pb.Token{}}, wantErr: false, wantToken: true},
		{name: "Do not find customer", args: args{ctx: ctx, user: badUser, token: &pb.Token{}}, wantErr: true, wantToken: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{}
			err := u.Auth(tt.args.ctx, tt.args.user, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("Auth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(tt.args.token.GetToken()) != 0 && !tt.wantToken {
				t.Errorf("Auth() Unexpected user token length. Wanted = 0, got %v", len(tt.args.token.GetToken()))
			}
		})
	}
}
