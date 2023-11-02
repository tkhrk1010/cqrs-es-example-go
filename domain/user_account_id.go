package domain

import "fmt"

type UserAccountId struct {
	value string
}

func NewUserAccountId(value string) *UserAccountId {
	return &UserAccountId{value: value}
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
