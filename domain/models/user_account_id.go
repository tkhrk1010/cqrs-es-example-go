package models

import (
	"fmt"
	"github.com/oklog/ulid/v2"
)

type UserAccountId struct {
	value string
}

func NewUserAccountId() *UserAccountId {
	id := ulid.Make()
	return &UserAccountId{value: id.String()}
}

func NewUserAccountIdFromString(value string) *UserAccountId {
	return &UserAccountId{value: value}
}

func ConvertUserAccountIdFromJSON(value map[string]interface{}) *UserAccountId {
	return NewUserAccountIdFromString(value["Value"].(string))
}

func (u *UserAccountId) GetValue() string {
	return u.value
}

func (u *UserAccountId) GetTypeName() string {
	return "user-account"
}

func (u *UserAccountId) AsString() string {
	return fmt.Sprintf("%s-%s", u.GetTypeName(), u.GetValue())
}

func (u *UserAccountId) String() string {
	return u.AsString()
}
