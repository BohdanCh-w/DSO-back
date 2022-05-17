package entities

type ResultAnalitics struct {
	Points []WavePoint `yaml:"points"`
	CoefsA []float64   `yaml:"coefs_a"`
	CoefsB []float64   `yaml:"coefs_b"`
}

type SaveResult struct {
	DiffMethodA      float64         `json:"diff_method_a,omitempty"`
	DiffMethodB      float64         `json:"diff_method_b,omitempty"`
	AnaliticsMethodA ResultAnalitics `json:"analitics_method_a,omitempty"`
	AnaliticsMethodB ResultAnalitics `json:"analitics_method_b,omitempty"`
}
