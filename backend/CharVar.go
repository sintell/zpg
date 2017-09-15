package main

import "github.com/go-pg/pg"

type CharVar struct {
	CharStatID       CharID      `json:"-"`
	CharStat         CharStat    `json:"-"`
	Stress           int         `json:"stress"`
	Resting          int         `json:"resting"`
	CurrentProjectID int         `json:"-"`
	CurrentProject   *Project    `json:"-"`
	SkillValueID     int         `json:"-"`
	SkillValue       *SkillValue `json:"skills"`
	Level            int         `json:"level"`
	Experience       int         `json:"exp"`
}

func createCharacter(userID int, name string, company string, prog int, testing int, analyze int) (
	*CharStat, *CharVar, error) {

	var charStat *CharStat
	var charVar *CharVar
	if err := GetDB().RunInTransaction(func(tx *pg.Tx) error {
		charStat = &CharStat{
			Name:    name,
			Company: company,
			UserID:  userID,
		}
		if err := tx.Insert(charStat); err != nil {
			return err
		}

		skillValue := &SkillValue{
			Prog:    prog,
			Testing: testing,
			Analyze: analyze,
		}
		if err := tx.Insert(skillValue); err != nil {
			return err
		}

		charVar = &CharVar{
			CharStatID:   charStat.ID,
			SkillValueID: skillValue.ID,
		}
		if err := tx.Insert(charVar); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, nil, err
	}

	return charStat, charVar, nil
}
