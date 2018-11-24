package goscen

import (
    "fmt"
    "errors"
)

var (
    scoring = read()
)

func init() {
    scoring.check()
    scoring.load()
}

func WithLoaders(loadersExecutions LoadersMapping, scoringExecution goscenExecution) {
    for loaderID, loaderExecution := range loadersExecutions {
        for _, node := range scoring.Nodes {
            if node.ID == scoring.ID {
                node.Execution = scoringExecution
            }
            if node.ID == loaderID {
                if node.Execution != nil {
                    panic(errors.New(fmt.Sprintf("%s node already has an execution assigned", node.ID)))
                }
                node.Execution = loaderExecution
            }
        }
    }
}

func Run() {
    scoring.checkExecutions()
    scoring.success()
    scoring.serve()
}
