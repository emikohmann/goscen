package goscen

import (
    "fmt"
    "strings"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "github.com/gin-gonic/gin"
    "github.com/mercadolibre/go-meli-toolkit/goutils/apierrors"
    "github.com/mercadolibre/go-meli-toolkit/gingonic/mlhandlers"
)

func read() *goscenScoring {
    bytes, err := ioutil.ReadFile("./config.json")
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

func (scoring *goscenScoring) success() {
    fmt.Println(goscen)
    fmt.Println(fmt.Sprintf("%s SCORING", strings.ToUpper(scoring.ID)))
    for _, loader := range scoring.Loaders {
        fmt.Println(fmt.Sprintf("\n\tWITH LOADER [%s]", strings.ToUpper(loader.ID)))
    }
    fmt.Println(fmt.Sprintf("\nSuccessfully initialized.\n"))
}

func (scoring *goscenScoring) serve() {
    router := mlhandlers.DefaultMeliRouter()

    router.GET("/ping", func(c *gin.Context) {
        c.String(http.StatusOK, "pong")
    })

    router.POST("/scoring", func(c *gin.Context) {
        if apiErr := scoring.run(); apiErr != nil {
            c.JSON(apiErr.Status(), apiErr)
            return
        }
        c.JSON(http.StatusCreated, nil)
    })

    router.Run(":8080")
}

func (scoring *goscenScoring) run() apierrors.ApiError {
    executions := make(map[*goscenLoader]bool)
    for _, loader := range scoring.Loaders {
        if apiErr := loader.run(executions); apiErr != nil {
            return apiErr
        }
    }
    return nil
}

func (loader *goscenLoader) run(executions map[*goscenLoader]bool) apierrors.ApiError {
    if executions[loader] == true {
        return nil
    }
    for _, dependencyLoader := range loader.DependenciesLoaders {
        if executions[dependencyLoader] == false {
            if apiErr := dependencyLoader.run(executions); apiErr != nil {
                return apiErr
            }
        }
    }
    inputs := make([]interface{}, 0)
    for _, dependencyLoader := range loader.DependenciesLoaders {
        inputs = append(inputs, dependencyLoader.ExecutionResult)
    }
    // fmt.Println("Running", loader.ID, "loader with inputs", inputs)
    res, apiErr := loader.Execution(inputs...)
    if apiErr != nil {
        return apiErr
    }
    loader.ExecutionResult = res
    executions[loader] = true
    return nil
}
