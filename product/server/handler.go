package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/arangodb/go-driver"
	"goTemp/product/proto"
	"goTemp/product/server/statements"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	"log"
	"reflect"
	"strings"
)

const (
	// productCollectionName: Name of the product collection (table)
	productCollectionName = "product"
)

//GetProductById: Get product from DB based on a given ID
func (p *Product) GetProductById(ctx context.Context, searchId *proto.SearchId, outProduct *proto.Product) error {
	col, err := conn.Collection(ctx, productCollectionName)
	if err != nil {
		log.Printf(prodErr.UnableToOpenCollection(productCollectionName))
		return err
	}
	customerMap := make(map[string]interface{})
	_, err = col.ReadDocument(ctx, searchId.GetXKey(), &customerMap)
	//_, err = col.ReadDocument(ctx, searchId.GetXKey(), outCustomer)
	if err != nil {
		log.Printf(prodErr.SelectRowReadError(err))
		return err
	}

	err = arangoMapToStruct(customerMap, outProduct)
	if err != nil {
		return err
	}
	return nil
}

//GetProducts: Search the products table in the DB based in a set of search parameters
func (p *Product) GetProducts(ctx context.Context, params *proto.SearchParams, products *proto.Products) error {
	values, sqlStatement, err2 := p.getSQLForSearch(params)
	if err2 != nil {
		return err2
	}

	cur, err := conn.Query(ctx, sqlStatement, values)
	if err != nil {
		log.Printf(prodErr.SelectReadError(err))
		return err
	}
	defer cur.Close()
	for cur.HasMore() {
		product := &proto.Product{}
		productMap := make(map[string]interface{})
		//_, err := cur.ReadDocument(ctx,&product)
		_, err := cur.ReadDocument(ctx, &productMap)
		if err != nil {
			log.Printf(prodErr.SelectScanError(err))
			return err
		}
		err = arangoMapToStruct(productMap, product)
		if err != nil {
			return err
		}

		products.Product = append(products.Product, product)
	}
	return nil
}

//CreateProduct: Creates a product in the Database. Calls before and after create functions
//for validations and sending record to the audit service via the broker
func (p *Product) CreateProduct(ctx context.Context, inProduct *proto.Product, response *proto.Response) error {
	outCustomer := &proto.Product{}
	customerMap := make(map[string]interface{})

	if errVal := p.BeforeCreateProduct(ctx, inProduct, &proto.ValidationErr{}); errVal != nil {
		return errVal
	}

	err := processExtraFields(inProduct)
	if err != nil {
		return err
	}

	//Get a reference to the collection
	col, err := conn.Collection(ctx, productCollectionName)
	if err != nil {
		log.Printf(prodErr.UnableToOpenCollection(productCollectionName))
		return err
	}

	//Create document
	ctxWithReturn := context.Background()
	ctxWithReturn = driver.WithReturnNew(ctxWithReturn, &customerMap)
	//ctxWithReturn =  driver.WithReturnNew(ctxWithReturn, outCustomer)
	_, err = col.CreateDocument(ctxWithReturn, inProduct)
	if err != nil {
		log.Printf(prodErr.InsertError(err))
		return err
	}

	//Map result back to Protobuf struct
	err = arangoMapToStruct(customerMap, outCustomer)
	if err != nil {
		return err
	}

	//after save processes
	response.Product = outCustomer
	failureDesc, err := p.getAfterAlerts(ctx, outCustomer, "AfterCreateProduct")
	if err != nil {
		return err
	}
	response.ValidationErr = &proto.ValidationErr{FailureDesc: failureDesc}

	return nil
}

//UpdateProduct: Update a product in the Database. Calls before and after create functions
//for validations and sending record to the audit service via the broker
func (p *Product) UpdateProduct(ctx context.Context, inProduct *proto.Product, response *proto.Response) error {

	outProduct := &proto.Product{}
	productMap := make(map[string]interface{})
	if errVal := p.BeforeUpdateProduct(ctx, inProduct, &proto.ValidationErr{}); errVal != nil {
		return errVal
	}

	err := processExtraFields(inProduct)
	if err != nil {
		return err
	}

	col, err := conn.Collection(ctx, productCollectionName)
	if err != nil {
		log.Printf(prodErr.UnableToOpenCollection(productCollectionName))
		return err
	}

	ctxWithReturn := context.Background()
	//ctxWithReturn =  driver.WithReturnNew(ctxWithReturn, outProduct)
	ctxWithReturn = driver.WithReturnNew(ctxWithReturn, &productMap)
	_, err = col.UpdateDocument(ctxWithReturn, inProduct.GetXKey(), inProduct)
	if err != nil {
		log.Printf(prodErr.UpdateError(err))
		return err
	}

	err = arangoMapToStruct(productMap, outProduct)
	if err != nil {
		return err
	}

	response.Product = outProduct
	failureDesc, err := p.getAfterAlerts(ctx, outProduct, "AfterUpdateProduct")
	if err != nil {
		return err
	}
	response.ValidationErr = &proto.ValidationErr{FailureDesc: failureDesc}

	return nil
}

//DeleteProduct: Delete a product in the Database based on the product ID (Xkey). Calls before and after create functions
//for validations and sending record to the audit service via the broker
func (p *Product) DeleteProduct(ctx context.Context, searchId *proto.SearchId, response *proto.Response) error {
	outCustomer := &proto.Product{}
	if err := p.GetProductById(ctx, searchId, outCustomer); err != nil {
		return err
	}
	if errVal := p.BeforeDeleteProduct(ctx, outCustomer, &proto.ValidationErr{}); errVal != nil {
		return errVal
	}

	col, err := conn.Collection(ctx, productCollectionName)
	if err != nil {
		log.Printf(prodErr.UnableToOpenCollection(productCollectionName))
		return err
	}
	_, err = col.RemoveDocument(ctx, searchId.GetXKey())
	if err != nil {
		log.Printf(prodErr.DeleteError(searchId.GetXKey(), err))
		return err
	}
	response.AffectedCount = 1

	failureDesc, err := p.getAfterAlerts(ctx, outCustomer, "AfterDeleteProduct")
	if err != nil {
		return err
	}
	response.ValidationErr = &proto.ValidationErr{FailureDesc: failureDesc}

	//log.Printf("deleted document with meta %v\n", meta)
	return nil
}

//getAfterAlerts: Call the appropriate after create/update/delete function and return the alert validation errors
//These alerts  are logged, but do not cause the record processing to fail
func (p *Product) getAfterAlerts(ctx context.Context, customer *proto.Product, operation string) ([]string, error) {
	afterFuncErr := &proto.AfterFuncErr{}
	var errVal error
	if operation == "AfterDeleteProduct" {
		errVal = p.AfterDeleteProduct(ctx, customer, afterFuncErr)
	}
	if operation == "AfterCreateProduct" {
		errVal = p.AfterCreateProduct(ctx, customer, afterFuncErr)
	}
	if operation == "AfterUpdateProduct" {
		errVal = p.AfterUpdateProduct(ctx, customer, afterFuncErr)
	}
	if errVal != nil {
		return []string{}, errVal
	}

	if len(afterFuncErr.GetFailureDesc()) > 0 {
		log.Printf("Alerts: %v: ", afterFuncErr.GetFailureDesc())
		return afterFuncErr.GetFailureDesc(), nil
	}
	return []string{}, nil
}

//arangoMapToStruct: takes the map returned by arangoDB and converts it to a struct
// any fields not found in the struct are added to the extrafields field of the struct.
//This function is necessary because the arango drive will either return a struct
// that ignores the extra fields or a map that contain every thing (which cannot be
// sent as a gRPC message)
func arangoMapToStruct(inMap map[string]interface{}, product *proto.Product) error {

	//TODO: find a better way to do this ?

	//Get known fields into the struct
	fullMapBytes, err := json.Marshal(inMap)
	if err != nil {
		log.Printf(glErr.MarshalFullMap(err))
		return err
	}

	if err = json.Unmarshal(fullMapBytes, product); err != nil {
		log.Printf(glErr.UnMarshalByteFullMap(err))
		return err
	}

	//Remove known fields from struct
	fields := reflect.TypeOf(proto.Product{})
	num := fields.NumField()
	for i := 0; i < num; i++ {
		field := fields.Field(i)
		if len(field.Tag) != 0 && field.PkgPath == "" {
			jsontags := strings.Split(field.Tag.Get("json"), ",")
			if len(jsontags) != 0 {
				delete(inMap, jsontags[0])
			}
		}
	}

	//Get additional fields into ExtraFields field of struct
	partialMapBytes, err := json.Marshal(inMap)
	if err != nil {
		log.Printf(glErr.MarshalPartialMap(err))
		return err
	}
	tempStruct := &structpb.Struct{}
	if err = protojson.Unmarshal(partialMapBytes, tempStruct); err != nil {
		log.Printf(glErr.UnMarshalBytePartialMap(err))
		return err
	}
	product.ExtraFields = tempStruct

	return nil
}

//processExtraFields: Manage the fact that we may have data not defined in the Protobuf definition
//those records come in the ExtreFields field of the message
func processExtraFields(product *proto.Product) error {
	//TODO: Need to process extra fields instead of removing them before saving the record
	product.ExtraFields = nil
	return nil
}

//getSQLForSearch: Combine the where clause built in the buildSearchWhereClause method with the rest of the sql
//statement to return the final search for users sql statement
func (p *Product) getSQLForSearch(searchParms *proto.SearchParams) (map[string]interface{}, string, error) {
	sql := statements.SqlSelectAll.String()
	sqlWhereClause, values, err := p.buildSearchWhereClause(searchParms)
	if err != nil {
		return nil, "", err
	}

	sqlStatement := fmt.Sprintf(sql, sqlWhereClause, statements.MaxRowsToFetch)
	return values, sqlStatement, nil
}

//buildSearchWhereClause: Builds a sql string to be used as the where clause in a sql statement. It also returns an interface
//slice with the values to be used as replacements in the sql statement. Currently only handles equality constraints, except
//for the date lookup which is done  as a contains clause
func (p *Product) buildSearchWhereClause(searchParms *proto.SearchParams) (string, map[string]interface{}, error) {
	sqlWhereClause := " FILTER 1==1 "
	values := make(map[string]interface{})

	if searchParms.GetXKey() != "" {
		sqlWhereClause += fmt.Sprint(" AND p._key ==  @xkey")
		values["xkey"] = searchParms.GetXKey()
	}
	if searchParms.GetName() != "" {
		sqlWhereClause += fmt.Sprint(" AND p.name == @name")
		values["name"] = searchParms.GetName()
	}
	if searchParms.GetValidDate() != nil {
		secs := searchParms.GetValidDate().GetSeconds()
		sqlWhereClause += fmt.Sprint(" AND p.validityDates.validFrom.seconds <= @validDateSecs AND p.validityDates.validThru.seconds >= @validDateSecs")
		values["validDateSecs"] = secs
	}
	return sqlWhereClause, values, nil
}
