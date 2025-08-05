package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"BANK_TRANSACTION_APPLICATION/model"
	"BANK_TRANSACTION_APPLICATION/services"

	"github.com/gorilla/mux"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var acc model.Account
	if err := json.NewDecoder(r.Body).Decode(&acc); err != nil {
		http.Error(w, `{"error":"Invalid request body"}`, http.StatusBadRequest)
		return
	}

	if err := services.InsertAccount(acc); err != nil {
		http.Error(w, `{"error":"Failed to create account"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Account created"})
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, `{"error":"Invalid account ID"}`, http.StatusBadRequest)
		return
	}

	acc, err := services.FetchAccount(id)
	if err != nil {
		http.Error(w, `{"error":"Account not found"}`, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(acc)
}
