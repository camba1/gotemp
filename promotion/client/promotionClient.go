package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"log"
	"os"
	"time"

	pb "goTemp/promotion"
)

//const serverAddress = "127.0.0.1:50051"

const serverAddressEnvVar = "SERVERADDRESS"

var serverAddress = os.Getenv(serverAddressEnvVar)

const dateLayoutISO = "2006-01-02"

func GetPromotionById(promotionClient pb.PromotionSrvService, promoId *pb.SearchId) {

	var promotion *pb.Promotion
	var err error
	if serverAddress != "" {
		promotion, err = promotionClient.GetPromotionById(context.Background(), promoId, client.WithAddress(serverAddress))
	} else {
		promotion, err = promotionClient.GetPromotionById(context.Background(), promoId)
	}

	if err != nil {
		log.Printf("Unable to find promotion by Id. Error: %v", err)
	}

	if promotion.Id == 0 {
		fmt.Printf("No Promotion found for id %d\n", promoId.Id)
	} else {
		fmt.Printf("Pulled promotion by id %v\n", promotion)
	}

}

func GetPromotions(promotionClient pb.PromotionSrvService) {
	_, searchDate := timeStringToTimestamp("2020-10-24")

	searchParms := pb.SearchParams{
		//Id:         2308345766332077057,
		Name: "Promo1",
		//Name: 		"Super Promo",
		//ProductId:  7308345766332077057,
		CustomerId: 3308341401806443521,
		ValidDate:  searchDate,
	}

	var promotion *pb.Promotions
	var err error
	if serverAddress != "" {
		promotion, err = promotionClient.GetPromotions(context.Background(), &searchParms, client.WithAddress(serverAddress))
	} else {
		promotion, err = promotionClient.GetPromotions(context.Background(), &searchParms)
	}

	if err != nil {
		log.Fatalf("Unable to find promotions. Error: %v", err)
	}
	if len(promotion.GetPromotion()) == 0 {
		fmt.Printf("Promotions not found for parameters %v\n", &searchParms)
	} else {
		fmt.Printf("Pulled promotions %v\n", promotion)
	}

}

func CreatePromotion(promotionClient pb.PromotionSrvService) *pb.Promotion {

	var promo *pb.Promotion
	var err error

	_, validThru := timeStringToTimestamp("2021-05-24")

	disc := &pb.Discount{
		Id:          123456789,
		Value:       0.59,
		Type:        0,
		Description: "Good customer",
	}
	prod1 := &pb.Product{
		Id:      7308345766332077057,
		UpcCode: "prod1a",
	}
	prod1.Discount = append(prod1.Discount, disc)
	prod2 := &pb.Product{
		Id:      8308345766441128962,
		UpcCode: "prod1",
	}
	prod2.Discount = append(prod1.Discount, disc)

	newPromo := pb.Promotion{
		Id:                 6308345766332077057,
		Name:               "Super Promo",
		Description:        "Super Promo",
		ValidFrom:          ptypes.TimestampNow(),
		ValidThru:          validThru,
		Active:             false,
		CustomerId:         3308341401806443521,
		Product:            nil,
		ApprovalStatus:     0,
		PrevApprovalStatus: 0,
	}
	newPromo.Product = append(newPromo.Product, prod1, prod2)

	if serverAddress != "" {
		promo, err = promotionClient.CreatePromotion(context.Background(), &newPromo, client.WithAddress(serverAddress))
	} else {
		promo, err = promotionClient.CreatePromotion(context.Background(), &newPromo)
	}

	if err != nil {
		log.Fatalf("Unable to create promotion. Error: %v", err)
	}
	fmt.Printf("Created promotion %v\n", promo)
	return promo
}

func UpdatePromotion(promotionClient pb.PromotionSrvService, promo *pb.Promotion) {
	_, validThru := timeStringToTimestamp("2021-06-26")

	//disc := &pb.Discount{
	//	Id:          123456789,
	//	Value:       0.59,
	//	Type:        0,
	//	Description: "Good customer",
	//}
	//prod1 := &pb.Product{
	//	Id:        7308345766332077057,
	//	UpcCode: "prod1a",
	//}
	//prod1.Discount = append(prod1.Discount, disc)
	//prod2 := &pb.Product{
	//	Id:        8308345766441128962,
	//	UpcCode: "prod1",
	//}
	//prod2.Discount = append(prod1.Discount, disc)

	promo.Name = "Super Promo2"
	promo.Description = "Super Promo2"
	promo.ValidFrom = ptypes.TimestampNow()
	promo.ValidThru = validThru
	promo.Active = true
	promo.CustomerId = 3308341401806443521
	promo.Product = nil
	promo.ApprovalStatus = 0
	promo.PrevApprovalStatus = 0

	//newPromo.Product =  append(newPromo.Product, prod1, prod2)

	var outPromo *pb.Promotion
	var err error
	if serverAddress != "" {
		outPromo, err = promotionClient.UpdatePromotion(context.Background(), promo, client.WithAddress(serverAddress))
	} else {
		outPromo, err = promotionClient.UpdatePromotion(context.Background(), promo)
	}

	if err != nil {
		log.Fatalf("Unable to update promotion. Error: %v", err)
	}
	fmt.Printf("Updated promotion %v\n", outPromo)
}

func DeletePromotion(promotionClient pb.PromotionSrvService, promoId *pb.SearchId) {

	//searchId := pb.SearchId{
	//	Id: 2312030045339653121,
	//}

	var affectedCount *pb.AffectedCount
	var err error
	if serverAddress != "" {
		affectedCount, err = promotionClient.DeletePromotion(context.Background(), promoId, client.WithAddress(serverAddress))
	} else {
		affectedCount, err = promotionClient.DeletePromotion(context.Background(), promoId)
	}

	if err != nil {
		log.Fatalf("Unable to find promotion by Id. Error: %v", err)
	}
	fmt.Printf("Count of promotions deleted %d\n", affectedCount.Value)
}

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

func main() {

	service := micro.NewService(
		micro.Name("promotion.client"),
	)
	service.Init()
	fmt.Println("Client Running")
	promotionClient := pb.NewPromotionSrvService("promotion", service.Client())

	createdPromo := CreatePromotion(promotionClient)
	UpdatePromotion(promotionClient, createdPromo)
	searchId := pb.SearchId{
		Id: createdPromo.Id,
	}
	GetPromotionById(promotionClient, &searchId)
	DeletePromotion(promotionClient, &searchId)
	GetPromotionById(promotionClient, &searchId)
	GetPromotions(promotionClient)
}
