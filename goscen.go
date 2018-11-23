package goscen

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
}

const (
    defaultConfigFile      = "./config.json"
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

func init() {
    scoring := read()
    scoring.check()
    scoring.load()
}
