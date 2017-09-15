package main

type Project struct {
	ID            int           `json:"id"`
	CharStatId    *CharStat     `json:"-"`
	Name          string        `json:"name"`
	Description   string        `json:"desc"`
	ReqValuesId   int           `json:"-"`
	ReqValues     *SkillValue   `json:"req_values"`
	ProgrValuesId int           `json:"-"`
	ProgrValues   *SkillValue   `json:"progress"`
	Active        bool          `json:"active"`
	Status        ProjectStatus `json:"status"`
}

type ProjectStatus string

const (
	Todo     ProjectStatus = "TODO"
	Analyze  ProjectStatus = "ANALYZE"
	Prog     ProjectStatus = "PROG"
	Testing  ProjectStatus = "TEST"
	Released ProjectStatus = "RELEASED"
)

func (p Project) isFinished() bool {
	progrValues := p.ProgrValues
	reqValues := p.ReqValues

	return (progrValues.Analyze >= reqValues.Analyze) &&
		(progrValues.Prog >= reqValues.Prog) &&
		(progrValues.Testing >= reqValues.Testing)
}

func CreateProjectsFor(id CharID) []*Project {
	return []*Project{
		&Project{
			Name:        "Project 1",
			Description: "Description 1",
			ReqValues:   &SkillValue{Analyze: 10, Prog: 10, Testing: 10},
		},
		&Project{
			Name:        "Project 2",
			Description: "Description 2",
			ReqValues:   &SkillValue{Analyze: 10, Prog: 10, Testing: 10},
		},
		&Project{
			Name:        "Project 3",
			Description: "Description 3",
			ReqValues:   &SkillValue{Analyze: 10, Prog: 10, Testing: 10},
		},
	}
}
