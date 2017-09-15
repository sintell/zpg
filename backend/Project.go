package main

type Project struct {
	Id            int           `json:"id"`
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

func (p Project) isFinished() bool {
	progrValues := p.ProgrValues
	reqValues := p.ReqValues

	return (progrValues.Analyze >= reqValues.Analyze) &&
		(progrValues.Prog >= reqValues.Prog) &&
		(progrValues.Testing >= reqValues.Testing)
}

type ProjectStatus string

const (
	Todo     ProjectStatus = "TODO"
	Analyze  ProjectStatus = "ANALYZE"
	Prog     ProjectStatus = "PROG"
	Testing  ProjectStatus = "TEST"
	Released ProjectStatus = "RELEASED"
)
