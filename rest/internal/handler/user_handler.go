package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anirudhaxe/go-api-servers/rest/internal/model"
	"github.com/anirudhaxe/go-api-servers/rest/internal/repository"
	"github.com/anirudhaxe/go-api-servers/rest/internal/service"
	"github.com/anirudhaxe/go-api-servers/rest/utils"
)

// handler is basically the controller, includes business logic

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) error {

	val, ok := r.Context().Value("usrEmail").(string)

	if !ok {
		return fmt.Errorf("Error while authorizing user")
	}

	reqUsr, err := h.service.GetUser(val)

	if err != nil {
		return fmt.Errorf("Error while authorizing user")
	}

	if reqUsr.Role != repository.UserRoleAdmin {
		return fmt.Errorf("Error while authorizing user: only admins can register new users.")
	}

	var request model.RegisterUserRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {

		return fmt.Errorf("Invalid input")

	}

	usr, err := h.service.RegisterUser(&request)

	if err != nil {
		return fmt.Errorf("Error while registering user: %s", err.Error())
	}

	tokenStr, err := usr.GenerateJwtToken()

	if err != nil {
		return fmt.Errorf("Error while generating token %s", err.Error())
	}

	resp := &model.UserResponse{
		Username: usr.Username,
		Email:    usr.Email,
		Role:     usr.Role,
		Token:    tokenStr,
	}

	utils.WriteJSON(w, http.StatusCreated, resp)

	return nil

}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) error {
	var request model.LoginUserRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return fmt.Errorf("Invalid input")
	}

	usr, err := h.service.GetUser(request.Email)

	if err != nil {
		return fmt.Errorf("Error while logging in: %s", err.Error())
	}

	if isPassCorr := usr.ValidatePassword(request.Password); isPassCorr == false {

		return fmt.Errorf("Incorrect password")
	}

	tokenStr, err := usr.GenerateJwtToken()

	if err != nil {
		return fmt.Errorf("Error while generating token %s", err.Error())
	}

	resp := &model.UserResponse{
		Username: usr.Username,
		Email:    usr.Email,
		Role:     usr.Role,
		Token:    tokenStr,
	}

	utils.WriteJSON(w, http.StatusOK, resp)

	return nil

}

// func (h *UserHandler) GetProducts(w http.ResponseWriter, r *http.Request) error {
//
// 	products, err := h.service.ListProducts()
//
// 	if err != nil {
// 		return fmt.Errorf("Error while getting products: %s", err.Error())
// 	}
//
// 	utils.WriteJSON(w, http.StatusOK, products)
//
// 	return nil
//
// }
