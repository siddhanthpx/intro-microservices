package handlers

import (
	"encoding/json"
	"intro-microservices/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	nenc := json.NewEncoder(rw)
	err := nenc.Encode(lp)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}
