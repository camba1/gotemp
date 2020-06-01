package main

import (
	"context"
	"goTemp/globalUtils"
	"goTemp/globalerrors"
	pb "goTemp/user/proto"
)

func checkMandatoryFields(user *pb.User) ([]string, error) {
	var FailureDesc []string
	if user.GetFirstname() == "" {
		FailureDesc = append(FailureDesc, glErr.MissingField("first name"))
	}
	if user.GetLastname() == "" {
		FailureDesc = append(FailureDesc, glErr.MissingField("last name"))
	}
	if user.GetPwd() == "" {
		FailureDesc = append(FailureDesc, glErr.MissingField("password"))
	}
	if user.GetEmail() == "" {
		FailureDesc = append(FailureDesc, glErr.MissingField("email"))
	}
	dateValidation, err := globalUtils.CheckValidityDates(user.ValidFrom, user.ValidThru)
	if err != nil {
		return nil, err
	}
	FailureDesc = append(FailureDesc, dateValidation...)

	return FailureDesc, nil
}

func checkEmail(ctx context.Context, user *pb.User, u *User, isInsert bool) (string, error) {
	var usersWithSameEmail pb.Users
	err := u.GetUsersByEmail(ctx, &pb.SearchString{Value: user.Email}, &usersWithSameEmail)
	if err != nil {
		return "", err
	}
	return checkEmailUnique(user, &usersWithSameEmail, isInsert), nil
}

func checkEmailUnique(user *pb.User, users *pb.Users, isInsert bool) string {

	if users == nil {
		return ""
	}
	usersCount := len(users.User)
	if isInsert && usersCount > 0 {
		return userErr.InsertDupEmail()
	}
	if !isInsert && usersCount > 0 {
		if usersCount > 1 || (usersCount == 1 && users.User[0].Id != user.Id) {
			return userErr.UpdateDupEmail()
		}
	}
	return ""

}

//func SetMandatoryFields(user *pb.User){
//	validThru, _ := globalUtils.TimeToTimeStampPPB(time.Now().AddDate(1,0,0))
//	user.Active = false
//	user.ValidFrom = ptypes.TimestampNow()
//	user.ValidThru = validThru[0]
//}

func (u *User) BeforeCreateUser(ctx context.Context, user *pb.User, validationErr *pb.ValidationErr) error {
	_ = ctx
	validation, err := checkMandatoryFields(user)
	if err != nil {
		return err
	}
	validationErr.FailureDesc = append(validationErr.FailureDesc, validation...)

	dupEmail, err2 := checkEmail(ctx, user, u, true)
	if err2 != nil {
		return err2
	}
	if dupEmail != "" {
		validationErr.FailureDesc = append(validationErr.FailureDesc, dupEmail)
	}

	if len(validationErr.FailureDesc) > 0 {
		return &globalerrors.ValidationError{Source: "BeforeCreateUser", FailureDesc: validationErr.FailureDesc}
	}
	return nil
}

func (u *User) BeforeUpdateUser(ctx context.Context, user *pb.User, validationErr *pb.ValidationErr) error {
	_ = ctx

	validation, err := checkMandatoryFields(user)
	if err != nil {
		return err
	}
	validationErr.FailureDesc = append(validationErr.FailureDesc, validation...)

	dupEmail, err2 := checkEmail(ctx, user, u, false)
	if err2 != nil {
		return err2
	}
	if dupEmail != "" {
		validationErr.FailureDesc = append(validationErr.FailureDesc, dupEmail)
	}

	if len(validationErr.FailureDesc) > 0 {
		return &globalerrors.ValidationError{Source: "BeforeUpdatePromotion", FailureDesc: validationErr.FailureDesc}
	}
	return nil
}

func (u *User) BeforeDeleteUser(ctx context.Context, user *pb.User, validationErr *pb.ValidationErr) error {
	_ = ctx
	if user.GetActive() {
		validationErr.FailureDesc = append(validationErr.FailureDesc, userErr.DelUserActive())
	}
	if len(validationErr.FailureDesc) > 0 {
		return &globalerrors.ValidationError{Source: "BeforeDeleteUser", FailureDesc: validationErr.FailureDesc}
	}
	return nil
}

func (u *User) AfterCreateUser(ctx context.Context, user *pb.User, afterFuncErr *pb.AfterFuncErr) error {
	_ = ctx
	_ = user
	if len(afterFuncErr.FailureDesc) > 0 {
		return &globalerrors.ValidationError{Source: "AfterCreatePromotion"}
	}
	return nil
}

func (u *User) AfterUpdateUser(ctx context.Context, user *pb.User, afterFuncErr *pb.AfterFuncErr) error {
	_ = ctx
	_ = user
	if len(afterFuncErr.FailureDesc) > 0 {
		return &globalerrors.ValidationError{Source: "AfterCreatePromotion"}
	}
	return nil
}

func (u *User) AfterDeleteUser(ctx context.Context, user *pb.User, afterFuncErr *pb.AfterFuncErr) error {
	_ = ctx
	_ = user
	if len(afterFuncErr.FailureDesc) > 0 {
		return &globalerrors.ValidationError{Source: "AfterCreatePromotion"}
	}
	return nil
}
