package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/arangodb/go-driver"
	"goTemp/customer/proto"
	"goTemp/customer/server/statements"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	"log"
	"reflect"
	"strings"
)

const (
	customerCollectionName = "customer"
)

// GetCustomerById gets customer from DB based on a given ID
func (c *customer) GetCustomerById(ctx context.Context, searchId *proto.SearchId, outCustomer *proto.Customer) error {
	col, err := conn.Collection(ctx, customerCollectionName)
	if err != nil {
		log.Printf(custErr.UnableToOpenCollection(customerCollectionName))
		return err
	}
	customerMap := make(map[string]interface{})
	_, err = col.ReadDocument(ctx, searchId.GetXKey(), &customerMap)
	//_, err = col.ReadDocument(ctx, searchId.GetXKey(), outCustomer)
	if err != nil {
		log.Printf(custErr.SelectRowReadError(err))
		return err
	}

	err = arangoMapToStruct(customerMap, outCustomer)
	if err != nil {
		return err
	}
	return nil
}

// GetCustomers searches the customers table in the DB based in a set of search parameters
func (c *customer) GetCustomers(ctx context.Context, params *proto.SearchParams, customers *proto.Customers) error {

	values, sqlStatement, err2 := c.getSQLForSearch(params)
	if err2 != nil {
		return err2
	}

	cur, err := conn.Query(ctx, sqlStatement, values)
	if err != nil {
		log.Printf(custErr.SelectReadError(err))
		return err
	}
	defer cur.Close()
	for cur.HasMore() {
		customer := &proto.Customer{}
		customerMap := make(map[string]interface{})
		//_, err := cur.ReadDocument(ctx,&customer)
		_, err := cur.ReadDocument(ctx, &customerMap)
		if err != nil {
			log.Printf(custErr.SelectScanError(err))
			return err
		}
		err = arangoMapToStruct(customerMap, customer)
		if err != nil {
			return err
		}

		customers.Customer = append(customers.Customer, customer)
	}
	return nil
}

// CreateCustomer creates a customer in the Database. Calls before and after create functions
// for validations and sending record to the audit service via the broker
func (c *customer) CreateCustomer(ctx context.Context, inCustomer *proto.Customer, response *proto.Response) error {

	outCustomer := &proto.Customer{}
	customerMap := make(map[string]interface{})

	if errVal := c.BeforeCreateCustomer(ctx, inCustomer, &proto.ValidationErr{}); errVal != nil {
		return errVal
	}

	err := processExtraFields(inCustomer)
	if err != nil {
		return err
	}

	//Get a reference to the collection
	col, err := conn.Collection(ctx, customerCollectionName)
	if err != nil {
		log.Printf(custErr.UnableToOpenCollection(customerCollectionName))
		return err
	}

	//Create document
	ctxWithReturn := context.Background()
	ctxWithReturn = driver.WithReturnNew(ctxWithReturn, &customerMap)
	//ctxWithReturn =  driver.WithReturnNew(ctxWithReturn, outCustomer)
	_, err = col.CreateDocument(ctxWithReturn, inCustomer)
	if err != nil {
		log.Printf(custErr.InsertError(err))
		return err
	}

	//Map result back to Protobuf struct
	err = arangoMapToStruct(customerMap, outCustomer)
	if err != nil {
		return err
	}

	//after save processes
	response.Customer = outCustomer
	failureDesc, err := c.getAfterAlerts(ctx, outCustomer, "AfterCreateCustomer")
	if err != nil {
		return err
	}
	response.ValidationErr = &proto.ValidationErr{FailureDesc: failureDesc}

	return nil
}

// UpdateCustomer updates a customer in the Database. Calls before and after create functions
// for validations and sending record to the audit service via the broker
func (c *customer) UpdateCustomer(ctx context.Context, inCustomer *proto.Customer, response *proto.Response) error {

	outCustomer := &proto.Customer{}
	customerMap := make(map[string]interface{})
	if errVal := c.BeforeUpdateCustomer(ctx, inCustomer, &proto.ValidationErr{}); errVal != nil {
		return errVal
	}

	err := processExtraFields(inCustomer)
	if err != nil {
		return err
	}

	col, err := conn.Collection(ctx, customerCollectionName)
	if err != nil {
		log.Printf(custErr.UnableToOpenCollection(customerCollectionName))
		return err
	}

	ctxWithReturn := context.Background()
	//ctxWithReturn =  driver.WithReturnNew(ctxWithReturn, outCustomer)
	ctxWithReturn = driver.WithReturnNew(ctxWithReturn, &customerMap)
	_, err = col.UpdateDocument(ctxWithReturn, inCustomer.GetXKey(), inCustomer)
	if err != nil {
		log.Printf(custErr.UpdateError(err))
		return err
	}

	err = arangoMapToStruct(customerMap, outCustomer)
	if err != nil {
		return err
	}

	response.Customer = outCustomer
	failureDesc, err := c.getAfterAlerts(ctx, outCustomer, "AfterUpdateCustomer")
	if err != nil {
		return err
	}
	response.ValidationErr = &proto.ValidationErr{FailureDesc: failureDesc}

	return nil
}

// DeleteCustomer deletes a customer in the Database based on the product ID (Xkey). Calls before and after create functions
// for validations and sending record to the audit service via the broker
func (c *customer) DeleteCustomer(ctx context.Context, searchId *proto.SearchId, response *proto.Response) error {

	outCustomer := &proto.Customer{}
	if err := c.GetCustomerById(ctx, searchId, outCustomer); err != nil {
		return err
	}
	if errVal := c.BeforeDeleteCustomer(ctx, outCustomer, &proto.ValidationErr{}); errVal != nil {
		return errVal
	}

	col, err := conn.Collection(ctx, customerCollectionName)
	if err != nil {
		log.Printf(custErr.UnableToOpenCollection(customerCollectionName))
		return err
	}
	_, err = col.RemoveDocument(ctx, searchId.GetXKey())
	if err != nil {
		log.Printf(custErr.DeleteError(searchId.GetXKey(), err))
		return err
	}
	response.AffectedCount = 1

	failureDesc, err := c.getAfterAlerts(ctx, outCustomer, "AfterDeleteCustomer")
	if err != nil {
		return err
	}
	response.ValidationErr = &proto.ValidationErr{FailureDesc: failureDesc}

	//log.Printf("deleted document with meta %v\n", meta)
	return nil
}

// getAfterAlerts calls the appropriate after create/update/delete function and return the alert validation errors
// These alerts  are logged, but do not cause the record processing to fail
func (c *customer) getAfterAlerts(ctx context.Context, customer *proto.Customer, operation string) ([]string, error) {
	afterFuncErr := &proto.AfterFuncErr{}
	var errVal error
	if operation == "AfterDeleteCustomer" {
		errVal = c.AfterDeleteCustomer(ctx, customer, afterFuncErr)
	}
	if operation == "AfterCreateCustomer" {
		errVal = c.AfterCreateCustomer(ctx, customer, afterFuncErr)
	}
	if operation == "AfterUpdateCustomer" {
		errVal = c.AfterUpdateCustomer(ctx, customer, afterFuncErr)
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

// arangoMapToStruct takes the map returned by arangoDB and converts it to a struct
// any fields not found in the struct are added to the extrafields field of the struct.
// This function is necessary because the arango drive will either return a struct
// that ignores the extra fields or a map that contain every thing (which cannot be
// sent as a gRPC message)
func arangoMapToStruct(inMap map[string]interface{}, customer *proto.Customer) error {

	//TODO: find a better way to do this ?

	//Get known fields into the struct
	fullMapBytes, err := json.Marshal(inMap)
	if err != nil {
		log.Printf(glErr.MarshalFullMap(err))
		return err
	}

	if err = json.Unmarshal(fullMapBytes, customer); err != nil {
		log.Printf(glErr.UnMarshalByteFullMap(err))
		return err
	}

	//Remove known fields from struct
	fields := reflect.TypeOf(proto.Customer{})
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
	customer.ExtraFields = tempStruct

	return nil
}

// processExtraFields Manages the fact that we may have data not defined in the Protobuf definition
// those records come in the ExtreFields field of the message
func processExtraFields(customer *proto.Customer) error {
	//TODO: Need to process extra fields instead of removing them before saving the record
	customer.ExtraFields = nil
	return nil
}

// getSQLForSearch combines the where clause built in the buildSearchWhereClause method with the rest of the sql
// statement to return the final search for users sql statement
func (c *customer) getSQLForSearch(searchParms *proto.SearchParams) (map[string]interface{}, string, error) {
	sql := statements.SqlSelectAll.String()
	sqlWhereClause, values, err := c.buildSearchWhereClause(searchParms)
	if err != nil {
		return nil, "", err
	}

	sqlStatement := fmt.Sprintf(sql, sqlWhereClause, statements.MaxRowsToFetch)
	return values, sqlStatement, nil
}

// buildSearchWhereClause builds a sql string to be used as the where clause in a sql statement. It also returns an interface
// slice with the values to be used as replacements in the sql statement. Currently only handles equality constraints, except
// for the date lookup which is done  as a contains clause
func (c *customer) buildSearchWhereClause(searchParms *proto.SearchParams) (string, map[string]interface{}, error) {
	sqlWhereClause := " FILTER 1==1 "
	values := make(map[string]interface{})

	if searchParms.GetXKey() != "" {
		sqlWhereClause += fmt.Sprint(" AND c._key ==  @xkey")
		values["xkey"] = searchParms.GetXKey()
	}
	if searchParms.GetName() != "" {
		sqlWhereClause += fmt.Sprint(" AND c.name == @name")
		values["name"] = searchParms.GetName()
	}
	if searchParms.GetValidDate() != nil {
		secs := searchParms.GetValidDate().GetSeconds()
		sqlWhereClause += fmt.Sprint(" AND c.validityDates.validFrom.seconds <= @validDateSecs AND c.validityDates.validThru.seconds >= @validDateSecs")
		values["validDateSecs"] = secs
	}
	return sqlWhereClause, values, nil
}
