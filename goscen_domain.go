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
