package internal

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type RecipeStore struct {
	db *sql.DB
}

type DBInterface interface {
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
	Close() error
}

func NewRecipeStore() (*RecipeStore, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("cannnot connect:%v", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("cannot ping,%v", err)
	}
	return &RecipeStore{db}, nil
}

func (r *RecipeStore) Init() error {
	return r.CreateRecipeTable()
}

func (r *RecipeStore) CreateRecipeTable() error {
	query := `CREATE TABLE IF NOT EXISTS recipe(
	id serial primary key,
	name varchar(50),
	category varchar(50),
	instructions text,
	createdat timestamp
	)`
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (r *RecipeStore) Close() error {
	if r.db != nil {
		return r.db.Close()
	}
	return nil
}

func (r *RecipeStore) Exec(query string, args ...interface{}) (sql.Result, error) {
	return r.db.Exec(query, args...)
}

func (r *RecipeStore) QueryRow(query string, args ...interface{}) *sql.Row {
	return r.db.QueryRow(query, args...)
}
