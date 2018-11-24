package goscen

import (
    "fmt"
    "errors"
    "strings"
)

func (scoring *goscenScoring) check() {
    scoring.checkID()
    scoring.checkMode()
    scoring.checkEntryPoint()
    scoring.checkNodes()
}

func (scoring *goscenScoring) checkID() {
    if scoring.ID == "" {
        panic(errors.New("scoring ID is empty"))
    }
}

func (scoring *goscenScoring) checkMode() {
    for _, scoringMode := range scoringModes {
        if scoring.Mode == scoringMode {
            return
        }
    }

    panic(errors.New(fmt.Sprintf("%s scoring mode is not valid", scoring.ID)))
}

func (scoring *goscenScoring) checkEntryPoint() {
    if strings.HasPrefix(scoring.EntryPoint, "/") == false ||
        len(strings.Split(scoring.EntryPoint, "/")[1]) == 0 {
        panic(errors.New(fmt.Sprintf("%s scoring entry point is not valid", scoring.ID)))
    }
}

func (scoring *goscenScoring) checkNodes() {
    scoring.checkNodesUniqueness()
    scoring.checkNodesDependencies()

    for _, node := range scoring.Nodes {
        node.check()
    }
}

func (scoring *goscenScoring) checkNodesUniqueness() {
    uniqueness := make(map[string]int)
    for _, node := range scoring.Nodes {
        uniqueness[node.ID]++
    }

    for id, count := range uniqueness {
        if count > 1 {
            panic(errors.New(fmt.Sprintf("%s node is duplicated", id)))
        }
    }
}

func (scoring *goscenScoring) checkNodesDependencies() {
    scoring.checkDependenciesExistence()
    scoring.checkDependenciesCycles()
}

func (scoring *goscenScoring) checkDependenciesExistence() {
    dependencies := make(map[string]int)
    for _, node := range scoring.Nodes {
        dependencies[node.ID]++
    }

    for _, node := range scoring.Nodes {
        for _, dependency := range node.DependenciesID {
            if dependencies[dependency] == 0 {
                panic(errors.New(fmt.Sprintf("dependency %s for node %s doesn't exists", dependency, node.ID)))
            }
        }
    }
}

func (scoring *goscenScoring) checkDependenciesCycles() {
    for _, node := range scoring.Nodes {
        visited := make(map[string]bool)
        if dependencyCycle := scoring.isDependencyCyclic(node.ID, visited); dependencyCycle == true {
            panic(errors.New(fmt.Sprintf("%s node has cyclic dependencies", node.ID)))
        }
    }
}

func (scoring *goscenScoring) isDependencyCyclic(dependency string, visited map[string]bool) bool {
    for _, node := range scoring.Nodes {
        if dependency == node.ID {
            if len(node.DependenciesID) == 0 {
                return false
            }

            visited[dependency] = true
            for _, dependencyID := range node.DependenciesID {
                if visited[dependencyID] || scoring.isDependencyCyclic(dependencyID, visited) {
                    return true
                }
            }
        }
    }
    return false
}

func (node *goscenNode) check() {
    node.checkID()
    node.checkType()
}

func (node *goscenNode) checkID() {
    if node.ID == "" {
        panic(errors.New("node ID is empty"))
    }
}

func (node *goscenNode) checkType() {
    for _, nodeType := range nodeTypes {
        if node.Type == nodeType {
            return
        }
    }

    panic(errors.New(fmt.Sprintf("%s node type is not valid", node.ID)))
}

func (scoring *goscenScoring) checkExecutions() {
    for _, node := range scoring.Nodes {
        if node.Execution == nil {
            panic(errors.New(fmt.Sprintf("%s node has nil execution", node.ID)))
        }
    }
}
