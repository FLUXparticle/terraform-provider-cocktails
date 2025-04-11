package model

type Ingredient struct {
	IngredientID int    `json:"id,omitempty"`
	Name         string `json:"name"`
}
