package main

import (
	"cms/model"
	"encoding/json"
	"net/http"
	"strconv"
  "github.com/gorilla/mux"
)

func (a *App) initCountryRoutes() {
	a.Router.HandleFunc("/api/v1/countries", a.GetCountries).Methods("GET")
	a.Router.HandleFunc("/api/v1/countries/{id}", a.GetCountry).Methods("GET")
	a.Router.HandleFunc("/api/v1/countries", a.CreateCountries).Methods("POST")
	a.Router.HandleFunc("/api/v1/countries/{id}", a.UpdateCountries).Methods("PUT")
	a.Router.HandleFunc("/api/v1/countries/{id}", a.DeleteCountries).Methods("DELETE")

}
func (a *App) GetCountries(res http.ResponseWriter, req *http.Request) {

	var response model.Response
	countries := model.GetAllCountries(a.DB)
	res.Header().Set("Content-Type", "application/json")
	response.Status = 200
	response.Error = false
	response.Data = countries
	json.NewEncoder(res).Encode(response)

}
func (a *App) GetCountry(res http.ResponseWriter, req *http.Request) {
	var response model.Response
	vars := mux.Vars(req)
	id, _ := vars["id"]
	idp, err := strconv.Atoi(id)
	res.Header().Set("Content-Type", "application/json")
	if err != nil {
		response.Error = true
		response.Data = err.Error()
	} else {
		coutry := model.Country{ID: idp}
		err = coutry.Get(a.DB)
		response.Data = coutry
		response.Error = false
		if err != nil {
			response.Error = true
			response.Data = err.Error()
		}
	}
	json.NewEncoder(res).Encode(response)

}
func (a *App) CreateCountries(res http.ResponseWriter, req *http.Request) {
	var country model.Country
	err := json.NewDecoder(req.Body).Decode(&country)
	var response model.Response
	res.Header().Set("Content-Type", "application/json")
	if err != nil {
		response.Error = true
		response.Data = err.Error()
		json.NewEncoder(res).Encode(response)
	}
	err = country.Create(a.DB)
	if err != nil {
		response.Error = true
		response.Data = err.Error()

	} else {
		response.Error = false
		response.Data = country
	}
	json.NewEncoder(res).Encode(response)

}
func (a *App) UpdateCountries(res http.ResponseWriter, req *http.Request) {
	var response model.Response
	vars := mux.Vars(req)
	id := vars["id"]
	conId, err := strconv.Atoi(id)
	var updatedFields map[string]interface{}
	res.Header().Set("Content-Type", "Application/json")
	err = json.NewDecoder(req.Body).Decode(&updatedFields)
	if err != nil {
		response.Status = http.StatusUnprocessableEntity
		response.Error = true
		response.Data = err.Error()
		json.NewEncoder(res).Encode(response)
	}

	country := model.Country{ID: conId}
	err = country.Update(a.DB, updatedFields)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Error = true
		response.Data = err.Error()
	} else {
		response.Status = http.StatusOK
		response.Error = true
		response.Data = country
	}
	json.NewEncoder(res).Encode(response)
}
func (a *App) DeleteCountries(res http.ResponseWriter, req *http.Request) {
	var response model.Response
	vars := mux.Vars(req)
	id := vars["id"]
	conId, err := strconv.Atoi(id)
	res.Header().Set("Content-Type", "Application/json")
	if err != nil {
		response.Status = http.StatusUnprocessableEntity
		response.Error = true
		response.Data = err.Error()
		json.NewEncoder(res).Encode(response)
	}

	country := model.Country{ID: conId}
	err = country.Delete(a.DB)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Error = true
		response.Data = err.Error()
	} else {
		response.Status = http.StatusOK
		response.Error = true
		response.Data = "Data delete successfully"
	}
	json.NewEncoder(res).Encode(response)

}
