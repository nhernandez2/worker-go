package models

type ValidateProductsRequest struct {
	ProductIDs []string `json:"product_ids"`
}

type ValidateProductsResponse struct {
	IsValid       bool     `json:"is_valid"`
	NotFoundItems []string `json:"not_found_items"`
}
