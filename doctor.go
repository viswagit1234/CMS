package main

import (
	"cms/model"
	"encoding/json"
	
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (a *App) initDoctorRoutes() {
	a.Router.HandleFunc("/api/v1/doctors", a.GetDoctors).Methods("GET")
	a.Router.HandleFunc("/api/v1/doctors/{id}", a.GetDoctor).Methods("GET")
	a.Router.HandleFunc("/api/v1/doctors", a.CreateDoctors).Methods("POST")
	a.Router.HandleFunc("/api/v1/doctors/{id}", a.UpdateDoctors).Methods("PUT")
	a.Router.HandleFunc("/api/v1/doctors/{id}", a.CreateDoctors).Methods("DELETE")
}

func (a *App) GetDoctors(res http.ResponseWriter, req *http.Request) {
	var response model.Response
	doctor := model.GetAllDoctors(a.DB)
	res.Header().Set("Contentent_Type","Application/json")
	response.Status = 200
	response.Error = true
	response.Data = doctor
	json.NewEncoder(res).Encode(response)

}
func (a *App) GetDoctor(res http.ResponseWriter, req *http.Request) {
	var response model.Response
	vars := mux.Vars(req)
	id, _ := vars["id"]
	idp, err := strconv.Atoi(id)
	res.Header().Set("Content-Type", "application/json")
	if err != nil {
		response.Error = true
		response.Data = err.Error()
	} else {
		doctor := model.Doctor{ID: idp}
		err = doctor.Get(a.DB)
		response.Data = doctor
		response.Error = false
		if err != nil {
			response.Error = true
			response.Data = err.Error()
		}
	}
	json.NewEncoder(res).Encode(response)

}
func (a *App) CreateDoctors(res http.ResponseWriter, req *http.Request) {
	var doctor model.Doctor
	err := json.NewDecoder(req.Body).Decode(&doctor)
	var response model.Response
	res.Header().Set("Content_Type","Application/json")
	if err != nil{
		
		response.Error = true
		response.Data = err.Error()
		json.NewEncoder(res).Encode(response)
	}
	err = doctor.Create(a.DB)
	if err != nil {
		response.Error = true
		response.Data = doctor
	}else {
		response.Error = false 
		response.Data = err.Error()
	}
	json.NewEncoder(res).Encode(response)

}
func (a *App) UpdateDoctors(res http.ResponseWriter, req *http.Request) {
	var response model.Response
	vars := mux.Vars(req)
	id := vars["id"]
	docId, err := strconv.Atoi(id)
	var updatedFields map[string]interface{}
	res.Header().Set("Content-Type", "Application/json")
	err = json.NewDecoder(req.Body).Decode(&updatedFields)
	if err != nil {
		response.Status = http.StatusUnprocessableEntity
		response.Error = true
		response.Data = err.Error()
		json.NewEncoder(res).Encode(response)
	}

	doctor := model.Doctor{ID: docId}
	err = doctor.Update(a.DB, updatedFields)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Error = true
		response.Data = err.Error()
	} else {
		response.Status = http.StatusOK
		response.Error = true
		response.Data = doctor
	}
	json.NewEncoder(res).Encode(response)
	    
}
func (a *App) DeleteDoctors(res http.ResponseWriter, req *http.Request) {
      var response model.Response
	  vars := mux.Vars(req)
	  id, _:= vars["id"]
	  docId, err:= strconv.Atoi(id)
	  res.Header().Set("Content_TyPe","Application/json")
	  if err != nil {
		  response.Status = http.StatusUnprocessableEntity
		  response.Error = true
		  response.Data = err.Error()
		  json.NewEncoder(res).Encode(response)
	  }
	  doctor := model.Doctor{ID : docId}
	  err = doctor.Delete(a.DB)
	  if err != nil {
           response.Status = http.StatusInternalServerError
		   response.Error = false
		   response.Data = err.Error()

	  }else{
		  response.Status = http.StatusOK
		  response.Data = "data delete successfully"
		  response.Error  = true
	  }
	   json.NewEncoder(res).Encode(response)

}
