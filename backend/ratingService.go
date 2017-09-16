package main

type RatingByExpResponse struct {
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Level    int    `json:"level"`
	CharId   CharID `json:"id"`
}

func getRatingByExp() ([]RatingByExpResponse, error) {
	cv := []CharVar{}
	err := GetDB().Model(&CharVar{}).
		Column("char_var.*", "CharStat.name", "CharStat.User.login").
		Order("level DESC").
		Limit(10).
		Select(&cv)

	if err != nil {
		return nil, err
	}

	rp := []RatingByExpResponse{}
	for _, charVar := range cv {
		rp = append(rp, RatingByExpResponse{Name: charVar.CharStat.Name, Level: charVar.Level, CharId: charVar.CharStatID, UserName: charVar.CharStat.User.Login})
	}

	return rp, nil
}
