package controllers

import (
	"encoding/json"
	"net/http"
	"worker-go/internal/models"
	"worker-go/internal/services"

	"github.com/gorilla/mux"
)

type ProductController struct {
	productService *services.ProductService
}

func NewProductController(productService *services.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (c *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := c.productService.CreateProduct(product); err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *ProductController) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	product, err := c.productService.GetProduct(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func (c *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := c.productService.UpdateProduct(id, product); err != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := c.productService.DeleteProduct(id); err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *ProductController) ValidateProducts(w http.ResponseWriter, r *http.Request) {
	var request models.ValidateProductsRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	isValid, notFoundProducts := c.productService.ValidateProducts(request.ProductIDs)

	response := models.ValidateProductsResponse{
		IsValid:       isValid,
		NotFoundItems: notFoundProducts,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
