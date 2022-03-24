package main

import (
	"cms/model"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (a *App) initHospitalRoutes() {
	a.Router.HandleFunc("/api/v1/hospitals", a.GetHospitals).Methods("GET")
	a.Router.HandleFunc("/api/v1/hospitals/{id}", a.GetHospital).Methods("GET")
	a.Router.HandleFunc("/api/v1/hospitals", a.CreateHospitals).Methods("POST")
	a.Router.HandleFunc("/api/v1/hospitals/{id}", a.UpdateHospitals).Methods("PUT")
	a.Router.HandleFunc("/api/v1/hospitals/{id}", a.DeleteHospitals).Methods("DELETE")

}
func (a *App) GetHospitals(res http.ResponseWriter, req *http.Request) {

	var response model.Response
	hospitals := model.GetAllHospitals(a.DB)
	res.Header().Set("Content-Type", "application/json")
	response.Status = 200
	response.Error = false
	response.Data = hospitals
	json.NewEncoder(res).Encode(response)

}
func (a *App) GetHospital(res http.ResponseWriter, req *http.Request) {
	var response model.Response
	vars := mux.Vars(req)
	id, _ := vars["id"]
	idp, err := strconv.Atoi(id)
	res.Header().Set("Content-Type", "application/json")
	if err != nil {
		response.Error = true
		response.Data = err.Error()
	} else {
		hospital := model.Hospital{ID: idp}
		err = hospital.Get(a.DB)
		response.Data = hospital
		response.Error = false
		if err != nil {
			response.Error = true
			response.Data = err.Error()
		}
	}
	json.NewEncoder(res).Encode(response)

}
func (a *App) CreateHospitals(res http.ResponseWriter, req *http.Request) {
	var hospital model.Hospital
	err := json.NewDecoder(req.Body).Decode(&hospital)
	var response model.Response
	res.Header().Set("Content-Type", "application/json")
	if err != nil {
		response.Error = true
		response.Data = err.Error()
		json.NewEncoder(res).Encode(response)
	}
	err = hospital.Create(a.DB)
	if err != nil {
		response.Error = true
		response.Data = err.Error()

	} else {
		response.Error = false
		response.Data = hospital
	}
	json.NewEncoder(res).Encode(response)

}
func (a *App) UpdateHospitals(res http.ResponseWriter, req *http.Request) {
	var response model.Response
	vars := mux.Vars(req)
	id := vars["id"]
	hosId, err := strconv.Atoi(id)
	var updatedFields map[string]interface{}
	res.Header().Set("Content-Type", "Application/json")
	err = json.NewDecoder(req.Body).Decode(&updatedFields)
	if err != nil {
		response.Status = http.StatusUnprocessableEntity
		response.Error = true
		response.Data = err.Error()
		json.NewEncoder(res).Encode(response)
	}

	hospital := model.Hospital{ID: hosId}
	err = hospital.Update(a.DB, updatedFields)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Error = true
		response.Data = err.Error()
	} else {
		response.Status = http.StatusOK
		response.Error = true
		response.Data = hospital
	}
	json.NewEncoder(res).Encode(response)
}
func (a *App) DeleteHospitals(res http.ResponseWriter, req *http.Request) {
	var response model.Response
	vars := mux.Vars(req)
	id := vars["id"]
	hosId, err := strconv.Atoi(id)
	res.Header().Set("Content-Type", "Application/json")
	if err != nil {
		response.Status = http.StatusUnprocessableEntity
		response.Error = true
		response.Data = err.Error()
		json.NewEncoder(res).Encode(response)
	}

	hospital := model.Hospital{ID: hosId}
	err = hospital.Delete(a.DB)
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
