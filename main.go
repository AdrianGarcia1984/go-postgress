package main

import (
	"net/http"

	"github.com/adriangarcia1984/go-postgress/database"
	"github.com/adriangarcia1984/go-postgress/models"
	"github.com/adriangarcia1984/go-postgress/routes"
	"github.com/gorilla/mux"
)



func main() {

	 database.DBConection()
	 database.DB.AutoMigrate(models.Medidor{})

	r:=mux.NewRouter()
	
	//endpoints
	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/medidores", routes.GetMedidoresHandler).Methods("GET")
	r.HandleFunc("/medidores/{id}", routes.GetMedidorHandler).Methods("GET")
	r.HandleFunc("/medidores", routes.PostMedidorHandler).Methods("POST")
	r.HandleFunc("/medidores/{id}", routes.DeleteMedidorHandler).Methods("DELETE")
	r.HandleFunc("/medidores/{id}", routes.UpdateMedidorHandler).Methods("PUT")
	r.HandleFunc("/medidores/active/{is_active}", routes.ActiveMedidor).Methods("GET")
	r.HandleFunc("/medidores/search/{brand}/{serial}", routes.BrandMedidor).Methods("GET")
	

	r.Use(mux.CORSMethodMiddleware(r))

	http.ListenAndServe(":3000", r)

}
