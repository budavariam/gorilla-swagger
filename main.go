// Package classification Account API.
//
// this is to show how to write RESTful APIs in golang.
// that is to provide a detailed overview of the language specs
//
// Terms Of Service:
//
//     Schemes: http, https
//     Host: localhost:8080
//     Version: 1.0.0
//     Contact: Supun Muthutantri<mydocs@example.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - api_key:
//
//     SecurityDefinitions:
//     api_key:
//          type: apiKey
//          name: KEY
//          in: header
//
// swagger:meta
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Account request model
type Account struct {
	// Id of the account
	ID string `json:"id"`
	// First Name of the account holder
	FirstName string `json:"first_name"`
	// Last Name of the account holder
	LastName string `json:"last_name"`
	// User Name of the account holder
	UserName string `json:"user_name"`
}

// Account response payload
// swagger:response accountRes
type swaggAccountRes struct {
	// in:body
	Body Account
}

// Success response
// swagger:response okResp
type swaggRespOk struct {
	// in:body
	Body struct {
		// HTTP status code 200 - OK
		Code int `json:"code"`
	}
}

// Error Bad Request
// swagger:response badReq
type swaggReqBadRequest struct {
	// in:body
	Body struct {
		// HTTP status code 400 -  Bad Request
		Code int `json:"code"`
	}
}

// Error Not Found
// swagger:response notFoundReq
type swaggReqNotFound struct {
	// in:body
	Body struct {
		// HTTP status code 404 -  Not Found
		Code int `json:"code"`
	}
}

var accountMap map[string]Account

func init() {
	accountMap = make(map[string]Account)
}
func main() {
	r := mux.NewRouter()
	// swagger:operation POST /accounts/ accounts createAccount
	// ---
	// summary: Creates a new account.
	// description: If account creation is success, account will be returned with Created (201).
	// parameters:
	// - name: account
	//   description: account to add to the list of accounts
	//   in: body
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/Account"
	// responses:
	//   "200":
	//     "$ref": "#/responses/okResp"
	//   "400":
	//     "$ref": "#/responses/badReq"
	r.HandleFunc("/accounts/", createAccountHandler).Methods("POST")
	// swagger:operation GET /accounts/{id} accounts getAccount
	// ---
	// summary: Return an Account provided by the id.
	// description: If the account is found, account will be returned else Error Not Found (404) will be returned.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of the account
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/accountRes"
	//   "400":
	//     "$ref": "#/responses/badReq"
	//   "404":
	//     "$ref": "#/responses/notFoundReq"
	r.HandleFunc("/accounts/{id}", getAccountHandler).Methods("GET")
	// swagger:operation DELETE /accounts/{id} accounts deleteAccount
	// ---
	// summary: Deletes requested account by account id.
	// description: Depending on the account id, HTTP Status Not Found (404) or HTTP Status OK (200) may be returned.
	// parameters:
	// - name: id
	//   in: path
	//   description: account id
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/okResp"
	//   "400":
	//     "$ref": "#/responses/badReq"
	//   "404":
	//     "$ref": "#/responses/notFoundReq"
	r.HandleFunc("/accounts/{id}", deleteAccountHandler).Methods("DELETE")

	sh := http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./swaggerui/")))
	r.PathPrefix("/swaggerui/").Handler(sh)
	log.Println("Serving on localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
func createAccountHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Request received to create an Account")

	var account Account
	json.NewDecoder(r.Body).Decode(&account)
	id := account.ID
	accountMap[id] = account
	log.Print("Successfully created the Account ", account)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}
func getAccountHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	log.Print("Request received to get an account by account id: ", id)
	account, key := accountMap[id]
	w.Header().Add("Content-Type", "application/json")
	if key {
		log.Print("Successfully retrieved the account ", account, " for account id: ", id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(account)
	} else {
		log.Print("Requested account is not found for account id: ", id)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w)
	}
}
func deleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Request received to delete an Account by account id")
	//add your own flavor to this function :)
}
