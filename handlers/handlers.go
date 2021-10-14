package handlers

import (
	"github.com/JonathanCR1/twitterGo/middlew"
	"github.com/JonathanCR1/twitterGo/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)
//Controllers has a port setting and the servers starts listen
func Controllers()  {
	router := mux.NewRouter()
	router.HandleFunc("/login", middlew.CheckBD(routers.SingUp)).Methods("POST")


	PORT := os.Getenv("PORT")
	if PORT == ""{
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT,handler))
}
