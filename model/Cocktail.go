package model

type Cocktail struct {
	CocktailID   int            `json:"id,omitempty"`
	Name         string         `json:"name"`
	Instructions []*Instruction `json:"instructions"`
}
