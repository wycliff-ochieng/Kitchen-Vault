package internal

import "time"

type Recipe struct {
	Id           string    `json:"idMeal"`
	Name         string    `json:"strMeal"`
	Category     string    `json:"strCategory"`
	Instructions string    `json:"strInstructions"`
	CreatedAt    time.Time `json:"time"`
}

type Ingredient struct {
	Id       int32  `json:"id"`
	RecipeID int32  `json:"recipeID"`
	Name     string `json:"name"`
}

type Meal struct {
	ID           string `json:"idMeal"`
	Title        string `json:"strMeal"`
	Category     string `json:"strCategory"`
	Area         string `json:"strArea"`
	Instructions string `json:"strInstructions"`
	Thumbnail    string `json:"strMealThumb"`
}

type MealResponse struct {
	Meals []Meal `json:"meals"`
}

func NewRecipe(id string, name string, category string, instructions string) (*Recipe, error) {
	return &Recipe{
		Id:           id,
		Name:         name,
		Category:     category,
		Instructions: instructions,
	}, nil
}
