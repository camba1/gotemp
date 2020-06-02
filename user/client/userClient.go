package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/micro/go-micro/v2"
	pb "goTemp/user/proto"
	"log"
	"time"
)

//const serverAddressEnvVar = "SERVERADDRESS"

//var serverAddress = os.Getenv(serverAddressEnvVar)

const dateLayoutISO = "2006-01-02"

func CreateUser(srvClient pb.UserSrvService) (*pb.User, error) {
	var outUser *pb.User
	var err error

	_, validThru := timeStringToTimestamp("2021-05-24")

	newUser := pb.User{
		Firstname: "Huge",
		Lastname:  "Microbe",
		ValidFrom: ptypes.TimestampNow(),
		ValidThru: validThru,
		Active:    true,
		Pwd:       "1234",
		Email:     "microbes@tiny.com",
		Company:   "Tiny",
	}

	//if serverAddress != "" {
	//	outUser, err = srvClient.CreateUser(context.Background(), &newUser, client.WithAddress(serverAddress))
	//} else {
	outUser, err = srvClient.CreateUser(context.Background(), &newUser)
	//}

	if err != nil {
		log.Printf("Unable to create user. Error: %v", err)
		return nil, err
	}
	fmt.Printf("Created user %v\n", outUser)
	return outUser, nil
}

func UpdateUser(srvClient pb.UserSrvService, user *pb.User) (*pb.User, error) {
	var outUser *pb.User
	var err error
	_, validThru := timeStringToTimestamp("2021-06-26")

	user.Firstname = "Incredible"
	user.Lastname = "Green Guy"
	user.ValidFrom = ptypes.TimestampNow()
	user.ValidThru = validThru
	user.Active = false
	user.Pwd = "5678"
	user.Email = "microbes@tiny.com"
	//user.Email = "cow@mymail.com"
	user.Company = "Tiny"

	//if serverAddress != "" {
	//outUser, err = srvClient.UpdateUser(context.Background(), user, client.WithAddress(serverAddress))
	//} else {
	outUser, err = srvClient.UpdateUser(context.Background(), user)
	//}

	if err != nil {
		log.Printf("Unable to update user. Error: %v", err)
		return nil, err
	}
	fmt.Printf("Updated user %v\n", outUser)
	return outUser, nil
}

func GetUserById(srvClient pb.UserSrvService, searchId *pb.SearchId) (*pb.User, error) {
	var outUser *pb.User
	var err error
	//if serverAddress != "" {
	//	outUser, err = srvClient.GetUserById(context.Background(), searchId, client.WithAddress(serverAddress))
	//} else {
	outUser, err = srvClient.GetUserById(context.Background(), searchId)
	//}

	if err != nil {
		log.Printf("Unable to find user by Id. Error: %v", err)
		return nil, err
	}

	if outUser.Id == 0 {
		log.Printf("No user found for id %d\n", searchId.Id)
		return nil, fmt.Errorf("No user found for id %d\n", searchId.Id)
	}

	fmt.Printf("Pulled user by id %v\n", outUser)
	return outUser, nil
}

func DeleteUser(srvClient pb.UserSrvService, searchId *pb.SearchId) (int64, error) {
	var affectedCount *pb.AffectedCount
	var err error
	//if serverAddress != "" {
	//affectedCount, err = srvClient.DeleteUser(context.Background(), searchId, client.WithAddress(serverAddress))
	//} else {
	affectedCount, err = srvClient.DeleteUser(context.Background(), searchId)
	//}

	if err != nil {
		log.Printf("Unable to find user by Id. Error: %v", err)
		return 0, err
	}
	fmt.Printf("Count of users deleted %d\n", affectedCount.Value)
	return affectedCount.GetValue(), nil
}

func GetUsers(srvClient pb.UserSrvService) (*pb.Users, error) {
	_, searchDate := timeStringToTimestamp("2020-10-24")

	searchParms := pb.SearchParams{
		//Id:        1234,
		Fisrtname: "Super",
		Lastname:  "Duck",
		ValidDate: searchDate,
		Email:     "duck@mymail.com",
	}

	var outUsers *pb.Users
	var err error
	//if serverAddress != "" {
	//	outUsers, err = srvClient.GetUsers(context.Background(), &searchParms, client.WithAddress(serverAddress))
	//} else {
	outUsers, err = srvClient.GetUsers(context.Background(), &searchParms)
	//}

	if err != nil {
		log.Fatalf("Unable to find users. Error: %v", err)
		return nil, err
	}
	if len(outUsers.GetUser()) == 0 {
		fmt.Printf("Users not found for parameters %v\n", &searchParms)
		return nil, fmt.Errorf("Users not found for parameters %v\n", &searchParms)
	}
	fmt.Printf("Pulled users %v\n", outUsers)
	return outUsers, nil

}

func authUser(srvClient pb.UserSrvService, user *pb.User) (*pb.Token, error) {
	token, err := srvClient.Auth(context.Background(), &pb.User{
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
		micro.Name("user.client"),
	)
	service.Init()
	fmt.Println("Client Running")
	srvClient := pb.NewUserSrvService("user", service.Client())

	createdUser, err := CreateUser(srvClient)
	if err != nil {
		return
	}

	_, err = authUser(srvClient, &pb.User{
		Pwd:   "1234",
		Email: createdUser.Email})
	if err != nil {
		return
	}
	_, _ = UpdateUser(srvClient, createdUser)

	searchId := pb.SearchId{
		Id: createdUser.Id,
	}

	_, _ = GetUserById(srvClient, &searchId)
	_, _ = DeleteUser(srvClient, &searchId)
	_, _ = GetUserById(srvClient, &searchId)
	_, _ = GetUsers(srvClient)
}
