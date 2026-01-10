package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/anirudhaxe/go-api-servers/rest/internal/model"
	"github.com/anirudhaxe/go-api-servers/rest/internal/service"
)

// handler is basically the controller, includes business logic

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.service.AddProduct(&product); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {

	log.Println("GET ALL PRODUCT handler hit")
	products := h.service.ListProducts()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(products)

}
