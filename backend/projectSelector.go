package main

import (
	"fmt"
	"math/rand"
)

func getNextProjectAfterProjectComplete(characterID CharID) *Project {
	projects := getUnfinishedProjects(characterID)
	projects = filterByWIPLimits(projects, characterID)
	project := projects[rand.Intn(len(projects))]
	fmt.Println("GET PROJECT WITH STATUS", project.Status)
	return project
}

func getNextProjectAfterProjectStageComplete(characterID CharID) *Project {
	projects := getUnfinishedProjects(characterID)
	projects = filterByWIPLimits(projects, characterID)
	currentProjectAllowed := false
	currentProject := GetGlobalState().get(characterID).CharVarValue.CurrentProject
	for _, value := range projects {
		if value.ID == currentProject.ID {
			currentProjectAllowed = true
			break
		}
	}
	if currentProjectAllowed && rand.Intn(2) == 0 {
		return GetGlobalState().get(characterID).CharVarValue.CurrentProject
	}
	return getNextProjectAfterProjectComplete(characterID)
}

func filterByWIPLimits(src []*Project, characterId CharID) []*Project{
	rejectedStatuses := make([]ProjectStatus, 0)
	counterAnalyze := 0
	counterProg := 0
	counterTest := 0
	WIPLimits := GetGlobalState().get(characterId).CharVarValue.WIPLimits
	for _, value := range src {
		if value.Status == Analyze {
			counterAnalyze++
		}
		if value.Status == Prog {
			counterProg++
		}
		if (value.Status == Testing) {
			counterTest++
		}
	}

	if counterAnalyze >= WIPLimits.Analyze {
		rejectedStatuses = append(rejectedStatuses, Todo)
	}
	if counterProg >= WIPLimits.Prog {
		rejectedStatuses = append(rejectedStatuses, Analyze)
	}
	if counterTest >= WIPLimits.Testing {
		rejectedStatuses = append(rejectedStatuses, Prog)
	}

	result := make([]*Project, 0)
	for _, value := range src {
		rejected := false
		for _, status := range rejectedStatuses {
			if value.Status == status {
				rejected = true
			}
		}
		if !rejected {
			result = append(result, value)
		}
	}
	return result
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
