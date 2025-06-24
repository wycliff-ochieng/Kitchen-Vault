package internal

import (
	"log"
	//"github.com/wycliff-ochieng/internal"
)

type KitchenService struct {
	l  *log.Logger
	db DBInterface
}

func NewKitcheService(l *log.Logger, db DBInterface) *KitchenService {
	return &KitchenService{
		l:  l,
		db: db,
	}
}

func (k *KitchenService) AddToVault(id string, name string, category string, instructions string) (*Recipe, error) {

	query := `INSERT INTO recipe(id,name,category,instructions,createdat) VALUES($1,$2,$3,$4,$5)`

	//create new recipe
	recipe, err := NewRecipe(id, name, category, instructions)
	if err != nil {
		return nil, err
	}

	_, err = k.db.Exec(query, recipe.Id, recipe.Name, recipe.Category, recipe.Instructions, recipe.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &Recipe{
		Id:           recipe.Id,
		Name:         recipe.Name,
		Category:     recipe.Category,
		Instructions: recipe.Instructions,
	}, nil
}
