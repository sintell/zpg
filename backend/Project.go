package main

import "math/rand"

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
	progrValues := p.ProgrValues
	reqValues := p.ReqValues

	return (progrValues.Analyze >= reqValues.Analyze) &&
		(progrValues.Prog >= reqValues.Prog) &&
		(progrValues.Testing >= reqValues.Testing)
}

func NewProject(id CharID) (*Project, error) {
	req := &SkillValue{
		Analyze: int(rand.Int31n(10)),
		Prog:    int(rand.Int31n(10)),
		Testing: int(rand.Int31n(10)),
	}
	progr := &SkillValue{}

	if err := GetDB().Insert(req, progr); err != nil {
		return nil, err
	}

	return &Project{
		CharStatID:    id,
		Name:          "Project 1",
		Description:   "Description 1",
		ReqValuesID:   req.ID,
		ReqValues:     req,
		ProgrValuesID: progr.ID,
		ProgrValues:   progr,
	}, nil
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
