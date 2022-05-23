// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package qlmodel

import (
	"fmt"
	"io"
	"strconv"

	commonModel "github.com/vbetsun/surgeon-intern-app/internal/graph/model"
)

type ActivateUserInput struct {
	Email          string `json:"email"`
	ActivationCode string `json:"activationCode"`
	GivenName      string `json:"givenName"`
	FamilyName     string `json:"familyName"`
	Password       string `json:"password"`
}

type ActivationVerification struct {
	Status ActivationVerificationStatus `json:"Status"`
}

type ActivationVerificationInput struct {
	Email          string `json:"email"`
	ActivationCode string `json:"activationCode"`
}

type DetailedUser struct {
	UserID      string                    `json:"userId"`
	Email       string                    `json:"email"`
	GivenName   string                    `json:"givenName"`
	FamilyName  string                    `json:"familyName"`
	Username    string                    `json:"username"`
	ClinicRoles []*commonModel.ClinicRole `json:"clinicRoles"`
	Activated   bool                      `json:"activated"`
}

type InviteUserInput struct {
	Email string `json:"email"`
}

type UpdateUserInput struct {
	UserID     string `json:"userId"`
	GivenName  string `json:"givenName"`
	FamilyName string `json:"familyName"`
}

type ActivationVerificationStatus string

const (
	ActivationVerificationStatusActive       ActivationVerificationStatus = "Active"
	ActivationVerificationStatusExpired      ActivationVerificationStatus = "Expired"
	ActivationVerificationStatusInvalid      ActivationVerificationStatus = "Invalid"
	ActivationVerificationStatusUnknownError ActivationVerificationStatus = "UnknownError"
)

var AllActivationVerificationStatus = []ActivationVerificationStatus{
	ActivationVerificationStatusActive,
	ActivationVerificationStatusExpired,
	ActivationVerificationStatusInvalid,
	ActivationVerificationStatusUnknownError,
}

func (e ActivationVerificationStatus) IsValid() bool {
	switch e {
	case ActivationVerificationStatusActive, ActivationVerificationStatusExpired, ActivationVerificationStatusInvalid, ActivationVerificationStatusUnknownError:
		return true
	}
	return false
}

func (e ActivationVerificationStatus) String() string {
	return string(e)
}

func (e *ActivationVerificationStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ActivationVerificationStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ActivationVerificationStatus", str)
	}
	return nil
}

func (e ActivationVerificationStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}