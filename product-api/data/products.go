package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Desciption string  `json:"description"`
	Price      float32 `json:"price"`
	SKU        string  `json:"sku"`
	CreatedOn  string  `json:"-"`
	UpdatedOn  string  `json:"-"`
	DeletedOn  string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productlist
}

var productlist = []*Product{
	&Product{
		ID:         1,
		Name:       "Latte",
		Desciption: "Frothy milky coffee",
		Price:      2.45,
		SKU:        "abc123",
		CreatedOn:  time.Now().UTC().String(),
		UpdatedOn:  time.Now().UTC().String(),
	},
	&Product{
		ID:         2,
		Name:       "Espresso",
		Desciption: "Short and Strong Coffee without milk",
		Price:      1.99,
		SKU:        "fjd123",
		CreatedOn:  time.Now().UTC().String(),
		UpdatedOn:  time.Now().UTC().String(),
	},
}
