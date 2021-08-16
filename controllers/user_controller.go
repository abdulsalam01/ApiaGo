package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/abdulsalam01/go_gorm/models"
	responses "github.com/abdulsalam01/go_gorm/utils"
	"github.com/gorilla/mux"
)

func (server *Server) GetAll(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	res, err := user.FindAll(server.Db)

	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
	}

	// get all user
	data := models.Response{
		Status:  http.StatusOK,
		Message: "Success!",
		Data:    res,
	}

	responses.JsonResponse(w, http.StatusOK, data)
}

func (server *Server) GetById(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
	}

	// get user by id
	res, _ := user.FindById(*server.Db, int(id))
	data := models.Response{
		Status:  http.StatusOK,
		Message: "Success!",
		Data:    res,
	}

	responses.JsonResponse(w, http.StatusOK, data)
}

func (server *Server) Create(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
	}

	json.Unmarshal(body, &user)

	// create user
	res, _ := user.Create(server.Db)
	data := models.Response{
		Status:  http.StatusOK,
		Message: "Success!",
		Data:    res,
	}

	responses.JsonResponse(w, http.StatusOK, data)
}

func (server *Server) Update(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 32)
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
	}

	json.Unmarshal(body, &user)
	// update user
	res, _ := user.Update(server.Db, int(id))
	data := models.Response{
		Status:  http.StatusOK,
		Message: "Success!",
		Data:    res,
	}

	responses.JsonResponse(w, http.StatusOK, data)
}

func (server *Server) Delete(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
	}

	// delete user
	res, _ := user.Delete(server.Db, int(id))
	data := models.Response{
		Status:  http.StatusOK,
		Message: "Success!",
		Data:    res,
	}

	responses.JsonResponse(w, http.StatusOK, data)
}
