package app

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/thisisommore/go-user-app-backend/db"
	"github.com/thisisommore/go-user-app-backend/user/userhandler"
)

var AppRouter *mux.Router

func Init() {
	CreateRouter()
	port := ":" + os.Getenv("PORT")
	db.Initialize()
	log.Fatal(http.ListenAndServe(port, AppRouter))
}

func CreateRouter() *mux.Router {
	AppRouter = mux.NewRouter()
	AppRouter.HandleFunc("/user", userhandler.AddUser).Methods("POST")
	AppRouter.HandleFunc("/user/{id}", userhandler.GetUser).Methods("GET")
	AppRouter.HandleFunc("/user/{id}", userhandler.DeleteUser).Methods("DELETE")
	AppRouter.HandleFunc("/user/{id}", userhandler.UpdateUser).Methods("PUT")
	return AppRouter
}
