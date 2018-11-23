package goscen

import (
    "fmt"
    "errors"
)

func (scoringConfig *scoringConfig) check() error {
    const (
        errNilScoringConfig = "scoring config is nil"
    )
    if scoringConfig == nil {
        return errors.New(errNilScoringConfig)
    }

    if err := scoringConfig.checkID(); err != nil {
        return err
    }

    if err := scoringConfig.checkType(); err != nil {
        return err
    }

    if err := scoringConfig.checkLoaders(); err != nil {
        return err
    }

    return nil
}

func (scoringConfig *scoringConfig) checkID() error {
    const (
        errEmptyScoringID = "scoring ID is empty"
    )
    if scoringConfig.ID == "" {
        return errors.New(errEmptyScoringID)
    }
    return nil
}

func (scoringConfig *scoringConfig) checkType() error {
    const (
        errInvalidScoringType  = "scoring type is not valid"
        scoringTypeComplete    = "complete"
        scoringTypeProgressive = "progressive"
    )
    switch scoringConfig.Type {
    case scoringTypeComplete:
    case scoringTypeProgressive:
    default:
        return errors.New(errInvalidScoringType)
    }
    return nil
}

func (scoringConfig *scoringConfig) checkLoaders() error {
    if err := scoringConfig.checkLoadersUniqueness(); err != nil {
        return err
    }

    if err := scoringConfig.checkLoadersDependencies(); err != nil {
        return err
    }

    for _, loader := range scoringConfig.Loaders {
        if err := loader.check(); err != nil {
            return err
        }
    }
    return nil
}

func (scoringConfig *scoringConfig) checkLoadersUniqueness() error {
    const (
        errDuplicatedLoader = "%s loader is duplicated"
    )
    uniqueness := make(map[string]int)
    for _, loader := range scoringConfig.Loaders {
        uniqueness[loader.ID]++
    }
    for id, count := range uniqueness {
        if count > 1 {
            return errors.New(fmt.Sprintf(errDuplicatedLoader, id))
        }
    }
    return nil
}

func (scoringConfig *scoringConfig) checkLoadersDependencies() error {
    if err := scoringConfig.checkDependenciesExistence(); err != nil {
        return err
    }

    if err := scoringConfig.checkDependenciesCycles(); err != nil {
        return err
    }

    return nil
}

func (scoringConfig *scoringConfig) checkDependenciesExistence() error {
    const (
        errDependencyNotExists = "dependency %s for loader %s doesn't exists"
    )
    dependencies := make(map[string]int)
    for _, loader := range scoringConfig.Loaders {
        dependencies[loader.ID]++
    }
    for _, loader := range scoringConfig.Loaders {
        for _, dependency := range loader.Dependencies {
            if dependencies[dependency] == 0 {
                return errors.New(fmt.Sprintf(errDependencyNotExists, dependency, loader.ID))
            }
        }
    }
    return nil
}

func (scoringConfig *scoringConfig) checkDependenciesCycles() error {
    const (
        errCyclicLoaderDependencies = "%s loader has cyclic dependencies"
    )
    for _, loader := range scoringConfig.Loaders {
        visited := make(map[string]bool)
        if dependencyCycle := scoringConfig.isDependencyCyclic(loader.ID, visited); dependencyCycle == true {
            return errors.New(fmt.Sprintf(errCyclicLoaderDependencies, loader.ID))
        }
        fmt.Println(visited)
    }
    return nil
}

func (scoringConfig *scoringConfig) isDependencyCyclic(dependency string, visited map[string]bool) bool {
    for _, loader := range scoringConfig.Loaders {
        if dependency == loader.ID {
            if len(loader.Dependencies) == 0 {
                return false
            }
            visited[dependency] = true
            for _, dep := range loader.Dependencies {
                if visited[dep] || scoringConfig.isDependencyCyclic(dep, visited) {
                    return true
                }
            }
        }
    }
    return false
}

func (loaderConfig *loaderConfig) check() error {
    const (
        errNilLoaderConfig = "loader config is nil"
    )
    if loaderConfig == nil {
        return errors.New(errNilLoaderConfig)
    }

    if err := loaderConfig.checkID(); err != nil {
        return err
    }

    if err := loaderConfig.checkType(); err != nil {
        return err
    }

    return nil
}

func (loaderConfig *loaderConfig) checkID() error {
    const (
        errEmptyLoaderID = "loader ID is empty"
    )
    if loaderConfig.ID == "" {
        return errors.New(errEmptyLoaderID)
    }
    return nil
}

func (loaderConfig *loaderConfig) checkType() error {
    const (
        errInvalidLoaderType = "loader type is not valid"
        loaderTypeAPI        = "API"
    )
    switch loaderConfig.Type {
    case loaderTypeAPI:
    default:
        return errors.New(errInvalidLoaderType)
    }
    return nil
}
