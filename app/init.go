package app

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/thisisommore/go-user-app-backend/db"
	"github.com/thisisommore/go-user-app-backend/user/userhandler"
)

type App struct {
	Router *mux.Router
}

func Init(app *App) {
	app.Router = mux.NewRouter()
	app.Router.HandleFunc("/user", userhandler.AddUser).Methods("POST")
	port := ":" + os.Getenv("PORT")
	db.Initialize()
	log.Fatal(http.ListenAndServe(port, app.Router))
}
