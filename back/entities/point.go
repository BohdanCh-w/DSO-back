package entities

type WavePoint struct {
	X float64 `json:"x"`
	Y float64 `json:"y,omitempty"`
	Z float64 `json:"z,omitempty"`
}
