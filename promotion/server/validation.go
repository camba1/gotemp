package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"goTemp/globalUtils"
	"goTemp/globalerrors"
	"goTemp/promotion/proto"
	"log"
	"strconv"
)

//mb: Broker instance to send/receive message from pub/sub system
var mb globalUtils.MyBroker

//checkValidityDates: Enusre that the valid date from and valid date thru are populated properly
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

//checkMandatoryFields: Ensure that all mandatory fields are populated properly
func checkMandatoryFields(promo *proto.Promotion) ([]string, error) {
	var FailureDesc []string
	if promo.GetName() == "" {
		FailureDesc = append(FailureDesc, promoErr.MissingField("name"))
	}
	if promo.GetCustomerId() == "" {
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

//BeforeCreatePromotion: Call data validations before creating a promotion
func (p *Promotion) BeforeCreatePromotion(ctx context.Context, promotion *proto.Promotion, validationErr *proto.ValidationErr) error {
	_ = ctx
	validation, err := checkMandatoryFields(promotion)
	if err != nil {
		return err
	}
	validationErr.FailureDesc = append(validationErr.FailureDesc, validation...)
	if len(validationErr.FailureDesc) > 0 {
		return &globalerrors.ValidationError{Source: "BeforeCreatePromotion", FailureDesc: validationErr.FailureDesc}
	}
	return nil
}

//BeforeUpdatePromotion: Call data validations before updating a promotion
func (p *Promotion) BeforeUpdatePromotion(ctx context.Context, promotion *proto.Promotion, validationErr *proto.ValidationErr) error {
	_ = ctx
	validation, err := checkMandatoryFields(promotion)
	if err != nil {
		return err
	}
	validationErr.FailureDesc = append(validationErr.FailureDesc, validation...)
	if len(validationErr.FailureDesc) > 0 {
		return &globalerrors.ValidationError{Source: "BeforeUpdatePromotion", FailureDesc: validationErr.FailureDesc}
	}
	return nil
}

//BeforeDeletePromotion: Call data validations before deleting a promotion
func (p *Promotion) BeforeDeletePromotion(ctx context.Context, promotion *proto.Promotion, validationErr *proto.ValidationErr) error {
	_ = ctx
	if promotion.ApprovalStatus > 0 {
		validationErr.FailureDesc = append(validationErr.FailureDesc, promoErr.DelPromoNotInitialState())
	}
	if len(validationErr.FailureDesc) > 0 {
		return &globalerrors.ValidationError{Source: "BeforeDeletePromotion", FailureDesc: validationErr.FailureDesc}
	}
	return nil
}

//AfterCreatePromotion: Call processes to be run after promotion create
func (p *Promotion) AfterCreatePromotion(ctx context.Context, promotion *proto.Promotion, afterFuncErr *proto.AfterFuncErr) error {
	_ = ctx

	failureDesc := p.sendAudit(ctx, serviceName, "AfterCreatePromotion", "insert", "promotion", promotion.GetId(), promotion)
	if len(failureDesc) > 0 {
		afterFuncErr.FailureDesc = append(afterFuncErr.FailureDesc, failureDesc)
	}
	//_ = promotion
	//if len(afterFuncErr.FailureDesc) > 0 {
	//	return &globalerrors.ValidationError{Source: "AfterCreatePromotion"}
	//}
	return nil
}

//AfterUpdatePromotion: Call processes to be run after promotion update
func (p *Promotion) AfterUpdatePromotion(ctx context.Context, promotion *proto.Promotion, afterFuncErr *proto.AfterFuncErr) error {
	_ = ctx

	failureDesc := p.sendAudit(ctx, serviceName, "AfterUpdatePromotion", "update", "promotion", promotion.GetId(), promotion)
	if len(failureDesc) > 0 {
		afterFuncErr.FailureDesc = append(afterFuncErr.FailureDesc, failureDesc)
	}

	//_ = promotion
	//if len(afterFuncErr.FailureDesc) > 0 {
	//	return &globalerrors.ValidationError{Source: "AfterUpdatePromotion"}
	//}
	return nil
}

//AfterDeletePromotion: Call processes to be run after promotion delete
func (p *Promotion) AfterDeletePromotion(ctx context.Context, promotion *proto.Promotion, afterFuncErr *proto.AfterFuncErr) error {
	_ = ctx

	failureDesc := p.sendAudit(ctx, serviceName, "AfterDeletePromotion", "Delete", "promotion", promotion.GetId(), promotion)
	if len(failureDesc) > 0 {
		afterFuncErr.FailureDesc = append(afterFuncErr.FailureDesc, failureDesc)
	}

	//_ = promotion
	//if len(afterFuncErr.FailureDesc) > 0 {
	//	return &globalerrors.ValidationError{Source: "AfterDeletePromotion"}
	//}
	return nil
}

//sendAudit: Convert a promotion to a byte array, and call AuditUtil to send message with updated promotion to audit service
func (p *Promotion) sendAudit(ctx context.Context, serviceName, actionFunc, actionType string, objectName string, iObjectId int64, promotion *proto.Promotion) string {

	if !glDisableAuditRecords {

		objectId := strconv.FormatInt(iObjectId, 10)
		byteUser, err := mb.ProtoToByte(promotion)
		if err != nil {
			return glErr.AudFailureSending(actionType, objectId, err)
		}

		log.Printf("sending audit record for %s:", actionFunc)
		return globalUtils.AuditSend(ctx, mb, serviceName, actionFunc, actionType, objectName, objectId, byteUser)

	}
	return ""
}
