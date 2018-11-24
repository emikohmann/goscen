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
    "errors"
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
    scoring.Nodes = append(scoring.Nodes, &goscenNode{
        ID:             scoring.ID,
        Type:           scoring.Type,
        DependenciesID: scoring.DependenciesID,
    })
    return &scoring
}

func (scoring *goscenScoring) load() {
    nodes := make(map[string]*goscenNode)
    for _, node := range scoring.Nodes {
        nodes[node.ID] = node
    }
    for _, node := range scoring.Nodes {
        node.DependenciesNodes = make([]*goscenNode, 0)
        for _, dependencyID := range node.DependenciesID {
            node.DependenciesNodes = append(node.DependenciesNodes, nodes[dependencyID])
        }
    }
}

func (scoring *goscenScoring) success() {
    fmt.Println(goscen)
    fmt.Println(fmt.Sprintf("%s SCORING", strings.ToUpper(scoring.ID)))
    for _, node := range scoring.Nodes {
        fmt.Println(fmt.Sprintf("\n\tWITH NODE [%s]", strings.ToUpper(node.ID)))
    }
    fmt.Println(fmt.Sprintf("\nSuccessfully initialized.\n"))
}

func (scoring *goscenScoring) serve() {
    router := mlhandlers.DefaultMeliRouter()

    router.GET("/ping", func(c *gin.Context) {
        c.String(http.StatusOK, "pong")
    })

    router.POST("/scoring", func(c *gin.Context) {
        res, apiErr := scoring.run()
        if apiErr != nil {
            c.JSON(apiErr.Status(), apiErr)
            return
        }
        c.JSON(http.StatusCreated, res)
    })

    router.Run(":8080")
}

func (scoring *goscenScoring) run() ([]interface{}, apierrors.ApiError) {
    executions := make(map[*goscenNode]bool)
    for _, node := range scoring.Nodes {
        if node.ID == scoring.ID {
            return node.run(executions)
        }
    }
    err := errors.New("scoring node not found")
    return nil, apierrors.NewInternalServerApiError(err.Error(), err)
}

func (node *goscenNode) run(executions map[*goscenNode]bool) ([]interface{}, apierrors.ApiError) {
    if executions[node] == true {
        return nil, nil
    }
    for _, dependencyNode := range node.DependenciesNodes {
        if executions[dependencyNode] == false {
            if _, apiErr := dependencyNode.run(executions); apiErr != nil {
                return nil, apiErr
            }
        }
    }
    inputs := make([]interface{}, 0)
    for _, dependencyNode := range node.DependenciesNodes {
        inputs = append(inputs, dependencyNode.ExecutionResult)
    }
    res, apiErr := node.Execution(inputs...)
    if apiErr != nil {
        return nil, apiErr
    }
    node.ExecutionResult = res
    executions[node] = true
    return res, nil
}
