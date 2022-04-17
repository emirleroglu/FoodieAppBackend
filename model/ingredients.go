package model

type Ingredient struct {
	Name   string
	Id     int
	Unit   int
	Amount int
}

func AddIngredientsToFrifgeId(fridgeId int) *Fridge {
	temp := NewFridge("Test", fridgeId)
	var list = []Ingredient{{Name: "Elma",
		Id:     1,
		Unit:   1,
		Amount: 2}}

	temp.Ingredient = list
	return temp

}
