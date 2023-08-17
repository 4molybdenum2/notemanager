package main

import (
	"log"
	"net/http"
	"time"

	"github.com/4molybdenum2/notemanager/noteservice"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	log.Println("Server Running...")
	// Note Controller
	ns := noteservice.NewNoteController()

	noteRouter := router.PathPrefix("/note").Subrouter()
	noteRouter.HandleFunc("/getAll", ns.GetAllNote).Methods("GET")
	noteRouter.HandleFunc("/get", ns.GetNote).Methods("GET")
	noteRouter.HandleFunc("/post", ns.PostNote).Methods("POST")
	noteRouter.HandleFunc("/update", ns.UpdateNote).Methods("PUT")
	noteRouter.HandleFunc("/delete", ns.DeleteNote).Methods("DELETE")

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		//Timeout
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
