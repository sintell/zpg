package main

type User struct {
	Id       int
	Login    string `json:"login,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}

func NewUser(u *User) error {
	if err := GetDB().Insert(u); err != nil {
		return err
	}

	return nil
}
