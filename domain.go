package goscen

type scoringConfig struct {
    ID      string         `json:"id"`
    Type    string         `json:"type"`
    Loaders []loaderConfig `json:"loaders"`
}

type loaderConfig struct {
    ID           string   `json:"id"`
    Type         string   `json:"type"`
    Dependencies []string `json:"dependencies"`
}
