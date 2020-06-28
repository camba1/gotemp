package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
	"goTemp/globalProtos"
	"goTemp/product/proto"
	userSrv "goTemp/user/proto"
	"log"
	"time"
)

//dateLayoutISO: Default time format for dates entered as strings
const dateLayoutISO = "2006-01-02"

//GetProductById: Call the product service and retrieve the product identified by a particular id
func GetProductById(ctx context.Context, productClient proto.ProductSrvService, custId *proto.SearchId) (*proto.Product, error) {

	product, err := productClient.GetProductById(ctx, custId)

	if err != nil {
		log.Printf("Unable to find product by Id. Error: %v", err)
		return nil, err
	}

	if product.XKey == "" {
		fmt.Printf("No Product found for id %s\n", custId.XKey)
		return nil, err
	}
	fmt.Printf("Pulled product by id %v\n", product)
	return product, nil

}

//GetProducts: Contact the product service and retrieve products based on a search criteria
func GetProducts(ctx context.Context, productClient proto.ProductSrvService) (*proto.Products, error) {
	_, searchDate := timeStringToTimestamp("2020-10-24")

	searchParms := proto.SearchParams{
		XKey:      "switch",
		Name:      "Play Switch Console",
		ValidDate: searchDate,
	}

	products, err := productClient.GetProducts(ctx, &searchParms)

	if err != nil {
		log.Printf("Unable to find products. Error: %v", err)
		return nil, err
	}
	if len(products.GetProduct()) == 0 {
		fmt.Printf("Products not found for parameters %v\n", &searchParms)
		return nil, fmt.Errorf("Products not found for parameters %v\n", &searchParms)
	}

	fmt.Printf("Pulled products %v\n", products)
	return products, nil

}

//CreateProduct: Call the product service and create a new product
func CreateProduct(ctx context.Context, productClient proto.ProductSrvService) (*proto.Product, error) {

	//var cust *proto.Product
	//var err error

	_, validThru := timeStringToTimestamp("2021-05-24")

	newProd := proto.Product{
		XKey:           "6308345766332077057",
		Name:           "Awesome Product",
		HierarchyLevel: "sku",
		ValidityDates: &globalProtos.GlValidityDate{
			ValidFrom: ptypes.TimestampNow(),
			ValidThru: validThru,
		},
		Modifications: &globalProtos.GlModification{
			CreateDate: ptypes.TimestampNow(),
			UpdateDate: ptypes.TimestampNow(),
			ModifiedBy: "123456789",
		},
	}

	resp, err := productClient.CreateProduct(ctx, &newProd)

	if err != nil {
		log.Printf("Unable to create product. Error: %v", err)
		return nil, err
	}
	fmt.Printf("Created product %v\n", resp.GetProduct())

	if len(resp.GetValidationErr().GetFailureDesc()) > 0 {
		fmt.Printf("Created product validations %v\n", resp.ValidationErr.FailureDesc)
	}
	return resp.GetProduct(), nil
}

//UpdateProduct: Call the product service and update a product
func UpdateProduct(ctx context.Context, productClient proto.ProductSrvService, prod *proto.Product) (*proto.Product, error) {
	_, validThru := timeStringToTimestamp("2021-06-26")

	prod.Name = "Just Ok Product"
	prod.HierarchyLevel = "cat"
	prod.ValidityDates.ValidFrom = ptypes.TimestampNow()
	prod.ValidityDates.ValidThru = validThru
	prod.Modifications.UpdateDate = ptypes.TimestampNow()
	prod.Modifications.ModifiedBy = "3308341401806443521"

	resp, err := productClient.UpdateProduct(ctx, prod)

	if err != nil {
		log.Printf("Unable to update product. Error: %v", err)
		return nil, err
	}
	fmt.Printf("Updated product %v\n", resp.GetProduct())

	if len(resp.GetValidationErr().GetFailureDesc()) > 0 {
		fmt.Printf("Update product validations %v\n", resp.GetValidationErr().GetFailureDesc())
	}

	return resp.GetProduct(), nil
}

//DeleteProduct: Call the product service and delete the user identified by a given id
func DeleteProduct(ctx context.Context, productClient proto.ProductSrvService, searchId *proto.SearchId) (int64, error) {

	resp, err := productClient.DeleteProduct(ctx, searchId)

	if err != nil {
		log.Printf("Unable to find product by Id. Error: %v", err)
		return 0, err
	}
	fmt.Printf("Count of products deleted %d\n", resp.AffectedCount)

	if len(resp.GetValidationErr().GetFailureDesc()) > 0 {
		fmt.Printf("Delete product validations %v\n", resp.GetValidationErr().GetFailureDesc())
	}
	return resp.GetAffectedCount(), nil
}

//timeStringToTimestamp: Convert time string to gRPC timestamp
func timeStringToTimestamp(priceVTstr string) (error, *timestamp.Timestamp) {
	priceVTtime, err := time.Parse(dateLayoutISO, priceVTstr)
	if err != nil {
		log.Fatalf("Unable to Format date %v", priceVTstr)
	}
	priceVT, err := ptypes.TimestampProto(priceVTtime)
	if err != nil {
		log.Fatalf("Unable to convert time to timestamp %v", priceVTtime)
	}
	return err, priceVT
}

//authUser: Call the user service and authenticate a user. receive a jwt token if successful
func authUser(srvClient userSrv.UserSrvService, user *userSrv.User) (*userSrv.Token, error) {
	token, err := srvClient.Auth(context.Background(), &userSrv.User{
		Email: user.Email,
		Pwd:   user.Pwd,
	})
	if err != nil {
		log.Printf("Unable to find token. Error: %v\n", err)
		return nil, err
	}
	fmt.Printf("Got token: %v\n", token)
	return token, err
}

//loginUser: Call authUser to get an authentication token and store it in the context for use on other tasks
func loginUser(srvClient userSrv.UserSrvService) (context.Context, error) {
	myUser := &userSrv.User{
		Pwd:   "1234",
		Email: "duck@mymail.com"}

	authToken, err := authUser(srvClient, myUser)
	if err != nil {
		return nil, err
	}

	ctx := metadata.NewContext(context.Background(), map[string]string{
		"token": authToken.Token,
	})
	return ctx, nil
}

func main() {
	service := micro.NewService(
		micro.Name("product.client"),
	)
	service.Init()
	fmt.Println("Client Running")

	// send requests
	ctx, err := loginUser(userSrv.NewUserSrvService("user", client.DefaultClient))
	if err != nil {
		return
	}

	productClient := proto.NewProductSrvService("product", service.Client())

	createdPromo, err := CreateProduct(ctx, productClient)
	if err != nil {
		return
	}

	_, _ = UpdateProduct(ctx, productClient, createdPromo)

	searchId := proto.SearchId{
		XKey: createdPromo.GetXKey(),
	}
	_, _ = GetProductById(ctx, productClient, &searchId)
	_, _ = DeleteProduct(ctx, productClient, &searchId)
	_, _ = GetProductById(ctx, productClient, &searchId)
	_, _ = GetProducts(ctx, productClient)
}
