package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", &homeHandler{})
	mux.Handle("/recipes", &recipesHandler{})
	mux.Handle("/recipes/", &recipesHandler{})

	http.ListenAndServe(":8080", mux)
}

type homeHandler struct {
}

type recipesHandler struct {
}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the home page"))
}

func (c *recipesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the recipe page"))
}

func (c *recipesHandler) CreateRecipe(w http.ResponseWriter, r *http.Request) {}

func (c *recipesHandler) ListRecipes(w http.ResponseWriter, r *http.Request) {}

func (c *recipesHandler) GetRecipe(w http.ResponseWriter, r *http.Request) {}

func (c *recipesHandler) UpdateRecipe(w http.ResponseWriter, r *http.Request) {}

func (c *recipesHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {}
