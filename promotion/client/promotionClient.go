package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/micro/go-micro/v2"
	"goTemp/promotion/proto"
	"log"
	"time"
)

const serverAddressEnvVar = "SERVERADDRESS"

const dateLayoutISO = "2006-01-02"

func GetPromotionById(promotionClient proto.PromotionSrvService, promoId *proto.SearchId) (*proto.Promotion, error) {

	//var promotion *proto.Promotion
	//var err error
	//if serverAddress != "" {
	//	promotion, err = promotionClient.GetPromotionById(context.Background(), promoId, client.WithAddress(serverAddress))
	//} else {
	promotion, err := promotionClient.GetPromotionById(context.Background(), promoId)
	//}

	if err != nil {
		log.Printf("Unable to find promotion by Id. Error: %v", err)
		return nil, err
	}

	if promotion.Id == 0 {
		fmt.Printf("No Promotion found for id %d\n", promoId.Id)
		return nil, err
	}
	fmt.Printf("Pulled promotion by id %v\n", promotion)
	return promotion, nil

}

func GetPromotions(promotionClient proto.PromotionSrvService) (*proto.Promotions, error) {
	_, searchDate := timeStringToTimestamp("2020-10-24")

	searchParms := proto.SearchParams{
		//Id:         2308345766332077057,
		Name: "Promo1",
		//Name: 		"Super Promo",
		//ProductId:  7308345766332077057,
		CustomerId: 3308341401806443521,
		ValidDate:  searchDate,
	}

	//var promotion *proto.Promotions
	//var err error
	//if serverAddress != "" {
	//	promotion, err = promotionClient.GetPromotions(context.Background(), &searchParms, client.WithAddress(serverAddress))
	//} else {
	//	promotion, err = promotionClient.GetPromotions(context.Background(), &searchParms)
	//}

	promotions, err := promotionClient.GetPromotions(context.Background(), &searchParms)

	if err != nil {
		log.Printf("Unable to find promotions. Error: %v", err)
		return nil, err
	}
	if len(promotions.GetPromotion()) == 0 {
		fmt.Printf("Promotions not found for parameters %v\n", &searchParms)
		return nil, fmt.Errorf("Promotions not found for parameters %v\n", &searchParms)
	}

	fmt.Printf("Pulled promotions %v\n", promotions)
	return promotions, nil

}

func CreatePromotion(promotionClient proto.PromotionSrvService) (*proto.Promotion, error) {

	//var promo *proto.Promotion
	//var err error

	_, validThru := timeStringToTimestamp("2021-05-24")

	disc := &proto.Discount{
		Id:          123456789,
		Value:       0.59,
		Type:        0,
		Description: "Good customer",
	}
	prod1 := &proto.Product{
		Id:      7308345766332077057,
		UpcCode: "prod1a",
	}
	prod1.Discount = append(prod1.Discount, disc)
	prod2 := &proto.Product{
		Id:      8308345766441128962,
		UpcCode: "prod1",
	}
	prod2.Discount = append(prod1.Discount, disc)

	newPromo := proto.Promotion{
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

	//if serverAddress != "" {
	//	promo, err = promotionClient.CreatePromotion(context.Background(), &newPromo, client.WithAddress(serverAddress))
	//} else {
	//	promo, err = promotionClient.CreatePromotion(context.Background(), &newPromo)
	//}

	resp, err := promotionClient.CreatePromotion(context.Background(), &newPromo)

	if err != nil {
		log.Printf("Unable to create promotion. Error: %v", err)
		return nil, err
	}
	fmt.Printf("Created promotion %v\n", resp.GetPromotion())

	if len(resp.GetValidationErr().GetFailureDesc()) > 0 {
		fmt.Printf("Created promotion validations %v\n", resp.ValidationErr.FailureDesc)
	}
	return resp.GetPromotion(), nil
}

func UpdatePromotion(promotionClient proto.PromotionSrvService, promo *proto.Promotion) (*proto.Promotion, error) {
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

	//var outPromo *proto.Promotion
	//var err error
	//if serverAddress != "" {
	//	outPromo, err = promotionClient.UpdatePromotion(context.Background(), promo, client.WithAddress(serverAddress))
	//} else {
	//	outPromo, err = promotionClient.UpdatePromotion(context.Background(), promo)
	//}

	resp, err := promotionClient.UpdatePromotion(context.Background(), promo)

	if err != nil {
		log.Printf("Unable to update promotion. Error: %v", err)
		return nil, err
	}
	fmt.Printf("Updated promotion %v\n", resp.GetPromotion())

	if len(resp.GetValidationErr().GetFailureDesc()) > 0 {
		fmt.Printf("Update promotion validations %v\n", resp.GetValidationErr().GetFailureDesc())
	}

	return resp.GetPromotion(), nil
}

func DeletePromotion(promotionClient proto.PromotionSrvService, promoId *proto.SearchId) (int64, error) {

	//searchId := pb.SearchId{
	//	Id: 2312030045339653121,
	//}
	//
	//var affectedCount *proto.AffectedCount
	//var err error
	//if serverAddress != "" {
	//	affectedCount, err = promotionClient.DeletePromotion(context.Background(), promoId, client.WithAddress(serverAddress))
	//} else {
	//	affectedCount, err = promotionClient.DeletePromotion(context.Background(), promoId)
	//}

	resp, err := promotionClient.DeletePromotion(context.Background(), promoId)

	if err != nil {
		log.Printf("Unable to find promotion by Id. Error: %v", err)
		return 0, err
	}
	fmt.Printf("Count of promotions deleted %d\n", resp.AffectedCount)

	if len(resp.GetValidationErr().GetFailureDesc()) > 0 {
		fmt.Printf("Delete user validations %v\n", resp.GetValidationErr().GetFailureDesc())
	}
	return resp.GetAffectedCount(), nil
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
	promotionClient := proto.NewPromotionSrvService("promotion", service.Client())

	createdPromo, err := CreatePromotion(promotionClient)
	if err != nil {
		return
	}

	_, _ = UpdatePromotion(promotionClient, createdPromo)

	searchId := proto.SearchId{
		Id: createdPromo.Id,
	}
	_, _ = GetPromotionById(promotionClient, &searchId)
	_, _ = DeletePromotion(promotionClient, &searchId)
	_, _ = GetPromotionById(promotionClient, &searchId)
	_, _ = GetPromotions(promotionClient)
}
