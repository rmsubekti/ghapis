package controllers

import (
	"encoding/json"
	"github.com/rmsubekti/ghapis/app"
	"github.com/rmsubekti/ghapis/app/models"
	"github.com/rmsubekti/ghapis/app/utils"
	"net/http"
)

// SignUp create new user
var SignUp = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		utils.Respond(w, utils.Message(false, "invalid request"))
		return
	}

	resp := account.Create()
	utils.Respond(w, resp)
}

// Login user
var Login = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}
	resp := models.Login(account.Email, account.Password)
	utils.Respond(w, resp)
}

// Me logged in user
var Me = func(w http.ResponseWriter, r *http.Request) {
	resp := models.GetUser(app.GetUserID(r))
	utils.Respond(w, resp)
}
