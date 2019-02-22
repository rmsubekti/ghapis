package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rmsubekti/ghapis/app"
	"github.com/rmsubekti/ghapis/app/models"
	"github.com/rmsubekti/ghapis/app/utils"

	"github.com/gorilla/mux"
)

// CreateProduct Controller
var CreateProduct = func(w http.ResponseWriter, r *http.Request) {
	if !app.RoleAccess(r, []models.RoleType{models.RoleAdmin, models.RolePM}) {
		utils.Respond(w, utils.Message(false, "Access restricted"))
		return
	}
	product := &models.Product{}

	err := json.NewDecoder(r.Body).Decode(product)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding payload"))
		return
	}

	resp := product.Create()
	utils.Respond(w, resp)
}

// GetProduct Controller
var GetProduct = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if err := models.GetProduct(vars["id"]); err != nil {
		resp := utils.Message(true, err.Name)
		resp["product"] = err
		utils.Respond(w, resp)
		return
	}
	utils.Respond(w, utils.Message(false, "Not found : Product is not available in this server"))
}

//GetProducts controller
var GetProducts = func(w http.ResponseWriter, r *http.Request) {
	if err := models.GetProducts(); err != nil {
		resp := utils.Message(true, "All product we have")
		resp["products"] = err
		utils.Respond(w, resp)
		return
	}
	utils.Respond(w, utils.Message(false, "Not found : Currently we dont have product to sale"))
}
