package main

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"goTemp/globalUtils"
	pb "goTemp/promotion"
	"testing"
	"time"
)

func Test_checkMandatoryFields(t *testing.T) {
	nextYear := getNextYear()
	type args struct {
		promo *pb.Promotion
	}
	emptyPromo := pb.Promotion{
		Id:                 0,
		Name:               "",
		Description:        "",
		ValidFrom:          nil,
		ValidThru:          nil,
		Active:             false,
		CustomerId:         0,
		Product:            nil,
		ApprovalStatus:     0,
		PrevApprovalStatus: 0,
	}
	goodPromo := pb.Promotion{
		Id:                 0,
		Name:               "test",
		Description:        "test",
		ValidFrom:          ptypes.TimestampNow(),
		ValidThru:          nextYear,
		Active:             false,
		CustomerId:         1233,
		Product:            nil,
		ApprovalStatus:     0,
		PrevApprovalStatus: 0,
	}
	noNamePromo := pb.Promotion{
		Id:                 0,
		Name:               "",
		Description:        "",
		ValidFrom:          ptypes.TimestampNow(),
		ValidThru:          nextYear,
		Active:             false,
		CustomerId:         1233,
		Product:            nil,
		ApprovalStatus:     0,
		PrevApprovalStatus: 0,
	}
	noCustomerPromo := pb.Promotion{
		Id:                 0,
		Name:               "Test",
		Description:        "",
		ValidFrom:          ptypes.TimestampNow(),
		ValidThru:          nextYear,
		Active:             false,
		CustomerId:         0,
		Product:            nil,
		ApprovalStatus:     0,
		PrevApprovalStatus: 0,
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "Empty Promo", args: args{&emptyPromo}, want: 4, wantErr: false},
		{name: "Missing Name Promo", args: args{&noNamePromo}, want: 1, wantErr: false},
		{name: "Missing Customer Promo", args: args{&noCustomerPromo}, want: 1, wantErr: false},
		{name: "Good Promo", args: args{&goodPromo}, want: 0, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkMandatoryFields(tt.args.promo)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkMandatoryFields() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("checkMandatoryFields() got %v failed validations, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkValidityDates(t *testing.T) {
	nextYear := getNextYear()
	type args struct {
		validFrom *timestamp.Timestamp
		validThru *timestamp.Timestamp
	}
	tests := []struct {
		name string
		args args
		//want    []string
		want    int
		wantErr bool
	}{
		{name: "Nil validity dates", args: args{validFrom: nil, validThru: nil}, want: 2, wantErr: false},
		{name: "Nil valid from date", args: args{validFrom: nil, validThru: ptypes.TimestampNow()}, want: 1, wantErr: false},
		{name: "Nil valid thru date", args: args{validFrom: ptypes.TimestampNow(), validThru: nil}, want: 1, wantErr: false},
		{name: "Invalid equal  Dates", args: args{validFrom: ptypes.TimestampNow(), validThru: ptypes.TimestampNow()}, want: 1, wantErr: false},
		{name: "Invalid Dates", args: args{validFrom: nextYear, validThru: ptypes.TimestampNow()}, want: 1, wantErr: false},
		{name: "Valid dates", args: args{validFrom: ptypes.TimestampNow(), validThru: nextYear}, want: 0, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkValidityDates(tt.args.validFrom, tt.args.validThru)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkValidityDates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("checkValidityDates() got %v failed validations, want %v", got, tt.want)
			}
		})
	}
}

func getNextYear() *timestamp.Timestamp {
	myDates, _ := globalUtils.TimeToTimeStampPPB(time.Now().AddDate(1, 0, 0))
	return myDates[0]
}
