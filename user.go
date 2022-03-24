package main

import (
	"cms/model"
	"encoding/json"
	"net/http"
	"strconv"
 "github.com/gorilla/mux"
	
)

func (a *App) initUserRoutes() {
	a.Router.HandleFunc("/api/v1/users", a.GetUsers).Methods("GET")
	a.Router.HandleFunc("/api/v1/users/{id}", a.GetUser).Methods("GET")
	a.Router.HandleFunc("/api/v1/users", a.CreateUsers).Methods("POST")
	a.Router.HandleFunc("/api/v1/users/{id}", a.UpdateUsers).Methods("PUT")
	a.Router.HandleFunc("/api/v1/users/{id}", a.DeleteUsers).Methods("DELETE")

}
func (a *App) GetUsers(res http.ResponseWriter, req *http.Request) {
	var response model.Response
	users := model.GetAllUsers(a.DB)
	res.Header().Set("Content-Type", "application/json")
	response.Status = 200
	response.Error = false
	response.Data = users
	json.NewEncoder(res).Encode(response)
}
func (a *App) GetUser(res http.ResponseWriter, req *http.Request) {
    var response model.Response
	vars := mux.Vars(req)
	id,_ :=vars["id"]
	idp,err := strconv.Atoi(id)
	res.Header().Set("Content-Type", "application/json")
	if err != nil {
		response.Error = true
		response.Data = err.Error()
	} else {
		user := model.User{ID: idp}
		err = user.Get(a.DB)
		response.Data = user
		response.Error = false
		if err != nil {
			response.Error = true
			response.Data = err.Error()
		}
	}
	json.NewEncoder(res).Encode(response)

}
func (a *App) CreateUsers(res http.ResponseWriter, req *http.Request) {
	var user model.User
	err := json.NewDecoder(req.Body).Decode(&user)
	var response model.Response
	res.Header().Set("Content-Type", "application/json")
	if err != nil {
		response.Error = true
		response.Data = err.Error()
		json.NewEncoder(res).Encode(response)
	}
	err = user.Create(a.DB)
	if err != nil {
		response.Error = true
		response.Data = err.Error()

	} else {
		response.Error = false
		response.Data = user
	}
	json.NewEncoder(res).Encode(response)
}
func (a *App) UpdateUsers(res http.ResponseWriter, req *http.Request) {
	var response model.Response
	vars := mux.Vars(req)
	id := vars["id"]
	userId, err := strconv.Atoi(id)
	var updatedFields map[string]interface{}
	res.Header().Set("Content-Type", "Application/json")
	err = json.NewDecoder(req.Body).Decode(&updatedFields)
	if err != nil {
		response.Status = http.StatusUnprocessableEntity
		response.Error = true
		response.Data = err.Error()
		json.NewEncoder(res).Encode(response)
	}

	user := model.User{ID: userId}
	err = user.Update(a.DB, updatedFields)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Error = true
		response.Data = err.Error()
	} else {
		response.Status = http.StatusOK
		response.Error = true
		response.Data = user
	}
	json.NewEncoder(res).Encode(response)
}
func (a *App) DeleteUsers(res http.ResponseWriter, req *http.Request) {
	var response model.Response
	vars := mux.Vars(req)
	id := vars["id"]
	userId, err := strconv.Atoi(id)
	res.Header().Set("Content-Type", "Application/json")
	if err != nil {
		response.Status = http.StatusUnprocessableEntity
		response.Error = true
		response.Data = err.Error()
		json.NewEncoder(res).Encode(response)
	}

	user := model.User{ID: userId}
	err = user.Delete(a.DB)
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
