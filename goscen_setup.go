package goscen

import (
    "github.com/mercadolibre/go-meli-toolkit/goutils/logger"
    "fmt"
)

var (
    scoringTypes = []string{
        "complete",
        "progressive",
    }

    loaderTypes = []string{
        "API",
    }
)

func init() {
    scoring := read()
    scoring.check()
    scoring.load()
    logger.Infof("%s scoring successfully initialized!", scoring.ID)
    fmt.Println(goscen)
}
