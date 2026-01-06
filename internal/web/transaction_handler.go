package web

import (
	"encoding/json"
	"net/http"

	createTransaction "github.com/leomoritz/fullcycle-arquitetura-microsservicos/internal/usecase/create-transaction"
)

type WebTransactionHandler struct {
	CreateTransactionUseCase createTransaction.CreateTransactionUseCase
}

func NewWebTransactionHandler(createTransactionUseCase createTransaction.CreateTransactionUseCase) *WebTransactionHandler {
	return &WebTransactionHandler{
		CreateTransactionUseCase: createTransactionUseCase,
	}
}

func (h *WebTransactionHandler) CreateTransaction(response http.ResponseWriter, request *http.Request) {
	var dto createTransaction.CreateTransactionInputDTO
	err := json.NewDecoder(request.Body).Decode(&dto)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := request.Context()
	output, err := h.CreateTransactionUseCase.Execute(ctx, dto)
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
