package main

import (
	"log"
	"net/http"
	"urbanAPI/controller"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	corsOpt := cors.New(
		cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
				http.MethodOptions,
				http.MethodHead,
			},
			AllowedHeaders: []string{"*"},
		},
	)

	r := mux.NewRouter()
	prefix := "/api"

	// CUSTOMER
	r.HandleFunc(prefix+"/customers", controller.GetCustomerByID).Methods("GET")

	// PRODUCT
	r.HandleFunc(prefix+"/products", controller.GetProductByID).Methods("GET")
	r.HandleFunc(prefix+"/images", controller.GetImagesbyProductID).Methods("GET")
	r.HandleFunc(prefix+"/orders", controller.GetProductsByOrderId).Methods("GET")

	// COMMENTS
	r.HandleFunc(prefix+"/comments", controller.GetAllComments).Methods("GET")

	http.Handle("/", r)

	log.Println("Servidor rodando na porta :3000")
	http.ListenAndServe(":3000", corsOpt.Handler(r))
}
