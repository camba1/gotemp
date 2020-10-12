package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
	"goTemp/promotion/proto"
	userSrv "goTemp/user/proto"
	"log"
	"time"
)

const (
	//serviceName: service identifier
	serviceName = "goTemp.api.promotion"
	//serviceNameUser: service identifier for user service
	serviceNameUser = "goTemp.api.user"
)

//dateLayoutISO defines the default time format for dates entered as strings
const dateLayoutISO = "2006-01-02"

//GetPromotionById calls the promotion service and retrieve the promotion identified by a particular id
func GetPromotionById(ctx context.Context, promotionClient proto.PromotionSrvService, promoId *proto.SearchId) (*proto.Promotion, error) {

	promotion, err := promotionClient.GetPromotionById(ctx, promoId)

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

//GetPromotions contacts the promotion service and retrieve promotions based on a search criteria
func GetPromotions(ctx context.Context, promotionClient proto.PromotionSrvService) (*proto.Promotions, error) {
	_, searchDate := timeStringToTimestamp("2020-10-24")

	searchParms := proto.SearchParams{
		Name:       "Promo1",
		CustomerId: "ducksrus",
		ValidDate:  searchDate,
	}

	promotions, err := promotionClient.GetPromotions(ctx, &searchParms)

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

//CreatePromotion calls the promotion service and create a new promotion
func CreatePromotion(ctx context.Context, promotionClient proto.PromotionSrvService) (*proto.Promotion, error) {

	_, validThru := timeStringToTimestamp("2021-05-24")

	//disc := &proto.Discount{
	//	Id:          123456789,
	//	Value:       0.59,
	//	Type:        0,
	//	Description: "Good customer",
	//}
	//prod1 := &proto.Product{
	//	Id:      7308345766332077057,
	//	UpcCode: "prod1a",
	//}
	//prod1.Discount = append(prod1.Discount, disc)
	//prod2 := &proto.Product{
	//	Id:      8308345766441128962,
	//	UpcCode: "prod1",
	//}
	//prod2.Discount = append(prod1.Discount, disc)

	newPromo := proto.Promotion{
		Id:                 6308345766332077057,
		Name:               "Super Promo",
		Description:        "Super Promo",
		ValidFrom:          ptypes.TimestampNow(),
		ValidThru:          validThru,
		Active:             false,
		CustomerId:         "ducksrus",
		Product:            nil,
		ApprovalStatus:     0,
		PrevApprovalStatus: 0,
	}
	//newPromo.Product = append(newPromo.Product, prod1, prod2)

	resp, err := promotionClient.CreatePromotion(ctx, &newPromo)

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

//UpdatePromotion calls the promotion service and update a promotion
func UpdatePromotion(ctx context.Context, promotionClient proto.PromotionSrvService, promo *proto.Promotion) (*proto.Promotion, error) {
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
	promo.CustomerId = "patoloco"
	promo.Product = nil
	promo.ApprovalStatus = 0
	promo.PrevApprovalStatus = 0

	resp, err := promotionClient.UpdatePromotion(ctx, promo)

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

//DeletePromotion calls the promotion service and delete the promotion identified by a given id
func DeletePromotion(ctx context.Context, promotionClient proto.PromotionSrvService, promoId *proto.SearchId) (int64, error) {

	resp, err := promotionClient.DeletePromotion(ctx, promoId)

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

//timeStringToTimestamp converts time string to gRPC timestamp
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

//authUser calls the user service and authenticate a user. receive a jwt token if successful
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

//loginUser calls authUser to get an authentication token and store it in the context for use on other tasks
func loginUser(srvClient userSrv.UserSrvService) (context.Context, error) {
	myUser := &userSrv.User{
		Pwd:   "1234",
		Email: "duck@mymail.com"}

	authToken, err := authUser(srvClient, myUser)
	if err != nil {
		return nil, err
	}

	ctx := metadata.NewContext(context.Background(), map[string]string{
		//"token": authToken.Token,
		"Authorization": "Bearer " + authToken.Token,
	})
	return ctx, nil
}

func main() {

	service := micro.NewService(
		micro.Name("promotion.client"),
	)
	service.Init()
	fmt.Println("Client Running")

	// send requests
	ctx, err := loginUser(userSrv.NewUserSrvService(serviceNameUser, client.DefaultClient))
	if err != nil {
		return
	}

	promotionClient := proto.NewPromotionSrvService(serviceName, service.Client())

	createdPromo, err := CreatePromotion(ctx, promotionClient)
	if err != nil {
		return
	}

	_, _ = UpdatePromotion(ctx, promotionClient, createdPromo)

	searchId := proto.SearchId{
		Id: createdPromo.Id,
	}
	_, _ = GetPromotionById(ctx, promotionClient, &searchId)
	_, _ = DeletePromotion(ctx, promotionClient, &searchId)
	_, _ = GetPromotionById(ctx, promotionClient, &searchId)
	_, _ = GetPromotions(ctx, promotionClient)
}
