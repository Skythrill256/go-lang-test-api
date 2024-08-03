package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type apiError struct {
	Error string
}
type apiFunc func(w http.ResponseWriter, r *http.Request) error

func writeJSON(w http.ResponseWriter, status int, v any) error {

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)

}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			//handle error
			writeJSON(w, http.StatusBadRequest, apiError{Error: err.Error()})
		}
	}
}

type ApiServer struct {
	listenAdder string
	store       storage
}

func newApiServer(listenAdder string, store storage) *ApiServer {
	return &ApiServer{listenAdder: listenAdder, store: store}
}

func (s *ApiServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id} ", makeHTTPHandleFunc(s.handleGetAccount))
	log.Println("JSON API is running on", s.listenAdder)

	http.ListenAndServe(s.listenAdder, router)
}

func (s *ApiServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)

	}
	return fmt.Errorf("METHOD NOT ALLOWED %s", r.Method)
}
func (s *ApiServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	// db get
	fmt.Println(id)
	return writeJSON(w, http.StatusOK, &Account{})

}

func (s *ApiServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	createAccountReq := new(createAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(createAccountReq); err != nil {
		return err
	}
	account := newAccount(createAccountReq.FirstName, createAccountReq.LastName)
	if err := s.store.CreateAccount(account); err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, account)
}

func (s *ApiServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) handleMoneyTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
