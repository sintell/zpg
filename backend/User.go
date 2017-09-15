package main

type User struct {
	ID       int
	Login    string `json:"login,omitempty" validate:"required" sql:", unique"`
	Password string `json:"password,omitempty" validate:"required"`
}

func NewUser(u *User) error {
	if err := GetDB().Insert(u); err != nil {
		return err
	}

	return nil
}

func (u *User) getCharacter() (*CharStat, error) {
	cs := &CharStat{}

	if err := GetDB().Model(cs).Select(&CharStat{UserID: u.ID}); err != nil {
		return nil, err
	}

	return cs, nil
}
