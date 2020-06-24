package main

import (
	"context"
	"encoding/json"
	"github.com/arangodb/go-driver"
	"goTemp/customer/proto"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	"log"
	"reflect"
	"strings"
)

const (
	customerCollectionName = "customer"
)

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

func (c *customer) GetCustomers(ctx context.Context, params *proto.SearchParams, customers *proto.Customers) error {
	//TODO: build dynamic params and pass to query
	values := make(map[string]interface{})
	values["name"] = params.Name

	cur, err := conn.Query(ctx, "FOR c IN customer FILTER c.name == @name RETURN c", values)
	if err != nil {
		log.Printf(custErr.SelectReadError(err))
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
		//log.Printf("Got document with meta %v\n", meta)
		customers.Customer = append(customers.Customer, customer)
	}
	return nil
}

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

func (c *customer) UpdateCustomer(ctx context.Context, inCustomer *proto.Customer, response *proto.Response) error {

	outCustomer := &proto.Customer{}
	customerMap := make(map[string]interface{})
	if errVal := c.BeforeUpdateCustomer(ctx, inCustomer, &proto.ValidationErr{}); errVal != nil {
		return errVal
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

//getAfterAlerts: Call the appropriate after create/update/delete function and return the alert validation errors
//These alerts  are logged, but do not cause the record processing to fail
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

//arangoMapToStruct: takes the map returned by arangoDB and converts it to a struct
// any fields not found in the struct are added to the extrafields field of the struct.
//This function is necessary because the arango drive will either return a struct
// that ignores the extra fields or a map that contain every thing (which cannot be
// sent as a gRPC message)
func arangoMapToStruct(inMap map[string]interface{}, customer *proto.Customer) error {

	//TODO: find a better way to do this ?

	//Get known fields into the struct
	fullMapBytes, err := json.Marshal(inMap)
	if err != nil {
		log.Printf("Unable to marshal full Arango map. Error: %v\n", err)
		return err
	}

	if err = json.Unmarshal(fullMapBytes, customer); err != nil {
		log.Printf("Unable to unmarshal byte full version of Arango map to struct. Error: %v\n", err)
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
		log.Printf("Unable to marshal partial Arango map. Error: %v\n", err)
		return err
	}
	tempStruct := &structpb.Struct{}
	if err = protojson.Unmarshal(partialMapBytes, tempStruct); err != nil {
		log.Printf("Unable to unmarshal byte partial version of Arango map to proto struct. Error: %v\n", err)
		return err
	}
	customer.ExtraFields = tempStruct

	return nil
}

//processExtraFields: Manage the fact that we may have data not defined in the Protobuf definition
//those records come in the ExtreFields field of the message
func processExtraFields(customer *proto.Customer) error {
	//TODO: Need to process extra fields instead of removing them before saving the record
	customer.ExtraFields = nil
	return nil
}
