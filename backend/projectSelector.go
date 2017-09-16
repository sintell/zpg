package main

import (
	"fmt"
	"math/rand"
)

func getNextProjectAfterProjectComplete(characterID CharID) *Project {
	projects := getUnfinishedProjects(characterID)
	project := projects[rand.Intn(len(projects))]
	fmt.Println("GET PROJECT WITH STATUS", project.Status)
	return project
}

func getNextProjectAfterProjectStageComplete(characterID CharID) *Project {
	if rand.Intn(2) == 0 {
		return GetGlobalState().get(characterID).CharVarValue.CurrentProject
	}
	return getNextProjectAfterProjectComplete(characterID)
}

func getUnfinishedProjects(characterID CharID) []*Project {
	return Filter(GetGlobalState().get(characterID).Projects, func(p *Project) bool {
		return !p.isFinished()
	})
}

func Filter(vs []*Project, f func(project *Project) bool) []*Project {
	vsf := make([]*Project, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
