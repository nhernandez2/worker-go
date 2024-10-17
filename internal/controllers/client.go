package controllers

import (
	"encoding/json"
	"net/http"
	"worker-go/internal/models"
	"worker-go/internal/services"

	"github.com/gorilla/mux"
)

type ClientController struct {
	clientService *services.ClientService
}

func NewClientController(clientService *services.ClientService) *ClientController {
	return &ClientController{clientService: clientService}
}

func (c *ClientController) CreateClient(w http.ResponseWriter, r *http.Request) {
	var client models.Client
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := c.clientService.CreateClient(client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *ClientController) GetClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	client, err := c.clientService.GetClient(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(client)
}

func (c *ClientController) UpdateClient(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var updatedData models.Client
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := c.clientService.UpdateClient(id, updatedData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (c *ClientController) DeleteClient(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := c.clientService.DeleteClient(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
