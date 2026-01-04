package web

import (
	"encoding/json"
	"net/http"

	createclient "github.com/leomoritz/fullcycle-arquitetura-microsservicos/internal/usecase/create-client"
)

type WebClientHandler struct {
	createClientUseCase createclient.CreateClientUseCase
}

func NewWebClientHandler(createClientUseCase createclient.CreateClientUseCase) *WebClientHandler {
	return &WebClientHandler{
		createClientUseCase: createClientUseCase,
	}
}

func (h *WebClientHandler) CreateClient(response http.ResponseWriter, request *http.Request) {
	var dto createclient.CreateClientInputDTO
	err := json.NewDecoder(request.Body).Decode(&dto)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.createClientUseCase.Execute(dto)
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
