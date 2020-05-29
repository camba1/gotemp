package main

import (
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
	sqlTestSearch := sqlEmptyWhereClause + " AND appuser.id = $1 AND appuser.firstname = $2 AND appuser.lastname = $3 AND appuser.validfrom <= $4 AND appuser.validthru >= $4"

	var intEmptySearch []interface{}
	intFullSearch := []interface{}{int64(1), "Super", "Duck", convertedDates[0]}
	intOnlyDateSearch := []interface{}{convertedDates[0]}
	intOnlyFirstNameSearch := []interface{}{"Super"}
	intTestSearch := []interface{}{int64(1234), "Incredible", "Green Guy", convertedDates[0]}

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
