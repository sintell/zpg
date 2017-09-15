package main

import "github.com/go-pg/pg"

type CharVar struct {
	CharStatId       int         `json:"-"`
	CharStatValue    CharStat    `json:"-"`
	Stress           int         `json:"stress"`
	Resting          int         `json:"resting"`
	CurrentProjectId int         `json:"-"`
	CurrentProject   *Project    `json:"-"`
	SkillValueId     int         `json:"-"`
	SkillValue       *SkillValue `json:"skills"`
	Level            int         `json:"level"`
	Experience       int         `json:"exp"`
}

func createCharacter(userId int, name string, company string, prog int, testing int, analyze int) {
	GetDB().RunInTransaction(func(tx *pg.Tx) error {
		charStat := &CharStat{
			Name:    name,
			Company: company,
			UserId:  userId,
		}
		if err := tx.Insert(charStat); err != nil {
			if err := tx.Rollback(); err != nil {
				return err
			}
			return err
		}

		skillValue := &SkillValue{
			Prog:    prog,
			Testing: testing,
			Analyze: analyze,
		}
		if err := tx.Insert(skillValue); err != nil {
			if err := tx.Rollback(); err != nil {
				return err
			}
			return err
		}

		charVar := &CharVar{
			CharStatId:   charStat.Id,
			SkillValueId: skillValue.Id,
		}
		if err := tx.Insert(charVar); err != nil {
			if err := tx.Rollback(); err != nil {
				return err
			}
			return err
		}

		return tx.Commit()
	})
}
