package main

import (
	"cms/model"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (a *App) initStateRoutes() {
	a.Router.HandleFunc("/api/v1/states", a.GetStates).Methods("GET")
	a.Router.HandleFunc("/api/v1/states/{id}", a.GetState).Methods("GET")
	a.Router.HandleFunc("/api/v1/states", a.CreateStates).Methods("POST")
	a.Router.HandleFunc("/api/v1/states/{id}", a.UpdateStates).Methods("PUT")
	a.Router.HandleFunc("/api/v1/states/{id}", a.CreateStates).Methods("DELETE")
}

func (a *App) GetStates(res http.ResponseWriter, req *http.Request) {
	var response model.Response
	states := model.GetAllStates(a.DB)
	res.Header().Set("Content-Type", "application/json")
	response.Status = 200
	response.Error = false
	response.Data = states
	json.NewEncoder(res).Encode(response)
}
func (a *App) GetState(res http.ResponseWriter, req *http.Request) {
	var response model.Response
	vars := mux.Vars(req)
	id, _ := vars["id"]
	idp, err := strconv.Atoi(id)
	res.Header().Set("Content-Type", "application/json")
	if err != nil {
		response.Error = true
		response.Data = err.Error()
	} else {
		state := model.Country{ID: idp}
		err = state.Get(a.DB)
		response.Data = state
		response.Error = false
		if err != nil {
			response.Error = true
			response.Data = err.Error()
		}
	}
	json.NewEncoder(res).Encode(response)
}
func (a *App) CreateStates(res http.ResponseWriter, req *http.Request) {
	var response model.Response
	var state model.State
	res.Header().Set("Content-Type", "Application/json")
	err := json.NewDecoder(req.Body).Decode(&state)
	if err != nil {
		response.Status = http.StatusUnprocessableEntity
		response.Error = true
		response.Data = err.Error()
		json.NewEncoder(res).Encode(response)
	}

	err = state.Create(a.DB)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Error = true
		response.Data = err.Error()
	} else {
		response.Status = http.StatusOK
		response.Error = false
		response.Data = state
	}
	json.NewEncoder(res).Encode(response)

}
func (a *App) UpdateStates(res http.ResponseWriter, req *http.Request) {
    var response model.Response
	vars := mux.Vars(req)
	id := vars["id"]
	stId, err := strconv.Atoi(id)
	var updatedFields map[string]interface{}
	res.Header().Set("Content-Type", "Application/json")
	err = json.NewDecoder(req.Body).Decode(&updatedFields)
	if err != nil {
		response.Status = http.StatusUnprocessableEntity
		response.Error = true
		response.Data = err.Error()
		json.NewEncoder(res).Encode(response)
	}

	state := model.State{ID: stId}
	err = state.Update(a.DB, updatedFields)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Error = true
		response.Data = err.Error()
	} else {
		response.Status = http.StatusOK
		response.Error = true
		response.Data = state
	}
	json.NewEncoder(res).Encode(response)
}
func (a *App) DeleteStates(res http.ResponseWriter, req *http.Request) {
	var response model.Response
	vars := mux.Vars(req)
	id := vars["id"]
	stId, err := strconv.Atoi(id)
	res.Header().Set("Content-Type", "Application/json")
	if err != nil {
		response.Status = http.StatusUnprocessableEntity
		response.Error = true
		response.Data = err.Error()
		json.NewEncoder(res).Encode(response)
	}

	state := model.State{ID: stId}
	err = state.Delete(a.DB)
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
