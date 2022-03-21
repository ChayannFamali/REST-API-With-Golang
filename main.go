package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Transaction struct {
	ID      string `json:"id"`
	Pin     string `json:"PIN"`
	Summary int    `json:"Transactions"`
	User    *User  `json:"User"`
}

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var transactions []Transaction

func getTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}
func getTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range transactions {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Transaction{})
}

func createTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var transaction Transaction
	_ = json.NewDecoder(r.Body).Decode(&transaction)
	transaction.ID = strconv.Itoa(rand.Intn(1000000))
	transaction.Pin = strconv.Itoa(rand.Intn(1000000))
	transactions = append(transactions, transaction)
	json.NewEncoder(w).Encode(transaction)
}

func updateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range transactions {
		if item.ID == params["id"] {
			transactions = append(transactions[:index], transactions[index+1:]...)
			var transaction Transaction
			_ = json.NewDecoder(r.Body).Decode(&transaction)
			transaction.ID = params["id"]
			transactions = append(transactions, transaction)
			json.NewEncoder(w).Encode(transaction)
			return
		}
	}
	json.NewEncoder(w).Encode(transactions)
}

func deleteTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range transactions {
		if item.ID == params["id"] {
			transactions = append(transactions[:index], transactions[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(transactions)
}

func main() {
	r := mux.NewRouter()
	transactions = append(transactions, Transaction{ID: "1", Summary: 1000, Pin: "1234", User: &User{Firstname: "Дмитрий", Lastname: "Иванов"}})
	transactions = append(transactions, Transaction{ID: "2", Summary: 1233334, Pin: "3123", User: &User{Firstname: "Валерия", Lastname: "Тухачевская"}})
	transactions = append(transactions, Transaction{ID: "3", Summary: 2000, Pin: "1231", User: &User{Firstname: "Сергей", Lastname: "Козлов"}})
	transactions = append(transactions, Transaction{ID: "4", Summary: 123, Pin: "4123", User: &User{Firstname: "Андрей", Lastname: "Толстой"}})
	transactions = append(transactions, Transaction{ID: "5", Summary: 123432412, Pin: "1333", User: &User{Firstname: "Ирина", Lastname: "Козлова"}})
	transactions = append(transactions, Transaction{ID: "6", Summary: 234234234, Pin: "3123", User: &User{Firstname: "Мурат", Lastname: "Юсупов"}})
	transactions = append(transactions, Transaction{ID: "7", Summary: 323123, Pin: "1235", User: &User{Firstname: "Василий", Lastname: "Третьяков"}})
	r.HandleFunc("/transactions", getTransactions).Methods("GET")
	r.HandleFunc("/transactions/{id}", getTransaction).Methods("GET")
	r.HandleFunc("/transactions", createTransaction).Methods("POST")
	r.HandleFunc("/transactions/{id}", updateTransaction).Methods("PUT")
	r.HandleFunc("/transactions/{id}", deleteTransaction).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}
