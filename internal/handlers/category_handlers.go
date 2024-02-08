package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ricardoraposo/blogs-api-go/internal/database"
	"github.com/ricardoraposo/blogs-api-go/internal/entity"
	"github.com/ricardoraposo/blogs-api-go/internal/utils"
)

type CategoryHandler struct {
	categoryDB database.CategoryDBInterface
}

type CreateCategoryDTO struct {
	Name string `json:"name"`
}

func NewCategoryHandler(categoryDB database.CategoryDBInterface) *CategoryHandler {
	return &CategoryHandler{categoryDB: categoryDB}
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var p CreateCategoryDTO
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		utils.WriteToJson(w, map[string]string{"error": "invalid request"})
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	category := entity.NewCategory(p.Name)

    newCategory, err := h.categoryDB.CreateCategory(category)
    if err != nil {
        utils.WriteToJson(w, map[string]string{"error": "could not create category"})
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    utils.WriteToJson(w, newCategory)
    w.WriteHeader(http.StatusCreated)
}

func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
    categories, err := h.categoryDB.GetCategories()
    if err != nil {
        utils.WriteToJson(w, map[string]string{"error": "could not get categories"})
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    utils.WriteToJson(w, categories)
    w.WriteHeader(http.StatusOK)
}
