package main

func MainloopTick() {
	charIDs := GetGlobalState().getIds()

	for _, charID := range charIDs {
		tickUser(charID)
	}
}

func tickUser(charID CharID) {
	state := GetGlobalState().get(charID)
	if state.CharVarValue.Resting > 0 {
		state.CharVarValue.Resting--
		return
	}
	progress(state)
}

func progress(state *InternalState) {
	for _, project := range state.Projects {
		if project.Active {
			progressValues := project.ProgrValues
			requiredValues := project.ReqValues
			skills := state.CharVarValue.SkillValue
			switch project.Status {
			case ProjectStatus(Analyze):
				{
					if progressValues.Analyze+skills.Analyze >= requiredValues.Analyze {
						progressValues.Analyze = requiredValues.Analyze
						switchState(state, project, ProjectStatus(Analyze), ProjectStatus(Prog))
					} else {
						progressValues.Analyze += skills.Analyze
					}
				}
			case ProjectStatus(Prog):
				{
					if progressValues.Prog+skills.Prog >= requiredValues.Prog {
						progressValues.Prog = requiredValues.Prog
						switchState(state, project, ProjectStatus(Prog), ProjectStatus(Testing))
					} else {
						progressValues.Prog += skills.Prog
					}
				}
			case ProjectStatus(Testing):
				{
					if progressValues.Testing+skills.Testing >= requiredValues.Testing {
						progressValues.Testing = requiredValues.Testing
						switchState(state, project, ProjectStatus(Testing), ProjectStatus(Released))
					} else {
						progressValues.Testing += skills.Testing
					}
				}
			}
			project.ProgrValues = progressValues
			break
		}
	}
}

func switchState(state *InternalState, project *Project, statusTo ProjectStatus, statusFrom ProjectStatus) {
	project.Status = statusFrom
	newProject := getNextProjectAfterProjectStageComplete(state.CharStatValue.ID)
	if newProject.ID != project.ID {
		project.Active = false
		newProject.Active = true
	}
}
