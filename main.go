package main

import (
	"cms/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type App struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (a *App) Init() {
	db, err := config.GetDB()
	if err != nil {
		log.Fatal("Unable to connect database:", err.Error())
	}
	a.DB = db
	a.Router = mux.NewRouter()
	a.initDoctorRoutes()
	a.initStateRoutes()
	a.initCountryRoutes()
	a.initHospitalRoutes()
	a.initUserRoutes()

}

func (a *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, a.Router))
}

func main() {
	app := App{}
	app.Init()
	app.Run(":8000")

}
