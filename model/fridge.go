package model

type Fridge struct {
	ID         int
	Name       string
	Ingredient []Ingredient
}

func NewFridge(name string, id int) *Fridge {
	fridge := Fridge{ID: id, Name: name}

	return &fridge
}
