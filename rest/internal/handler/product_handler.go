package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anirudhaxe/go-api-servers/rest/internal/model"
	"github.com/anirudhaxe/go-api-servers/rest/internal/service"
	"github.com/anirudhaxe/go-api-servers/rest/utils"
)

// handler is basically the controller, includes business logic

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) error {
	var product model.Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {

		return fmt.Errorf("Invalid input")

	}

	if err := h.service.CreateProduct(&product); err != nil {
		return fmt.Errorf("Error while creating product: %s", err.Error())
	}

	w.WriteHeader(http.StatusCreated)

	return nil

}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) error {

	products, err := h.service.ListProducts()

	if err != nil {
		return fmt.Errorf("Error while getting products: %s", err.Error())
	}

	utils.WriteJSON(w, http.StatusOK, products)

	return nil

}
