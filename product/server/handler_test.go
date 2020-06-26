package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/micro/go-micro/v2/metadata"
	"goTemp/globalProtos"
	"goTemp/globalUtils"
	"goTemp/product/proto"
	"goTemp/product/server/statements"
	"log"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestProduct_CreateProduct(t *testing.T) {
	type args struct {
		ctx       context.Context
		inProduct *proto.Product
		response  *proto.Response
	}

	loadConfig()
	conn = connectToDB()

	id := base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(123456789, 10)))
	ctx := metadata.Set(context.Background(), "userid", id)

	dtValidFrom, _ := time.Parse(globalUtils.DateLayoutISO, "2021-06-26")
	dtValidThru, _ := time.Parse(globalUtils.DateLayoutISO, "2021-07-26")
	testData := getTestData(dtValidFrom, dtValidThru)
	goodItem := testData["goodItem"]
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Good Customer", args: args{ctx: ctx, inProduct: goodItem, response: &proto.Response{}}, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Product{}
			if err := p.CreateProduct(tt.args.ctx, tt.args.inProduct, tt.args.response); (err != nil) != tt.wantErr {
				t.Errorf("CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProduct_DeleteProduct(t *testing.T) {
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
	testData := getTestData(dtValidFrom, dtValidThru)
	goodItem := testData["goodItem"]
	newRecordId, err := InsertTestProduct(ctx, goodItem)
	if err != nil {
		t.Errorf("DeleteProduct() Unable to create test product. error = %v", err)
	}
	recordToDeleteId := newRecordId

	validIdArg := args{
		ctx:      ctx,
		searchId: &proto.SearchId{XKey: recordToDeleteId},
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
		{name: "Delete a record", args: validIdArg, want: 1, wantErr: false},
		{name: "Nothing to delete", args: nothingToDelArg, want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Product{}
			if err := p.DeleteProduct(tt.args.ctx, tt.args.searchId, tt.args.response); (err != nil) != tt.wantErr {
				t.Errorf("DeleteProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.args.response.AffectedCount != tt.want {
				t.Errorf("DeleteProduct() wanted to delete %v records, deleted %v", tt.args.response.AffectedCount, tt.want)
			}
		})
	}
}

func TestProduct_GetProductById(t *testing.T) {
	type args struct {
		ctx        context.Context
		searchId   *proto.SearchId
		outProduct *proto.Product
	}

	loadConfig()
	conn = connectToDB()
	ctx := context.Background()
	validIdArg := args{
		ctx:        ctx,
		searchId:   &proto.SearchId{XKey: "switch"},
		outProduct: &proto.Product{},
	}
	validIdArg2 := args{
		ctx:        ctx,
		searchId:   &proto.SearchId{XKey: "tele"},
		outProduct: &proto.Product{},
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "Get a customer", args: validIdArg, want: "Play Switch Console", wantErr: false},
		{name: "Get a second customer", args: validIdArg2, want: "Watch me TV", wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Product{}
			if err := p.GetProductById(tt.args.ctx, tt.args.searchId, tt.args.outProduct); (err != nil) != tt.wantErr {
				t.Errorf("GetProductById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProduct_GetProducts(t *testing.T) {
	type args struct {
		ctx      context.Context
		params   *proto.SearchParams
		products *proto.Products
	}

	loadConfig()
	conn = connectToDB()
	ctx := context.Background()
	validIdArg := args{
		ctx:      ctx,
		params:   &proto.SearchParams{Name: "Watch me TV"},
		products: &proto.Products{},
	}
	inValidIdArg := args{
		ctx:      ctx,
		params:   &proto.SearchParams{Name: "XYZ"},
		products: &proto.Products{},
	}

	tests := []struct {
		name      string
		args      args
		want      string
		wantCount int
		wantErr   bool
	}{
		{name: "Find one record", args: validIdArg, want: "Watch me TV", wantCount: 1, wantErr: false},
		{name: "No records to be found", args: inValidIdArg, want: "", wantCount: 0, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Product{}
			err := p.GetProducts(tt.args.ctx, tt.args.params, tt.args.products)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProducts() error = %v, wantErr %v", err, tt.wantErr)
			}

			if len(tt.args.products.Product) != tt.wantCount {
				t.Errorf("GetCustomers() record count = %v, wantCount %v", len(tt.args.products.Product), tt.wantCount)
			}
			if len(tt.args.products.Product) == 1 && tt.args.products.Product[0].Name != tt.want {
				t.Errorf(" expected customer name == %s, got %s", tt.want, tt.args.products.Product[0].Name)
			}

		})
	}
}

func TestProduct_UpdateProduct(t *testing.T) {
	type args struct {
		ctx       context.Context
		inProduct *proto.Product
		response  *proto.Response
	}

	loadConfig()
	conn = connectToDB()

	id := base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(123456789, 10)))
	ctx := metadata.Set(context.Background(), "userid", id)

	dtValidFrom, _ := time.Parse(globalUtils.DateLayoutISO, "2021-06-26")
	dtValidThru, _ := time.Parse(globalUtils.DateLayoutISO, "2021-07-26")
	testData := getTestData(dtValidFrom, dtValidThru)
	updateItem := testData["updateItem"]

	//create a record
	goodItem := testData["goodItem"]
	newRecordId, err := InsertTestProduct(ctx, goodItem)
	if err != nil {
		t.Errorf("UpdateProduct() Unable to create test product. error = %v", err)
	}

	updateItem.XKey = newRecordId

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Update Product", args: args{ctx: ctx, inProduct: updateItem, response: &proto.Response{}}, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Product{}
			if err := p.UpdateProduct(tt.args.ctx, tt.args.inProduct, tt.args.response); (err != nil) != tt.wantErr {
				t.Errorf("UpdateProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	//Clean up
	p := &Product{}
	err = p.DeleteProduct(ctx, &proto.SearchId{XKey: newRecordId}, &proto.Response{})
	if err != nil {
		t.Errorf("UpdateCustomer() Unable to delete created customer %v. Error: %v", newRecordId, err)
	}
}

func TestProduct_getSQLForSearch(t *testing.T) {
	type args struct {
		searchParms *proto.SearchParams
	}

	sql := statements.SqlSelectAll.String()

	dtTmForSearch, _ := time.Parse(globalUtils.DateLayoutISO, "2021-06-26")
	protoDate, _ := globalUtils.TimeToTimeStampPPB(dtTmForSearch)
	dtForSearch := protoDate[0]

	sqlEmptyWhereClause := " FILTER 1==1 "
	sqlFullSearch := sqlEmptyWhereClause + " AND p._key ==  @xkey AND p.name == @name AND p.validityDates.validFrom.seconds <= @validDateSecs AND p.validityDates.validThru.seconds >= @validDateSecs"
	sqlOnlyDateSearch := sqlEmptyWhereClause + " AND p.validityDates.validFrom.seconds <= @validDateSecs AND p.validityDates.validThru.seconds >= @validDateSecs"
	sqlOnlyObjectIdSearch := sqlEmptyWhereClause + " AND p._key ==  @xkey"

	sqlEmptyWhereClauseFinal := fmt.Sprintf(sql, sqlEmptyWhereClause, statements.MaxRowsToFetch)
	sqlFullSearchFinal := fmt.Sprintf(sql, sqlFullSearch, statements.MaxRowsToFetch)
	sqlOnlyDateSearchFinal := fmt.Sprintf(sql, sqlOnlyDateSearch, statements.MaxRowsToFetch)
	sqlOnlyObjectIdSearchFinal := fmt.Sprintf(sql, sqlOnlyObjectIdSearch, statements.MaxRowsToFetch)

	intEmptySearch := make(map[string]interface{})
	intFullSearch := map[string]interface{}{"xkey": "switch", "name": "Play Switch Console", "validDateSecs": dtForSearch.GetSeconds()}
	intOnlyDateSearch := map[string]interface{}{"validDateSecs": dtForSearch.GetSeconds()}
	intOnlyObjectIdSearch := map[string]interface{}{"xkey": "tele"}

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
			p := &Product{}
			got, got1, err := p.getSQLForSearch(tt.args.searchParms)
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

func getTestData(validfrom time.Time, validThru time.Time) map[string]*proto.Product {
	products := make(map[string]*proto.Product)
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
	goodItem := proto.Product{
		Name:           "goodItem",
		HierarchyLevel: "sku",
		ValidityDates:  &vdates,
		Modifications:  &mods,
	}
	updateItem := proto.Product{
		XKey:           "",
		Name:           "goodUpdatedItem",
		HierarchyLevel: "cat",
		ValidityDates:  &vdates,
		Modifications:  &mods,
	}
	products["goodItem"] = &goodItem
	products["updateItem"] = &updateItem
	return products
}

func InsertTestProduct(ctx context.Context, customerToInsert *proto.Product) (string, error) {
	c := &Product{}
	var insertedResponse proto.Response
	err := c.CreateProduct(ctx, customerToInsert, &insertedResponse)
	if err != nil {
		return "", fmt.Errorf("UpdateProduct() unable to create product before update error = %v", err)
	}
	return insertedResponse.Product.XKey, nil
}

func getSearchParmsData(dtForSearch *timestamp.Timestamp) (map[string]*proto.SearchParams, error) {

	data := make(map[string]*proto.SearchParams)
	data["emptySearch"] = &proto.SearchParams{}
	data["fullSearch"] = &proto.SearchParams{
		XKey:      "switch",
		Name:      "Play Switch Console",
		ValidDate: dtForSearch,
	}
	data["onlyDateSearch"] = &proto.SearchParams{
		ValidDate: dtForSearch,
	}
	data["onlyObjectIdSearch"] = &proto.SearchParams{
		XKey: "tele",
	}

	return data, nil
}
