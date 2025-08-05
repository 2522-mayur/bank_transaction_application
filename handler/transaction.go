package handler

import (
	"BANK_TRANSACTION_APPLICATION/model"
	"BANK_TRANSACTION_APPLICATION/services"
	"encoding/json"
	"net/http"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var tx model.Transaction
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"error":"Invalid input"}`, http.StatusBadRequest)
		return
	}

	if err := services.ProcessTransaction(tx); err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"error":"Transaction failed: `+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Transaction successful",
	})
}

