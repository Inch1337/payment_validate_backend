package main

import (
	"encoding/json"
	"log"
	"net/http"
	"payment_backend/internal/factory"
)

type PayRequest struct {
	Method string `json:"method"`
	Amount int    `json:"amount"`
}

type PayResponse struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

func payHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req PayRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	payment, err := factory.NewPayment(req.Method)
	if err != nil {
		json.NewEncoder(w).Encode(PayResponse{Status: "failed", Error: err.Error()})
		return
	}

	if err := payment.Pay(req.Amount); err != nil {
		json.NewEncoder(w).Encode(PayResponse{Status: "failed", Error: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(PayResponse{Status: "ok"})
}

func main() {
	http.HandleFunc("/payments", payHandler)
	log.Println("API started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
