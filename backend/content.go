package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var projectNames ProjectNames

type CompanyNamesText struct {
	Parts []string
}

type ProjectNames struct {
	Action      []string `json:"action,omitempty"`
	Description []string `json:"description,omitempty"`
	Subject     []string `json:"subject,omitempty"`
}

type EventsText struct {
	Parts []struct {
	}
}

func LoadProjectNamesFrom(path string) error {
	return ReadConfig(&projectNames, path)
}

func GenerateProjectName() string {
	return fmt.Sprintf("%s %s",
		getRandomPart(projectNames.Action),
		joinWithAlignedKinds(getRandomPart(projectNames.Description),
			getRandomPart(projectNames.Subject)))
}

func joinWithAlignedKinds(dep string, source string) string {
	switch {
	case strings.HasSuffix(source, "а"):
		dep = dep + "ого"

	case strings.HasSuffix(source, "и"):
		dep = dep + "ой"
	}

	return fmt.Sprintf("%s %s", strings.ToLower(dep), strings.ToLower(source))
}

func getRandomPart(parts []string) string {
	return parts[rand.Intn(len(parts))]
}
