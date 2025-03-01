package domain

type Indicator struct {
	Name  string
	Label string
	Mark  float64
	Value float64
	Operator string
	Good  bool
}