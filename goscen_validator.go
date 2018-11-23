package goscen

import (
    "fmt"
    "errors"
)

func (scoring *goscenScoring) check() {
    scoring.checkID()
    scoring.checkType()
    scoring.checkLoaders()
}

func (scoring *goscenScoring) checkID() {
    if scoring.ID == "" {
        panic(errors.New("scoring ID is empty"))
    }
}

func (scoring *goscenScoring) checkType() {
    for _, scoringType := range scoringTypes {
        if scoring.Type == scoringType {
            return
        }
    }
    panic(errors.New("scoring type is not valid"))
}

func (scoring *goscenScoring) checkLoaders() {
    scoring.checkLoadersUniqueness()
    scoring.checkLoadersDependencies()
    for _, loader := range scoring.Loaders {
        loader.check()
    }
}

func (scoring *goscenScoring) checkLoadersUniqueness() {
    uniqueness := make(map[string]int)
    for _, loader := range scoring.Loaders {
        uniqueness[loader.ID]++
    }
    for id, count := range uniqueness {
        if count > 1 {
            panic(errors.New(fmt.Sprintf("%s loader is duplicated", id)))
        }
    }
}

func (scoring *goscenScoring) checkLoadersDependencies() {
    scoring.checkDependenciesExistence()
    scoring.checkDependenciesCycles()
}

func (scoring *goscenScoring) checkDependenciesExistence() {
    dependencies := make(map[string]int)
    for _, loader := range scoring.Loaders {
        dependencies[loader.ID]++
    }
    for _, loader := range scoring.Loaders {
        for _, dependency := range loader.DependenciesID {
            if dependencies[dependency] == 0 {
                panic(errors.New(fmt.Sprintf("dependency %s for loader %s doesn't exists", dependency, loader.ID)))
            }
        }
    }
}

func (scoring *goscenScoring) checkDependenciesCycles() {
    for _, loader := range scoring.Loaders {
        visited := make(map[string]bool)
        if dependencyCycle := scoring.isDependencyCyclic(loader.ID, visited); dependencyCycle == true {
            panic(errors.New(fmt.Sprintf("%s loader has cyclic dependencies", loader.ID)))
        }
    }
}

func (scoring *goscenScoring) isDependencyCyclic(dependency string, visited map[string]bool) bool {
    for _, loader := range scoring.Loaders {
        if dependency == loader.ID {
            if len(loader.DependenciesID) == 0 {
                return false
            }
            visited[dependency] = true
            for _, dep := range loader.DependenciesID {
                if visited[dep] || scoring.isDependencyCyclic(dep, visited) {
                    return true
                }
            }
        }
    }
    return false
}

func (loader *goscenLoader) check() {
    loader.checkID()
    loader.checkType()
}

func (loader *goscenLoader) checkID() {
    if loader.ID == "" {
        panic(errors.New("loader ID is empty"))
    }
}

func (loader *goscenLoader) checkType() {
    for _, loaderType := range loaderTypes {
        if loader.Type == loaderType {
            return
        }
    }
    panic(errors.New("loader type is not valid"))
}
