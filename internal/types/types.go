package types

// InterfaceWeighted interface to represent a weighted object
type InterfaceWeighted interface {
	GetProbability() float64
}

// WeightedElement represents a probability based weighted object
type WeightedElement struct {
	Probability float64 `json:"probability"`
}

// GetProbability gets match probability
func (we WeightedElement) GetProbability() float64 {
	return we.Probability
}
