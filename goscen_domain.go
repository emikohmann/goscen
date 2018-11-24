package goscen

import (
    "github.com/mercadolibre/go-meli-toolkit/goutils/apierrors"
)

const (
    scoringModePassive = "passive"
    scoringModeActive  = "active"
)

var (
    scoringModes = []string{
        scoringModePassive,
        scoringModeActive,
    }

    nodeTypes = []string{
        "",
    }
)

type goscenScoring struct {
    ID             string        `json:"id"`
    Mode           string        `json:"mode"`
    Type           string        `json:"type"`
    EntryPoint     string        `json:"entry_point"`
    DependenciesID []string      `json:"dependencies"`
    Nodes          []*goscenNode `json:"loaders"`
}

type goscenNode struct {
    ID                string          `json:"id"`
    Type              string          `json:"type"`
    DependenciesID    []string        `json:"dependencies"`
    DependenciesNodes []*goscenNode   `json:"-"`
    Execution         goscenExecution `json:"-"`
    ExecutionResult   []interface{}   `json:"-"`
}

type goscenExecution func(...interface{}) ([]interface{}, apierrors.ApiError)

type LoadersMapping map[string]goscenExecution
