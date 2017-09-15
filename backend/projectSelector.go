package main

import "math/rand"

func getNextProjectAfterProjectComplete(characterId int) *Project {
	projects := getUnfinishedProjects(characterId)
	return projects[rand.Intn(len(projects))]
}

func getNextProjectAfterProjectStageComplete(characterId int) *Project {
	if rand.Intn(2) == 0 {
		return GetGlobalState().get(characterId).CharVarValue.CurrentProject
	} else {
		return getNextProjectAfterProjectComplete(characterId)
	}
}

func getUnfinishedProjects(characterId int) []*Project {
	return Filter(GetGlobalState().get(characterId).Projects, func(p *Project) bool {
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
