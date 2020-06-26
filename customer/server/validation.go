package main

import (
	"context"
	"goTemp/customer/proto"
	"goTemp/globalUtils"
	"goTemp/globalerrors"
	"log"
	"strconv"
	"time"
)

//checkMandatoryFields: Ensure that all mandatory fields are populated properly
func checkMandatoryFields(customer *proto.Customer) ([]string, error) {
	var FailureDesc []string
	if customer.GetName() == "" {
		FailureDesc = append(FailureDesc, glErr.MissingField("name"))
	}
	dateValidation, err := globalUtils.CheckValidityDates(customer.GetValidityDates().GetValidFrom(), customer.GetValidityDates().GetValidThru())
	if err != nil {
		return nil, err
	}
	FailureDesc = append(FailureDesc, dateValidation...)

	return FailureDesc, nil
}

//SetMandatoryFields: Preset the mandatory fields that need to be populated before insert,delete or update
func SetMandatoryFields(ctx context.Context, customer *proto.Customer, isInsert bool) error {
	tempDates, _ := globalUtils.TimeToTimeStampPPB(time.Now(), time.Now().AddDate(1, 0, 0))
	if isInsert {
		if customer.GetValidityDates().GetValidFrom() == nil {
			customer.GetValidityDates().ValidFrom = tempDates[0]
			customer.GetValidityDates().ValidThru = tempDates[1]
		}
		customer.Modifications.CreateDate = tempDates[0]
	}
	customer.Modifications.UpdateDate = tempDates[0]

	currentUser, err := getCurrentUser(ctx)
	if err != nil {
		return err
	}
	customer.Modifications.ModifiedBy = currentUser
	return nil
}

//getCurrentUser: Get the user from the context. Notice that the authorization service returns a int64 and we convert to string
func getCurrentUser(ctx context.Context) (string, error) {
	var auth globalUtils.AuthUtils
	currentUser, err := auth.GetCurrentUserFromContext(ctx)
	if err != nil {
		log.Printf(glErr.AuthNoUserInToken(err))
		return "", err
	}
	return strconv.FormatInt(currentUser, 10), nil
}

func (c *customer) BeforeCreateCustomer(ctx context.Context, customer *proto.Customer, validationErr *proto.ValidationErr) error {
	_ = ctx

	err := SetMandatoryFields(ctx, customer, true)
	if err != nil {
		return err
	}

	validation, err := checkMandatoryFields(customer)
	if err != nil {
		return err
	}
	validationErr.FailureDesc = append(validationErr.FailureDesc, validation...)

	if len(validationErr.FailureDesc) > 0 {
		return &globalerrors.ValidationError{Source: "BeforeCreateUser", FailureDesc: validationErr.FailureDesc}
	}
	return nil
}

func (c *customer) BeforeUpdateCustomer(ctx context.Context, customer *proto.Customer, validationErr *proto.ValidationErr) error {
	_ = ctx

	err := SetMandatoryFields(ctx, customer, false)
	if err != nil {
		return err
	}

	validation, err := checkMandatoryFields(customer)
	if err != nil {
		return err
	}
	validationErr.FailureDesc = append(validationErr.FailureDesc, validation...)

	if len(validationErr.FailureDesc) > 0 {
		return &globalerrors.ValidationError{Source: "BeforeCreateCustomer", FailureDesc: validationErr.FailureDesc}
	}
	return nil
}

func (c *customer) BeforeDeleteCustomer(ctx context.Context, customer *proto.Customer, validationErr *proto.ValidationErr) error {
	_ = ctx

	err := SetMandatoryFields(ctx, customer, false)
	if err != nil {
		return err
	}

	if len(validationErr.FailureDesc) > 0 {
		return &globalerrors.ValidationError{Source: "BeforeDeleteUser", FailureDesc: validationErr.FailureDesc}
	}
	return nil
}

func (c *customer) AfterCreateCustomer(ctx context.Context, customer *proto.Customer, afterFuncErr *proto.AfterFuncErr) error {
	_ = ctx

	failureDesc := c.sendUserAudit(ctx, serviceName, "AfterCreateCustomer", "insert", "customer", customer.GetXKey(), customer)
	if len(failureDesc) > 0 {
		afterFuncErr.FailureDesc = append(afterFuncErr.FailureDesc, failureDesc)
	}

	//if len(afterFuncErr.FailureDesc) > 0 {
	//	return &globalerrors.ValidationError{Source: "AfterCreateUser", FailureDesc: afterFuncErr.FailureDesc}
	//}
	return nil
}

func (c *customer) AfterUpdateCustomer(ctx context.Context, customer *proto.Customer, afterFuncErr *proto.AfterFuncErr) error {
	_ = ctx

	failureDesc := c.sendUserAudit(ctx, serviceName, "AfterUpdateCustomer", "update", "customer", customer.GetXKey(), customer)
	if len(failureDesc) > 0 {
		afterFuncErr.FailureDesc = append(afterFuncErr.FailureDesc, failureDesc)
	}

	//if len(afterFuncErr.FailureDesc) > 0 {
	//	return &globalerrors.ValidationError{Source: "AfterCreatePromotion"}
	//}
	return nil
}

func (c *customer) AfterDeleteCustomer(ctx context.Context, customer *proto.Customer, afterFuncErr *proto.AfterFuncErr) error {
	_ = ctx

	failureDesc := c.sendUserAudit(ctx, serviceName, "AfterDeleteCustomer", "Delete", "customer", customer.GetXKey(), customer)
	if len(failureDesc) > 0 {
		afterFuncErr.FailureDesc = append(afterFuncErr.FailureDesc, failureDesc)
	}

	//if len(afterFuncErr.FailureDesc) > 0 {
	//	return &globalerrors.ValidationError{Source: "AfterCreatePromotion"}
	//}
	return nil
}

//sendUserAudit: Convert a user to a byte array, and call AuditUtil to send message with updated promotion to audit service
func (c *customer) sendUserAudit(ctx context.Context, serviceName, actionFunc, actionType string, objectName string, objectId string, customer *proto.Customer) string {
	if !glDisableAuditRecords {
		byteMessage, err := mb.ProtoToByte(customer)
		if err != nil {
			return glErr.AudFailureSending(actionType, objectId, err)
		}

		return globalUtils.AuditSend(ctx, mb, serviceName, actionFunc, actionType, objectName, objectId, byteMessage)
	}
	return ""

}
