package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"goTemp/globalUtils"
	"goTemp/globalerrors"
	pb "goTemp/promotion"
)

func checkValidityDates(validFrom *timestamp.Timestamp, validThru *timestamp.Timestamp) ([]string, error) {
	var FailureDesc []string
	validDates := true
	if validFrom == nil {
		FailureDesc = append(FailureDesc, promoErr.MissingField("valid from"))
		validDates = false
	}
	if validThru == nil {
		FailureDesc = append(FailureDesc, promoErr.MissingField("valid thru"))
		validDates = false
	}
	if validDates {
		vd, err := globalUtils.TimeStampPPBToTime(validFrom, validThru)
		if err != nil {
			return nil, err
		}
		if vd[0].After(vd[1]) || vd[1].Equal(vd[0]) {
			FailureDesc = append(FailureDesc, promoErr.DtInvalidValidityDates(vd[0], vd[1]))
		}
	}
	return FailureDesc, nil
}

func checkMandatoryFields(promo *pb.Promotion) ([]string, error) {
	var FailureDesc []string
	if promo.GetName() == "" {
		FailureDesc = append(FailureDesc, promoErr.MissingField("name"))
	}
	if promo.GetCustomerId() == 0 {
		FailureDesc = append(FailureDesc, promoErr.MissingField("customer"))
	}
	dateValidation, err := checkValidityDates(promo.ValidFrom, promo.ValidThru)
	if err != nil {
		return nil, err
	}
	FailureDesc = append(FailureDesc, dateValidation...)

	return FailureDesc, nil
}

//func SetMandatoryFields(promo *pb.Promotion){
//	validThru, _ := globalUtils.TimeToTimeStampPPB(time.Now().AddDate(1,0,0))
//	promo.Active = false
//	promo.ApprovalStatus = 0
//	promo.PrevApprovalStatus = 0
//	promo.ValidFrom = ptypes.TimestampNow()
//	promo.ValidThru = validThru[0]
//}

func (p *Promotion) BeforeCreatePromotion(ctx context.Context, promotion *pb.Promotion, validationErr *pb.ValidationErr) error {
	_ = ctx
	validation, err := checkMandatoryFields(promotion)
	if err != nil {
		return err
	}
	validationErr.FailureDesc = append(validationErr.FailureDesc, validation...)
	if len(validationErr.FailureDesc) > 0 {
		return &globalerrors.ValidationError{Source: "BeforeCreatePromotion"}
	}
	return nil
}

func (p *Promotion) BeforeUpdatePromotion(ctx context.Context, promotion *pb.Promotion, validationErr *pb.ValidationErr) error {
	_ = ctx
	validation, err := checkMandatoryFields(promotion)
	if err != nil {
		return err
	}
	validationErr.FailureDesc = append(validationErr.FailureDesc, validation...)
	if len(validationErr.FailureDesc) > 0 {
		return &globalerrors.ValidationError{Source: "BeforeUpdatePromotion"}
	}
	return nil
}

func (p *Promotion) BeforeDeletePromotion(ctx context.Context, promotion *pb.Promotion, validationErr *pb.ValidationErr) error {
	_ = ctx
	if promotion.ApprovalStatus > 0 {
		validationErr.FailureDesc = append(validationErr.FailureDesc, promoErr.DelPromoNotInitialState())
	}
	if len(validationErr.FailureDesc) > 0 {
		return &globalerrors.ValidationError{Source: "BeforeDeletePromotion"}
	}
	return nil
}

func (p *Promotion) AfterCreatePromotion(ctx context.Context, promotion *pb.Promotion, afterFuncErr *pb.AfterFuncErr) error {
	_ = ctx
	_ = promotion
	if len(afterFuncErr.FailureDesc) > 0 {
		return &globalerrors.ValidationError{Source: "AfterCreatePromotion"}
	}
	return nil
}

func (p *Promotion) AfterUpdatePromotion(ctx context.Context, promotion *pb.Promotion, afterFuncErr *pb.AfterFuncErr) error {
	_ = ctx
	_ = promotion
	if len(afterFuncErr.FailureDesc) > 0 {
		return &globalerrors.ValidationError{Source: "AfterUpdatePromotion"}
	}
	return nil
}

func (p *Promotion) AfterDeletePromotion(ctx context.Context, promotion *pb.Promotion, afterFuncErr *pb.AfterFuncErr) error {
	_ = ctx
	_ = promotion
	if len(afterFuncErr.FailureDesc) > 0 {
		return &globalerrors.ValidationError{Source: "AfterDeletePromotion"}
	}
	return nil
}
