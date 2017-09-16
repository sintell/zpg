package main

import (
	"fmt"
	"math/rand"

	"github.com/go-pg/pg"
)

type Project struct {
	ID            int           `json:"id"`
	CharStatID    CharID        `json:"-"`
	Name          string        `json:"name"`
	Description   string        `json:"desc"`
	ReqValuesID   int           `json:"-"`
	ReqValues     *SkillValue   `json:"req_values"`
	ProgrValuesID int           `json:"-"`
	ProgrValues   *SkillValue   `json:"progress"`
	Active        bool          `json:"active"`
	Status        ProjectStatus `json:"status"`
}

type ProjectStatus string

const (
	Todo     ProjectStatus = "TODO"
	Analyze                = "ANALYZE"
	Prog                   = "PROG"
	Testing                = "TEST"
	Released               = "RELEASED"
)

func (p Project) isFinished() bool {
	return p.Status == Released
}

func NewProject(id CharID) (*Project, error) {
	req := &SkillValue{
		Analyze: int(rand.Int31n(10)),
		Prog:    int(rand.Int31n(10)),
		Testing: int(rand.Int31n(10)),
	}
	progr := &SkillValue{}
	var p *Project
	if err := GetDB().RunInTransaction(func(tx *pg.Tx) error {
		if err := tx.Insert(req, progr); err != nil {
			return err
		}
		p = &Project{
			CharStatID:    id,
			Name:          fmt.Sprintf("%s %d", "Project", req.ID),
			Description:   fmt.Sprintf("%s %d", "Description", req.ID),
			ReqValuesID:   req.ID,
			ReqValues:     req,
			ProgrValuesID: progr.ID,
			ProgrValues:   progr,
			Status:        Todo,
		}
		if err := tx.Insert(p); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return p, nil
}

func CreateProjectsFor(id CharID) ([]*Project, error) {
	projects := make([]*Project, 3)
	for idx := range projects {
		p, err := NewProject(id)
		if err != nil {
			return nil, err
		}
		projects[idx] = p
	}

	return projects, nil
}
