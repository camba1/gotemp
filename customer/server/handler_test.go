package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/micro/go-micro/v2/metadata"
	"goTemp/customer/proto"
	"goTemp/customer/server/statements"
	"goTemp/globalProtos"
	"goTemp/globalUtils"
	"log"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func Test_customer_GetCustomerById(t *testing.T) {
	type args struct {
		ctx         context.Context
		searchId    *proto.SearchId
		outCustomer *proto.Customer
	}
	loadConfig()
	conn = connectToDB()
	ctx := context.Background()
	validIdArg := args{
		ctx:         ctx,
		searchId:    &proto.SearchId{XKey: "ducksrus"},
		outCustomer: &proto.Customer{},
	}
	validIdArg2 := args{
		ctx:         ctx,
		searchId:    &proto.SearchId{XKey: "canard"},
		outCustomer: &proto.Customer{},
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "Get a customer", args: validIdArg, want: "Ducks R Us", wantErr: false},
		{name: "Get a second customer", args: validIdArg2, want: "Canard Oui Oui", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &customer{}
			if err := c.GetCustomerById(tt.args.ctx, tt.args.searchId, tt.args.outCustomer); (err != nil) != tt.wantErr {
				t.Errorf("GetCustomerById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.args.outCustomer.Name != tt.want {
				t.Errorf(" expected customer name == %s, got %s", tt.want, tt.args.outCustomer.Name)
			}
		})
	}
}

func Test_customer_GetCustomers(t *testing.T) {
	type args struct {
		ctx       context.Context
		params    *proto.SearchParams
		customers *proto.Customers
	}
	loadConfig()
	conn = connectToDB()
	ctx := context.Background()
	validIdArg := args{
		ctx:       ctx,
		params:    &proto.SearchParams{Name: "Ducks R Us"},
		customers: &proto.Customers{},
	}
	inValidIdArg := args{
		ctx:       ctx,
		params:    &proto.SearchParams{Name: "XYZ"},
		customers: &proto.Customers{},
	}

	tests := []struct {
		name      string
		args      args
		want      string
		wantCount int
		wantErr   bool
	}{
		{name: "Find one customer", args: validIdArg, want: "Ducks R Us", wantCount: 1, wantErr: false},
		{name: "No customers to be found", args: inValidIdArg, want: "", wantCount: 0, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &customer{}
			err := c.GetCustomers(tt.args.ctx, tt.args.params, tt.args.customers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCustomers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(tt.args.customers.Customer) != tt.wantCount {
				t.Errorf("GetCustomers() record count = %v, wantCount %v", len(tt.args.customers.Customer), tt.wantCount)
			}
			if len(tt.args.customers.Customer) == 1 && tt.args.customers.Customer[0].Name != tt.want {
				t.Errorf(" expected customer name == %s, got %s", tt.want, tt.args.customers.Customer[0].Name)
			}

		})
	}
}

func Test_customer_DeleteCustomer(t *testing.T) {
	type args struct {
		ctx      context.Context
		searchId *proto.SearchId
		response *proto.Response
	}

	//setup test
	loadConfig()
	conn = connectToDB()
	id := base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(123456789, 10)))
	ctx := metadata.Set(context.Background(), "userid", id)

	//create a customer
	dtValidFrom, _ := time.Parse(globalUtils.DateLayoutISO, "2021-06-26")
	dtValidThru, _ := time.Parse(globalUtils.DateLayoutISO, "2021-07-26")
	custData := getTestData(dtValidFrom, dtValidThru)
	goodCustomer := custData["goodCustomer"]
	newCustomerId, err := InsertTestCustomer(ctx, goodCustomer)
	if err != nil {
		t.Errorf("UpdateCustomer() Unable to create test customer. error = %v", err)
	}
	custToDeleteId := newCustomerId

	//try to delete  customer
	//custToDeleteId := "delme"
	validIdArg := args{
		ctx:      ctx,
		searchId: &proto.SearchId{XKey: custToDeleteId},
		response: &proto.Response{},
	}
	nothingToDelArg := args{
		ctx:      ctx,
		searchId: &proto.SearchId{XKey: "XYZ"},
		response: &proto.Response{},
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{name: "Delete a customer", args: validIdArg, want: 1, wantErr: false},
		{name: "Nothing to delete", args: nothingToDelArg, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &customer{}
			err := c.DeleteCustomer(tt.args.ctx, tt.args.searchId, tt.args.response)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.args.response.AffectedCount != tt.want {
				t.Errorf("DeleteCustomer() wanted to delete %v records, deleted %v", tt.args.response.AffectedCount, tt.want)
			}
		})
	}
}

func Test_customer_CreateCustomer(t *testing.T) {
	type args struct {
		ctx        context.Context
		inCustomer *proto.Customer
		response   *proto.Response
	}
	loadConfig()
	conn = connectToDB()

	id := base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(123456789, 10)))
	ctx := metadata.Set(context.Background(), "userid", id)

	dtValidFrom, _ := time.Parse(globalUtils.DateLayoutISO, "2021-06-26")
	dtValidThru, _ := time.Parse(globalUtils.DateLayoutISO, "2021-07-26")
	custData := getTestData(dtValidFrom, dtValidThru)
	goodCustomer := custData["goodCustomer"]
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Good Customer", args: args{ctx: ctx, inCustomer: goodCustomer, response: &proto.Response{}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &customer{}
			if err := c.CreateCustomer(tt.args.ctx, tt.args.inCustomer, tt.args.response); (err != nil) != tt.wantErr {
				t.Errorf("CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_customer_UpdateCustomer(t *testing.T) {
	type args struct {
		ctx        context.Context
		inCustomer *proto.Customer
		response   *proto.Response
	}
	loadConfig()
	conn = connectToDB()

	id := base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(123456789, 10)))
	ctx := metadata.Set(context.Background(), "userid", id)

	dtValidFrom, _ := time.Parse(globalUtils.DateLayoutISO, "2021-06-26")
	dtValidThru, _ := time.Parse(globalUtils.DateLayoutISO, "2021-07-26")
	custData := getTestData(dtValidFrom, dtValidThru)
	updateCustomer := custData["updateCustomer"]

	//create a customer
	goodCustomer := custData["goodCustomer"]
	newCustomerId, err := InsertTestCustomer(ctx, goodCustomer)
	if err != nil {
		t.Errorf("UpdateCustomer() Unable to create test customer. error = %v", err)
	}

	updateCustomer.XKey = newCustomerId

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Update Customer", args: args{ctx: ctx, inCustomer: updateCustomer, response: &proto.Response{}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &customer{}
			if err := c.UpdateCustomer(tt.args.ctx, tt.args.inCustomer, tt.args.response); (err != nil) != tt.wantErr {
				t.Errorf("UpdateCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	//Clean up
	c := &customer{}
	err = c.DeleteCustomer(ctx, &proto.SearchId{XKey: newCustomerId}, &proto.Response{})
	if err != nil {
		t.Errorf("UpdateCustomer() Unable to delete created customer %v. Error: %v", newCustomerId, err)
	}
}

func getTestData(validfrom time.Time, validThru time.Time) map[string]*proto.Customer {
	customers := make(map[string]*proto.Customer)
	validDates, err := globalUtils.TimeToTimeStampPPB(validfrom, validThru)
	if err != nil {
		log.Fatalf("Unable to convert time.time %v, %v to timestamp", validfrom, validThru)
	}
	mods := globalProtos.GlModification{
		CreateDate: validDates[0],
		UpdateDate: validDates[1],
		ModifiedBy: "12345",
	}
	vdates := globalProtos.GlValidityDate{
		ValidFrom: validDates[0],
		ValidThru: validDates[1],
	}
	goodCustomer := proto.Customer{
		Name:          "goodCustomer",
		ValidityDates: &vdates,
		Modifications: &mods,
	}
	updateCustomer := proto.Customer{
		XKey:          "",
		Name:          "goodUpdatedCustomer",
		ValidityDates: &vdates,
		Modifications: &mods,
	}
	customers["goodCustomer"] = &goodCustomer
	customers["updateCustomer"] = &updateCustomer
	return customers
}

func InsertTestCustomer(ctx context.Context, customerToInsert *proto.Customer) (string, error) {
	c := &customer{}
	var insertedResponse proto.Response
	err := c.CreateCustomer(ctx, customerToInsert, &insertedResponse)
	if err != nil {
		return "", fmt.Errorf("UpdateCustomer() unable to create customer before update error = %v", err)
	}
	return insertedResponse.Customer.XKey, nil
}

func Test_customer_getSQLForSearch(t *testing.T) {
	type args struct {
		searchParms *proto.SearchParams
	}

	sql := statements.SqlSelectAll.String()

	dtTmForSearch, _ := time.Parse(globalUtils.DateLayoutISO, "2021-06-26")
	protoDate, _ := globalUtils.TimeToTimeStampPPB(dtTmForSearch)
	dtForSearch := protoDate[0]

	sqlEmptyWhereClause := " FILTER 1==1 "
	sqlFullSearch := sqlEmptyWhereClause + " AND c._key ==  @xkey AND c.name == @name AND c.validityDates.validFrom.seconds <= @validDateSecs AND c.validityDates.validThru.seconds >= @validDateSecs"
	sqlOnlyDateSearch := sqlEmptyWhereClause + " AND c.validityDates.validFrom.seconds <= @validDateSecs AND c.validityDates.validThru.seconds >= @validDateSecs"
	sqlOnlyObjectIdSearch := sqlEmptyWhereClause + " AND c._key ==  @xkey"

	sqlEmptyWhereClauseFinal := fmt.Sprintf(sql, sqlEmptyWhereClause, statements.MaxRowsToFetch)
	sqlFullSearchFinal := fmt.Sprintf(sql, sqlFullSearch, statements.MaxRowsToFetch)
	sqlOnlyDateSearchFinal := fmt.Sprintf(sql, sqlOnlyDateSearch, statements.MaxRowsToFetch)
	sqlOnlyObjectIdSearchFinal := fmt.Sprintf(sql, sqlOnlyObjectIdSearch, statements.MaxRowsToFetch)

	intEmptySearch := make(map[string]interface{})
	intFullSearch := map[string]interface{}{"xkey": "ducksrus", "name": "Ducks R Us", "validDateSecs": dtForSearch.GetSeconds()}
	intOnlyDateSearch := map[string]interface{}{"validDateSecs": dtForSearch.GetSeconds()}
	intOnlyObjectIdSearch := map[string]interface{}{"xkey": "patoloco"}

	data, err := getSearchParmsData(dtForSearch)
	if err != nil {

		return
	}

	emptySearch := data["emptySearch"]
	fullSearch := data["fullSearch"]
	onlyDateSearch := data["onlyDateSearch"]
	onlyObjectIdSearch := data["onlyObjectIdSearch"]

	tests := []struct {
		name       string
		args       args
		want       string
		wantValues map[string]interface{}
		wantErr    bool
	}{
		{name: "Empty search", args: args{emptySearch}, want: sqlEmptyWhereClauseFinal, wantValues: intEmptySearch, wantErr: false},
		{name: "Full search", args: args{fullSearch}, want: sqlFullSearchFinal, wantValues: intFullSearch, wantErr: false},
		{name: "Only date search", args: args{onlyDateSearch}, want: sqlOnlyDateSearchFinal, wantValues: intOnlyDateSearch, wantErr: false},
		{name: "Only object Id search", args: args{onlyObjectIdSearch}, want: sqlOnlyObjectIdSearchFinal, wantValues: intOnlyObjectIdSearch, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &customer{}
			got, got1, err := c.getSQLForSearch(tt.args.searchParms)
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

func getSearchParmsData(dtForSearch *timestamp.Timestamp) (map[string]*proto.SearchParams, error) {

	data := make(map[string]*proto.SearchParams)
	data["emptySearch"] = &proto.SearchParams{}
	data["fullSearch"] = &proto.SearchParams{
		XKey:      "ducksrus",
		Name:      "Ducks R Us",
		ValidDate: dtForSearch,
	}
	data["onlyDateSearch"] = &proto.SearchParams{
		ValidDate: dtForSearch,
	}
	data["onlyObjectIdSearch"] = &proto.SearchParams{
		XKey: "patoloco",
	}

	return data, nil
}
