package api

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/wycliff-ochieng/internal"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{addr}
}

func (s *APIServer) Run() {
	l := log.New(os.Stdout, "SIMPLE RECIPE APPLICATION WITH NO BUSINESS VALUE", log.LstdFlags)

	db, err := internal.NewRecipeStore()
	if err != nil {
		log.Fatalf("Error setting pg:%v", err)
	}

	if err = db.Init(); err != nil {
		log.Fatalf("Error Initializing:%v", err)
	}

	u := internal.NewKitcheService(l, db)

	sh := internal.NewRecipeHandler(l, u)

	router := mux.NewRouter()

	searchRouter := router.Methods("GET").Subrouter()
	searchRouter.HandleFunc("/api/search", sh.SearchByMeal)

	postDB := router.Methods("POST").Subrouter()
	postDB.HandleFunc("/api/add", sh.AddRecipeTODB)

	http.ListenAndServe(s.addr, router)
}
