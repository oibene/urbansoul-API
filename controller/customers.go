package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"urbanAPI/database"
	"urbanAPI/models"
)

type Customer struct {
	Customer_id int     `db:"customer_id" json:",omitempty"`
	Name        string  `db:"name" json:",omitempty"`
	Last_name   *string `db:"last_name" json:",omitempty"`
	Email       string  `db:"email" json:",omitempty"`
	Password    string  `db:"password" json:",omitempty"`
	Img_URL     *string `db:"img_url" json:",omitempty"`
}

func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	var ct []Customer
	data := database.ConnectDB()

	var input models.GetCustomerInput
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())

		return
	}

	err = data.Select(&ct, `SELECT *
							FROM customer 
							WHERE customer_id = $1`, input.Customer_id)

	if err != nil {
		log.Println("Erro ao consultar usuario!", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ct)
}
