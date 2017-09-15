package main

type Project struct {
	Id            int         `json:"id"`
	CharStatId    *CharStat   `json:"-"`
	Name          string      `json:"name"`
	Description   string      `json:"desc"`
	ReqValuesId   int         `json:"-"`
	ReqValues     *SkillValue `json:"req_values"`
	ProgrValuesId int         `json:"-"`
	ProgrValues   *SkillValue `json:"progress"`
}
