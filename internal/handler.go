package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type RecipeHandler struct {
	l *log.Logger
	k *KitchenService
}

type VaultReq struct {
	Id           string `json:"idMeal"`
	Name         string `json:"strMeal"`
	Category     string `json:"strCategory"`
	Instructions string `json:"strInstructions"`
}

func NewRecipeHandler(l *log.Logger, k *KitchenService) *RecipeHandler {
	return &RecipeHandler{
		l: l,
		k: k,
	}
}

var (
	ErrInvalidInput = errors.New("some fields are missing")
	ErrNoContent    = errors.New("no recipe for this meal")
)

func (n *RecipeHandler) SearchByMeal(w http.ResponseWriter, r *http.Request) {

	n.l.Println("search recipe my specific meal or food")

	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "missing search query", http.StatusBadRequest)
		return
	}
	//mealURL := "www.themealdb.com/api/json/v1/1/search.php?s=Arrabiata"

	mealURL := fmt.Sprintf("https://www.themealdb.com/api/json/v1/1/search.php?s=%s", url.QueryEscape(query))

	resp, err := http.Get(mealURL)
	if err != nil {
		http.Error(w, "something happened", http.StatusExpectationFailed)
		n.l.Printf("Error:%v", err)
		return
	}
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	defer resp.Body.Close()
	//fmt.Println(string(respData))

	var respObject MealResponse

	err = json.Unmarshal(respData, &respObject)
	if err != nil {
		http.Error(w, "failed to unmarshal", http.StatusExpectationFailed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respObject.Meals)

}

var vault []Recipe

func (n *RecipeHandler) AddRecipeToValt(w http.ResponseWriter, r *http.Request) {
	var recipe Recipe

	err := json.NewDecoder(r.Body).Decode(&recipe)
	if err != nil {
		http.Error(w, "error marshlling the response", http.StatusBadRequest)
		return
	}

	vault = append(vault, recipe)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "added successfully",
	})
}

func (n *RecipeHandler) AddRecipeTODB(w http.ResponseWriter, r *http.Request) {
	var req VaultReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "trouble unmarshlling data", http.StatusExpectationFailed)
		n.l.Printf("error:%v", err)
		return
	}

	if req.Id <= "" || req.Name == "" {
		http.Error(w, "import credentials ,missing", http.StatusExpectationFailed)
		n.l.Printf("Something is off: %v", ErrInvalidInput)
		return
	}

	if req.Instructions == "" {
		http.Error(w, "empty instruction list", http.StatusBadRequest)
		n.l.Printf("no cooking instructions: %v", ErrNoContent)
		return
	}
	//create recipe

	recipe, err := n.k.AddToVault((req.Id), req.Name, req.Category, req.Instructions)
	if err != nil {
		http.Error(w, "faield to add to vault", http.StatusFailedDependency)
		n.l.Printf("the issue:%v", err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&recipe)
}
