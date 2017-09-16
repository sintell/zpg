package main

type RatingByExpResponse struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
}

func getRatingByExp() ([]RatingByExpResponse, error) {
	cv := []CharVar{}
	err := GetDB().Model(&CharVar{}).
		Column("char_var.*", "CharStat.name").
		Order("level DESC").
		Limit(10).
		Select(&cv)

	if err != nil {
		return nil, err
	}

	rp := []RatingByExpResponse{}
	for _, charVar := range cv {
		rp = append(rp, RatingByExpResponse{Name: charVar.CharStat.Name, Level: charVar.Level})
	}

	return rp, nil
}
