package main

import (
	"math/rand"
)

func generateProject(charVar *CharVar) (*Project, error) {
	charValues := charVar.SkillValue
	if p, err := NewProject(charVar.CharStatID,
		generateProjectReqValues(charValues.Analyze),
		generateProjectReqValues(charValues.Prog),
		generateProjectReqValues(charValues.Testing),
		generateProjectName(),
		generateProjectDescription()); err == nil {
		return p, nil
	} else {
		return nil, err
	}
}

func generateProjectReqValues(value int) int {
	return rand.Intn(value) + value
}

func generateProjectName() string {
	return "new project ( " + string(rand.Intn(50)) + ")"
}

func generateProjectDescription() string {
	return "new project description blalbla ( " + string(rand.Intn(50)) + ")"
}

func needToGenerateProject(id CharID) bool {
	return (rand.Intn(50) > 40 || (len(GetGlobalState().get(id).Projects) < 3)) && (len(getUnfinishedProjects(id)) < 10)
}
