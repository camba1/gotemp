package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/micro/go-micro/v2/metadata"
	"goTemp/customer/proto"
	"goTemp/globalProtos"
	"goTemp/globalUtils"
	"log"
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
		searchId:    &proto.SearchId{XKey: "13741"},
		outCustomer: &proto.Customer{},
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "Get a customer", args: validIdArg, want: "Ducks R Us", wantErr: false},
		{name: "Get a second customer", args: validIdArg2, want: "goodUpdatedCustomer", wantErr: false},
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

	//goodCustomer :=  custData["goodCustomer"]
	//newCustomerId, err := InsertTestCustomer(ctx, goodCustomer)
	//if err != nil {
	//	t.Errorf("UpdateCustomer() Unable to create test customer. error = %v", err)
	//}
	//updateCustomer.Id = newCustomerId

	updateCustomer.XKey = "13741"
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
