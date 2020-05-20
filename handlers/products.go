package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hzhyvinskyi/micro/data"
)

type Products struct {
	l	*log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	productList := data.GetProducts()
	json.NewEncoder(w)
	productJson, err := json.Marshal(productList)
	if err != nil {
		http.Error(w, "Unable to marshall to JSON", http.StatusBadRequest)
		return
	}
	w.Write(productJson)
	w.WriteHeader(http.StatusOK)
}
