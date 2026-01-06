package web

import (
	"encoding/json"
	"net/http"

	createaccount "github.com/leomoritz/fullcycle-arquitetura-microsservicos/internal/usecase/create-account"
)

type WebAccountHandler struct {
	CreateAccountUseCase createaccount.CreateAccountUseCase
}

func NewWebAccountHandler(createAccountUseCase createaccount.CreateAccountUseCase) *WebAccountHandler {
	return &WebAccountHandler{
		CreateAccountUseCase: createAccountUseCase,
	}
}

func (h *WebAccountHandler) CreateAccount(response http.ResponseWriter, request *http.Request) {
	var dto createaccount.CreateAccountInputDTO
	err := json.NewDecoder(request.Body).Decode(&dto)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.CreateAccountUseCase.Execute(dto)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(response).Encode(output)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusCreated)
}
