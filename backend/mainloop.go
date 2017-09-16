package main

import (
	"fmt"
	"math"
	"math/rand"
)

func MainloopTick() {
	charIDs := GetGlobalState().getIds()

	for _, charID := range charIDs {
		tickUser(charID)
	}
}

func tickUser(charID CharID) {
	state := GetGlobalState().get(charID)
	state.CharVarValue.Stress ++
	randomEvents(state)

	if needToGenerateProject(charID) {
		if p, err := generateProject(state.CharVarValue); p != nil && err == nil {
			state.Projects = append(state.Projects, p)
		} else {
			fmt.Println(err.Error())
		}
	}
	expireEvents(state)
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
	if state.CharVarValue.Stress >= 100 {
		state.CharVarValue.Stress = 100
		state.CharVarValue.Resting = 10
		state.EventQueue.push(
			NewEvent(
				Forced_rest,
				"Слишком высокий стресс. Персонаж идет отдыхать"))
		return 
	}
	progress(state)
}

func expireEvents(state *InternalState) {
	newActiveEffects := make([]*ActiveEffect, 0)
	for _, value := range state.ActiveEffects {
		value.Expires--
		if value.Expires == 0 {
			state.EventQueue.push(
				NewEvent(
					Effect_expires,
					"ифект экспайрс"))
		} else {
			newActiveEffects = append(newActiveEffects, value)
		}
	}

	state.ActiveEffects = newActiveEffects
}

func randomEvents(state *InternalState) {
	if rand.Intn(6) == 0 {
		state.EventQueue.push(
			NewEvent(
				New_event,
				"Огонь ивент"))
	}

	if rand.Intn(10) == 0 {
		state.EventQueue.push(
			NewEvent(
				New_event,
				"Вы словили баф на статы"))

		var effect Effect

		GetDB().Model(&effect).Where("id = ?", 0).Select()

		state.ActiveEffects = append(state.ActiveEffects,
			&ActiveEffect{
				CharStat:   *state.CharStatValue,
				CharStatID: state.CharStatValue.ID,
				Expires:    4,
				EffectID:   0,
				Effect:     &effect})

		state.EventQueue.push(
			NewEvent(
				New_effect,
				effect.Description))
	}

	if rand.Intn(14) == 0 {
		state.EventQueue.push(
			NewEvent(
				New_event,
				"Вы словили дебаф на статы"))
		var effect Effect

		GetDB().Model(&effect).Where("id = ?", 1).Select()

		state.ActiveEffects = append(state.ActiveEffects,
			&ActiveEffect{
				CharStat:   *state.CharStatValue,
				CharStatID: state.CharStatValue.ID,
				Expires:    5,
				EffectID:   0,
				Effect:     &effect})

		state.EventQueue.push(
			NewEvent(
				New_effect,
				effect.Description))
	}
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
						progressValues.Analyze += int(math.Max(float64(skills.Analyze), 0))
					}
					break
				}
			case Prog:
				{
					if progressValues.Prog+skills.Prog >= requiredValues.Prog {
						progressValues.Prog = requiredValues.Prog
						switchState(state, project, Prog, Testing)
					} else {
						progressValues.Prog += int(math.Max(float64(skills.Prog), 0))
					}
					break
				}
			case Testing:
				{
					if progressValues.Testing+skills.Testing >= requiredValues.Testing {
						progressValues.Testing = requiredValues.Testing
						switchState(state, project, Testing, Released)
					} else {
						progressValues.Testing += int(math.Max(float64(skills.Testing), 0))
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
	if statusTo == Released {
		state.EventQueue.push(
			NewEvent(
				Complete_project,
				fmt.Sprintf("Проект %s завершен.",
					project.Name)))
		state.CharVarValue.Stress /= 2
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
		state.CharVarValue.CurrentProject = newProject
		state.CharVarValue.CurrentProjectID = newProject.ID
		if newProject.Status == ProjectStatus(Todo) {
			newProject.Status = ProjectStatus(Analyze)
			state.EventQueue.push(
				NewEvent(
					New_project,
					fmt.Sprintf("Проект %s взят в работу",
						newProject.Name)))
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
	state.EventQueue.push(
		NewEvent(
			Level_up,
			"Уровень персонажа повышен"))
	state.CharVarValue.Stress = 0
	progIncrease := rand.Intn(5)
	testIncrease := rand.Intn(5)
	analyzeIncrease := rand.Intn(5)
	state.CharVarValue.SkillValue.Prog += progIncrease;
	if progIncrease > 0 {
		state.EventQueue.push(
			NewEvent(
				Level_up,
				fmt.Sprintf("Навык программирования повышен на %s", progIncrease)))
	}
	state.CharVarValue.SkillValue.Testing += testIncrease;
	if testIncrease > 0 {
		state.EventQueue.push(
			NewEvent(
				Level_up,
				fmt.Sprintf("Навык тестирования повышен на %s", testIncrease)))
	}
	state.CharVarValue.SkillValue.Analyze += analyzeIncrease;
	if analyzeIncrease > 0 {
		state.EventQueue.push(
			NewEvent(
				Level_up,
				fmt.Sprintf("Навык аналитики повышен на %s", analyzeIncrease)))
	}
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
