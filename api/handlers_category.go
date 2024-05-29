package api

import (
	"fmt"
	"net/http"

	"evvoli.ru/internal/utils"
)

func (app *ApiApplication) CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	category := map[string]string{
		"id":   "1",
		"name": "Phones",
	}

	err := app.writeJson(w, 201, category, nil)
	if err != nil {
		app.Logger.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (app *ApiApplication) GetAllCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "get all categories")
}

func (app *ApiApplication) GetCategoryHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadInt64IDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of movie %d\n", id)
}
