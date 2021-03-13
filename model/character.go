package model

type Character struct {
	ID 	 int 	`json:"id"`
	Name string `json:"name"`
}

type Characters []Character

func GetCharacter(c *Character) *Character {
	db := Init()
	db.Find(c)
	return c
}