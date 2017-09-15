package main

type Project struct {
	Id            int
	ReqValuesId   int
	ReqValues     *SkillValue
	ProgrValuesId int
	ProgrValues   *SkillValue
}
