package api

import (
	"errors"
	"net/http"

	gson "github.com/DiscoFighter47/gSON"
	"github.com/DiscoFighter47/todo/backend/model"
)

type authSignUpBody struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (body *authSignUpBody) validate() gson.ValidationError {
	errV := gson.ValidationError{}
	if body.ID == "" {
		errV.Add("id", "required")
	}
	if body.Name == "" {
		errV.Add("name", "required")
	}
	if body.Password == "" {
		errV.Add("password", "required")
	}
	if len(errV) > 0 {
		return errV
	}
	return nil
}

func (api *API) authSignUp(w http.ResponseWriter, r *http.Request) {
	body := &authSignUpBody{}
	if err := gson.ParseBody(r, body); err != nil {
		panic(gson.NewAPIerror("Unable To Parse Body", http.StatusUnprocessableEntity, err))
	}
	if err := body.validate(); err != nil {
		panic(gson.NewAPIerror("Invalid Request Body", http.StatusBadRequest, err))
	}

	if err := api.store.AddUser(&model.User{
		ID:       body.ID,
		Name:     body.Name,
		Password: body.Password,
	}); err != nil {
		panic(gson.NewAPIerror("Unable To Add User", http.StatusInternalServerError, err, body.ID))
	}

	gson.ServeData(w, gson.Object{
		"id":   body.ID,
		"name": body.Name,
	})
}

type authSignInBody struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

func (body *authSignInBody) validate() gson.ValidationError {
	errV := gson.ValidationError{}
	if body.ID == "" {
		errV.Add("id", "required")
	}
	if body.Password == "" {
		errV.Add("password", "required")
	}
	if len(errV) > 0 {
		return errV
	}
	return nil
}

func (api *API) authSignIn(w http.ResponseWriter, r *http.Request) {
	body := &authSignInBody{}
	if err := gson.ParseBody(r, body); err != nil {
		panic(gson.NewAPIerror("Unable To Parse Body", http.StatusUnprocessableEntity, err))
	}
	if err := body.validate(); err != nil {
		panic(gson.NewAPIerror("Invalid Request Body", http.StatusBadRequest, err))
	}

	usr, err := api.store.GetUser(body.ID)
	if err != nil {
		panic(gson.NewAPIerror("Unable To Find User", http.StatusBadRequest, err, body.ID))
	}
	if usr.Password != body.Password {
		panic(gson.NewAPIerror("Incorrect Password", http.StatusBadRequest, errors.New("password dosen't match"), body.ID))
	}

	token, err := api.auth.GenerateToken(body.ID)
	if err != nil {
		panic(gson.NewAPIerror("Unable To Generate Token", http.StatusInternalServerError, err))
	}

	gson.ServeData(w, gson.Object{
		"id":    body.ID,
		"token": token,
	})
}

func (api *API) authCheck(w http.ResponseWriter, r *http.Request) {
	gson.ServeData(w, gson.Object{
		"message": "Hello Secret Universe! Welcome " + r.Header.Get("subject"),
	})
}
