package main

import (
	"context"
	"fmt"
	"goTemp/audit/server/statements"
	"goTemp/globalUtils"
	"reflect"
	"testing"
	"time"
)

func TestAuditSrv_GetAudit(t *testing.T) {
	type args struct {
		ctx      context.Context
		searchId *globalUtils.AuditSearchId
	}

	ctx := context.Background()

	validIdArg := args{
		ctx:      ctx,
		searchId: &globalUtils.AuditSearchId{Id: 2330739450113430529},
	}
	inValidIdArg := args{
		ctx:      ctx,
		searchId: &globalUtils.AuditSearchId{Id: 12345},
	}
	conn = connectToDB()
	defer conn.Close(context.Background())

	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{name: "Valid Id", args: validIdArg, want: validIdArg.searchId.Id, wantErr: false},
		{name: "Invalid Id", args: inValidIdArg, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AuditSrv{}
			dbRowStruct, dbMessage, err := a.GetAudit(tt.args.ctx, tt.args.searchId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAudit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want == 0 {
				return
			}
			if dbRowStruct == nil {
				t.Error("GetAudit() dbRowStruct is  nil")
				return
			}
			if dbRowStruct.Id != tt.want {
				t.Errorf("GetAudit() dbRowStruct = %v, want %v", dbRowStruct.Id, tt.want)
				return
			}
			if dbMessage == nil {
				t.Error("GetAudit() dbMessage is nil")
				return
			}
		})
	}
}

func TestAuditSrv_GetAudits(t *testing.T) {
	type args struct {
		ctx         context.Context
		searchParms *globalUtils.AuditSearchParams
	}

	ctx := context.Background()

	conn = connectToDB()
	defer conn.Close(context.Background())

	dtForSearch, _ := time.Parse(globalUtils.DateLayoutISO, "2000-01-01")
	dtForSearchEnd, _ := time.Parse(globalUtils.DateLayoutISO, "2021-07-26")
	data := getSearchParmsData(dtForSearch, dtForSearchEnd)

	emptySearch := data["emptySearch"]
	fullSearch := data["fullSearch"]
	onlyDateSearch := data["onlyDateSearch"]
	onlyObjectIdSearch := data["onlyObjectIdSearch"]

	tests := []struct {
		name    string
		args    args
		want    int
		wantid  int64
		wantErr bool
	}{
		{name: "Empty search", args: args{ctx, emptySearch}, want: 0, wantid: 0, wantErr: true},
		{name: "Full search", args: args{ctx, fullSearch}, want: 1, wantid: fullSearch.ObjectId, wantErr: false},
		{name: "Only date search", args: args{ctx, onlyDateSearch}, want: 0, wantid: 0, wantErr: true},
		{name: "Only object Id search", args: args{ctx, onlyObjectIdSearch}, want: 0, wantid: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AuditSrv{}
			got, err := a.GetAudits(tt.args.ctx, tt.args.searchParms)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAudits() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				if got == nil {
					t.Errorf("GetAudits() got is nil want %d", tt.want)
					return
				}
				if len(got.Header) != tt.want {
					t.Errorf("GetAudits() got = %v, want %v", got, tt.want)
				}
				if got.Header[0].ObjectId != tt.wantid {
					t.Errorf("GetAudits() object id got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestAuditSrv_getSQLForSearch(t *testing.T) {
	type args struct {
		searchParms *globalUtils.AuditSearchParams
	}

	sql := statements.SqlSelectAll.String()

	dtForSearch, _ := time.Parse(globalUtils.DateLayoutISO, "2021-06-26")
	dtForSearchEnd, _ := time.Parse(globalUtils.DateLayoutISO, "2021-07-26")

	sqlEmptyWhereClause := " where 1=1"
	sqlFullSearch := sqlEmptyWhereClause + " AND audit.objectname = $1 AND audit.objectid = $2 AND audit.actiontime >= $3 AND audit.actiontime <= $4"
	sqlOnlyDateSearch := sqlEmptyWhereClause + " AND audit.actiontime >= $1 AND audit.actiontime <= $2"
	sqlOnlyObjectIdSearch := sqlEmptyWhereClause + " AND audit.objectid = $1"

	sqlEmptyWhereClauseFinal := fmt.Sprintf(sql, sqlEmptyWhereClause, statements.MaxRowsToFetch)
	sqlFullSearchFinal := fmt.Sprintf(sql, sqlFullSearch, statements.MaxRowsToFetch)
	sqlOnlyDateSearchFinal := fmt.Sprintf(sql, sqlOnlyDateSearch, statements.MaxRowsToFetch)
	sqlOnlyObjectIdSearchFinal := fmt.Sprintf(sql, sqlOnlyObjectIdSearch, statements.MaxRowsToFetch)

	var intEmptySearch []interface{}
	intFullSearch := []interface{}{"user", int64(123456789), dtForSearch, dtForSearchEnd}
	intOnlyDateSearch := []interface{}{dtForSearch, dtForSearchEnd}
	intOnlyObjectIdSearch := []interface{}{int64(123456789)}

	data := getSearchParmsData(dtForSearch, dtForSearchEnd)

	emptySearch := data["emptySearch"]
	fullSearch := data["fullSearch"]
	onlyDateSearch := data["onlyDateSearch"]
	onlyObjectIdSearch := data["onlyObjectIdSearch"]

	tests := []struct {
		name       string
		args       args
		want       string
		wantValues []interface{}
		wantErr    bool
	}{
		{name: "Empty search", args: args{emptySearch}, want: sqlEmptyWhereClauseFinal, wantValues: intEmptySearch, wantErr: false},
		{name: "Full search", args: args{fullSearch}, want: sqlFullSearchFinal, wantValues: intFullSearch, wantErr: false},
		{name: "Only date search", args: args{onlyDateSearch}, want: sqlOnlyDateSearchFinal, wantValues: intOnlyDateSearch, wantErr: false},
		{name: "Only object Id search", args: args{onlyObjectIdSearch}, want: sqlOnlyObjectIdSearchFinal, wantValues: intOnlyObjectIdSearch, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AuditSrv{}
			got, got1, err := a.getSQLForSearch(tt.args.searchParms)
			if (err != nil) != tt.wantErr {
				t.Errorf("getSQLForSearch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got1 != tt.want {
				t.Errorf("getSQLForSearch() got1 = %v, want %v", got1, tt.want)
			}
			if !reflect.DeepEqual(got, tt.wantValues) {
				t.Errorf("getSQLForSearch() got = %v, want %v", got, tt.wantValues)
			}

		})
	}
}

func TestAuditSrv_validateSearchParams(t *testing.T) {
	type args struct {
		searchParms *globalUtils.AuditSearchParams
	}

	dtForSearch, _ := time.Parse(globalUtils.DateLayoutISO, "2021-06-26")
	dtForSearchEnd, _ := time.Parse(globalUtils.DateLayoutISO, "2021-07-26")
	data := getSearchParmsData(dtForSearch, dtForSearchEnd)

	emptySearch := data["emptySearch"]
	fullSearch := data["fullSearch"]
	onlyDateSearch := data["onlyDateSearch"]
	onlyObjectIdSearch := data["onlyObjectIdSearch"]

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "Empty search", args: args{emptySearch}, want: 5, wantErr: true},
		{name: "Full search", args: args{fullSearch}, want: 0, wantErr: false},
		{name: "Only date search", args: args{onlyDateSearch}, want: 2, wantErr: true},
		{name: "Only object Id search", args: args{onlyObjectIdSearch}, want: 4, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AuditSrv{}
			got, err := a.validateSearchParams(tt.args.searchParms)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateSearchParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("validateSearchParams() got = %v validations, want %v validations", got, tt.want)
			}
		})
	}
}

func getSearchParmsData(dtForSearch time.Time, dtForSearchEnd time.Time) map[string]*globalUtils.AuditSearchParams {
	data := make(map[string]*globalUtils.AuditSearchParams)
	data["emptySearch"] = &globalUtils.AuditSearchParams{}
	data["fullSearch"] = &globalUtils.AuditSearchParams{
		ObjectName:      "user",
		ObjectId:        123456789,
		ActionTimeStart: dtForSearch,
		ActionTimeEnd:   dtForSearchEnd,
	}
	data["onlyDateSearch"] = &globalUtils.AuditSearchParams{
		ActionTimeStart: dtForSearch,
		ActionTimeEnd:   dtForSearchEnd,
	}
	data["onlyObjectIdSearch"] = &globalUtils.AuditSearchParams{ObjectId: 123456789}

	return data
}
