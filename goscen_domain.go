package goscen

import "github.com/mercadolibre/go-meli-toolkit/goutils/apierrors"

const (
    scoringTypeComplete    = "complete"
    scoringTypeProgressive = "progressive"
    loaderTypeAPI          = "API"
)

var (
    scoringTypes = []string{
        scoringTypeComplete,
        scoringTypeProgressive,
    }

    loaderTypes = []string{
        loaderTypeAPI,
    }
)

type goscenScoring struct {
    ID      string          `json:"id"`
    Type    string          `json:"type"`
    Loaders []*goscenLoader `json:"loaders"`
}

type goscenLoader struct {
    ID                  string          `json:"id"`
    Type                string          `json:"type"`
    DependenciesID      []string        `json:"dependencies"`
    DependenciesLoaders []*goscenLoader `json:"-"`
    Execution           goscenExecution `json:"-"`
    ExecutionResult     []interface{}   `json:"-"`
}

type goscenExecution func(...interface{}) ([]interface{}, apierrors.ApiError)
