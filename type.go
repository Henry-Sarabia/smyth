package smyth

type TypeList struct {
	ValueFactor float64  `json:"value_factor"`
	Types       []string `json:"types"`
}

type Type struct {
	Name         string
	ValueFactor  float64
	WeightFactor float64
}
