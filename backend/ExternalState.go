package main

type ExternalState struct {
	Character struct {
		CharStat
		CharVar
	} `json:"char"`
	Projects      []*Project      `json:"projects"`
	ActiveEffects []*ActiveEffect `json:"effects"`
}

func from(InternalStateValue InternalState) ExternalState {
	result := ExternalState{}
	result.Character.CharStat = InternalStateValue.CharStatValue
	result.Character.CharVar = InternalStateValue.CharVarValue
	result.Projects = InternalStateValue.Projects
	result.ActiveEffects = InternalStateValue.ActiveEffects
	for _, effect := range result.ActiveEffects {
		result.Character.CharVar.SkillValue.Prog += effect.Effect.Effect.Prog
		result.Character.CharVar.SkillValue.Testing += effect.Effect.Effect.Testing
		result.Character.CharVar.SkillValue.Analyze += effect.Effect.Effect.Analyze
	}
	return result
}
