package goscen

import (
    "io/ioutil"
    "encoding/json"
    "github.com/mercadolibre/go-meli-toolkit/goutils/logger"
)

func init() {
    logger.Infof("Reading config")
    scoringConfig, err := readScoringConfig()
    if err != nil {
        panic(err)
    }

    logger.Infof("Validating config")
    if err := scoringConfig.check(); err != nil {
        panic(err)
    }

    logger.Infof("%s scoring with %d loaders successfully loaded", scoringConfig.ID, len(scoringConfig.Loaders))
}

func readScoringConfig() (*scoringConfig, error) {
    const (
        defaultConfigFileName = "./config.json"
    )
    bytes, err := ioutil.ReadFile(defaultConfigFileName)
    if err != nil {
        return nil, err
    }
    var scoringConfig scoringConfig
    if err := json.Unmarshal(bytes, &scoringConfig); err != nil {
        return nil, err
    }
    return &scoringConfig, nil
}
