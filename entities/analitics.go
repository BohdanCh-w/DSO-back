package entities

type ResultAnalitics struct {
	Points []WawePoint `yaml:"points"`
	CoefsA []float64   `yaml:"coefs_a"`
	CoefsB []float64   `yaml:"coefs_b"`
}

type SaveResult struct {
	Diff      float64         `json:"diff"`
	Analitics ResultAnalitics `json:"analitics"`
}
