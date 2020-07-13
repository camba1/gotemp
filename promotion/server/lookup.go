package main

import (
	"context"
	"github.com/micro/go-micro/v2/client"
	custServ "goTemp/customer/proto"
	"log"
)

//idLookup: Struct in charge of looking up names based on ids
type idLookup struct{}

const (
	//cacheCustomerIdPrefix: Prefix to be used when storing a new customer id value in the cache
	cacheCustomerIdPrefix = "customerId"
)

//getCustomerNameFromService: Get a customer id directly from the customer service
func (i *idLookup) getCustomerNameFromService(ctx context.Context, customerId string) (string, error) {
	customerClient := custServ.NewCustomerSrvService(serviceNameCustomer, client.DefaultClient)
	customer, err := customerClient.GetCustomerById(ctx, &custServ.SearchId{XKey: customerId})
	if err != nil {
		return "", err
	}
	return customer.Name, nil
}

//getCustomerName: Lookup the customer id. first search the cache If customer if not found in cache
//call function to get it from the customer service and store it in the cache.
func (i *idLookup) getCustomerName(ctx context.Context, customerId string) (string, error) {
	customerName := ""
	customerName, err := glCache.GetCacheValue(cacheCustomerIdPrefix, customerId)
	if err != nil {
		log.Printf(promoErr.CacheCustomerNameNotFound(err))
		customerName, err = i.getCustomerNameFromService(ctx, customerId)
		if err != nil {
			log.Printf(promoErr.CustomerNameNotFound(customerId, err))
			return "", err
		}
		err := glCache.SetCacheValue(cacheCustomerIdPrefix, customerId, customerName)
		if err != nil {
			log.Printf(glErr.CacheUnableToWrite(customerId, err))
		}
	}

	return customerName, nil
}

//
//func testStore2() {
//	bal := glCache.Store.Options()
//	log.Printf("glCache settings %v\n", bal)
//	key := "mytest"
//	rec := store2.Record{
//		Key:    key,
//		Value:  []byte("mytest2"),
//		Expiry: 2 * time.Hour,
//	}
//
//	err := glCache.Store.Write(&rec)
//	if err != nil {
//		log.Printf("error writting. Error: %v", err)
//	}
//	rec1, err := glCache.Store.Read(key)
//	if err != nil {
//		log.Printf("Uanble to read. Error: %v\n", err)
//	}
//	log.Printf("Read 1: %v\n", rec1)
//
//
//
//	//err = glCache.Delete(key)
//	//if err != nil {
//	//	log.Printf("delete error %v\n", err)
//	//}
//	//myList, err = glCache.List(listLimit)
//	//if err != nil {
//	//	log.Printf("listing error %v\n", err)
//	//}
//	//log.Printf("list: %v", myList)
//}
//
