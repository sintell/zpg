package main

import (
	"fmt"
)

func MainloopTick() {
	charIDs := GetGlobalState().getIds()

	for _, charID := range charIDs {
		tickUser(charID)
	}
}

func tickUser(charID CharID) {
	state := GetGlobalState().get(charID)
	randomEvents(state)
	fmt.Println(state)
	if state.CharVarValue.Resting > 0 {
		state.CharVarValue.Resting--
		if state.CharVarValue.Stress > 10 {
			state.CharVarValue.Stress -= 10
		} else {
			state.CharVarValue.Stress = 0
		}
		return
	}
	progress(state)
}

func randomEvents(state *InternalState) {

}

func progress(state *InternalState) {
	for _, project := range state.Projects {
		if project.Active {
			progressValues := project.ProgrValues
			requiredValues := project.ReqValues
			skills := state.CharVarValue.SkillValue
			switch project.Status {
			case Analyze:
				{
					if progressValues.Analyze+skills.Analyze >= requiredValues.Analyze {
						progressValues.Analyze = requiredValues.Analyze
						switchState(state, project, Analyze, Prog)
					} else {
						progressValues.Analyze += skills.Analyze
					}
					break
				}
			case Prog:
				{
					if progressValues.Prog+skills.Prog >= requiredValues.Prog {
						progressValues.Prog = requiredValues.Prog
						switchState(state, project, Prog, Testing)
					} else {
						progressValues.Prog += skills.Prog
					}
					break
				}
			case Testing:
				{
					if progressValues.Testing+skills.Testing >= requiredValues.Testing {
						progressValues.Testing = requiredValues.Testing
						switchState(state, project, Testing, Released)
					} else {
						progressValues.Testing += skills.Testing
					}
					break
				}
			}
			project.ProgrValues = progressValues
			break
		}
	}
}

func switchState(state *InternalState, project *Project, statusFrom ProjectStatus, statusTo ProjectStatus) {
	project.Status = statusTo
	var newProject *Project
	if statusTo == ProjectStatus(Released) {
		state.CharVarValue.Experience += countExpFromProject(project, state.CharVarValue)
		if state.CharVarValue.Experience >= 100 {
			levelUp(state)
		}
		newProject = getNextProjectAfterProjectComplete(state.CharStatValue.ID)
	} else {
		newProject = getNextProjectAfterProjectStageComplete(state.CharStatValue.ID)
	}
	if newProject.ID != project.ID {
		project.Active = false
		newProject.Active = true
		if newProject.Status == ProjectStatus(Todo) {
			newProject.Status = ProjectStatus(Analyze)
			state.EventQueue.push(
				NewEvent(
					New_project,
					fmt.Sprintf("Проект %s завершен. Проект %s взят в работу",
						project.Name, newProject.Name)))
		} else {
			state.EventQueue.push(
				NewEvent(
					Change_project_status,
					fmt.Sprintf("Стадия %s проекта %s завершена. Переключение на проект %s",
						statusFrom, project.Name, newProject.Name)))
		}
	} else {
		state.EventQueue.push(
			NewEvent(
				Change_project_status,
				fmt.Sprintf("Стадия %s проекта %s завершена",
					statusFrom, project.Name)))
	}
}

func levelUp(state *InternalState) {
	state.CharVarValue.Level += 1
	state.CharVarValue.Experience = 0
}

func countExpFromProject(project *Project, charVar *CharVar) int {
	rv := project.ReqValues
	cv := charVar.SkillValue
	return countExpFromSkillValue(rv.Analyze, cv.Analyze) +
		countExpFromSkillValue(rv.Prog, cv.Prog) +
		countExpFromSkillValue(rv.Testing, cv.Testing)
}

func countExpFromSkillValue(reqValue, charValue int) int {
	return reqValue/charValue + 1
}
