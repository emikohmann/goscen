package goscen

import (
    "io/ioutil"
    "encoding/json"
)

func read() *goscenScoring {
    bytes, err := ioutil.ReadFile(defaultConfigFile)
    if err != nil {
        panic(err)
    }
    var scoring goscenScoring
    if err := json.Unmarshal(bytes, &scoring); err != nil {
        panic(err)
    }
    return &scoring
}

func (scoring *goscenScoring) load() {
    loaders := make(map[string]*goscenLoader)
    for _, loader := range scoring.Loaders {
        loaders[loader.ID] = loader
    }
    for _, loader := range scoring.Loaders {
        loader.DependenciesLoaders = make([]*goscenLoader, 0)
        for _, dependencyID := range loader.DependenciesID {
            loader.DependenciesLoaders = append(loader.DependenciesLoaders, loaders[dependencyID])
        }
    }
}
