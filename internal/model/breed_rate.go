package model

type BreedRate struct {
	ID         int     `json:"id" db:"id"`
	Species    Species `json:"species" db:"species"`
	Breed      string  `json:"breed" db:"breed"`
	SizeClass  SizeClass `json:"size_class" db:"size_class"`
	RiskFactor float64 `json:"risk_factor" db:"risk_factor"`
}
