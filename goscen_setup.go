package goscen

import (
    "fmt"
    "errors"
    "github.com/mercadolibre/go-meli-toolkit/goutils/apierrors"
)

var (
    scoring = read()
)

func init() {
    scoring.check()
    scoring.load()
}

func AddLoaderExecution(loaderID string, execution func(inputs ...interface{}) ([]interface{}, apierrors.ApiError)) {
    for _, loader := range scoring.Loaders {
        if loader.ID == loaderID {
            if loader.Execution != nil {
                panic(errors.New(fmt.Sprintf("%s loader already has an execution assigned", loader.ID)))
            }
            loader.Execution = execution
        }
    }
}

func Run() {
    scoring.checkLoadersExecutions()
    scoring.success()
    scoring.serve()
}
