package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"urbanAPI/database"
	"urbanAPI/models"
)

type Product struct {
	Product_id     int     `db:"product_id" json:",omitempty"`
	Product_name   string  `db:"product_name" json:",omitempty"`
	Gender         string  `db:"gender" json:",omitempty"`
	Size           string  `db:"size" json:",omitempty"`
	Color          *string `db:"color" json:",omitempty"`
	Price          string  `db:"price" json:",omitempty"`
	Descount_price *string `db:"descount_price" json:",omitempty"`

	Category    *string `db:"category" json:,omitempty`
	Description *string `db:"description" json:,omitempty`
	Notes       *string `db:"notes" json:,omitempty`
	Composition *string `db:"composition" json:,omitempty`
}

type Order_Items struct {
	Order_id       int     `db:"order_items_id" json:",omitempty"`
	Product_id     int     `db:"product_id" json:",omitempty"`
	Product_name   string  `db:"product_name" json:",omitempty"`
	Gender         string  `db:"gender" json:",omitempty"`
	Size           string  `db:"size" json:",omitempty"`
	Color          *string `db:"color" json:",omitempty"`
	Price          string  `db:"price" json:",omitempty"`
	Descount_price *string `db:"descount_price" json:",omitempty"`
}

type Images struct {
	Product_id int    `db:"product_id" json:",omitempty"`
	Img_url    string `db:"img_url" json:",omitempty"`
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	var p []Product
	data := database.ConnectDB()

	var input models.GetProductInput
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())

		return
	}

	err = data.Select(&p, `SELECT p.product_name, p.gender, p.size, p.color, p.price, p.descount_price,
							c.description as category,
							d.description, d.notes, d.composition

							FROM product p 

							INNER JOIN category c
							ON c.category_code = p.category_code

							INNER JOIN description d
							ON d.model_code = p.model_code

							WHERE p.product_id = $1`, input.Product_id)

	if err != nil {
		log.Println("Erro ao consultar produto!", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func GetProductsByOrderId(w http.ResponseWriter, r *http.Request) {
	var o []Order_Items
	data := database.ConnectDB()

	var input models.GetProductInput
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())

		return
	}

	err = data.Select(&o, `SELECT o.order_items_id,
							p.product_id, p.product_name, p.gender, p.size, p.color, p.price, p.descount_price
							FROM order_items o

							inner join product p on o.product_id = p.product_id
							WHERE o.order_items_id = $1`, input.Orders_id)

	if err != nil {
		log.Println("Erro ao consultar lista de produtos!", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(o)
}

func GetImagesbyProductID(w http.ResponseWriter, r *http.Request) {
	var i []Images
	data := database.ConnectDB()

	var input models.GetProductInput
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())

		return
	}

	err = data.Select(&i, `SELECT * FROM images
							WHERE product_id = $1`, input.Product_id)

	if err != nil {
		log.Println("Erro ao consultar images!", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(i)
}
