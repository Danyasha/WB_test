package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func writeOKResponse(w http.ResponseWriter, m interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&m)
}

func writeErrorResponse(w http.ResponseWriter, errorCode int, errorMsg string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(errorCode)
	json.NewEncoder(w).Encode(errorMsg)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func getTvs(w http.ResponseWriter, r *http.Request) {
	tvs, err := dbGetTvs()
	if err != nil {
		writeErrorResponse(w, 501, "Error")
	}
	writeOKResponse(w, tvs)
}

func getTv(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil || id <= 0 {
		writeErrorResponse(w, 401, "Invalid request")
		return
	}
	tv, err := dbGetTv(id)
	if err == sql.ErrNoRows {
		writeErrorResponse(w, 404, "No such model")
		return
	}
	if err != nil {
		writeErrorResponse(w, 501, "Error")
		return
	}
	writeOKResponse(w, tv)
}

func createTv(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil || id <= 0 {
		writeErrorResponse(w, 401, "Invalid request")
		return
	}
	manufacturer := params["manufacturer"]
	if manufacturer == "" || len(manufacturer) < 3 {
		writeErrorResponse(w, 401, "Invalid request")
		return
	}
	model := params["model"]
	if model == "" || len(model) < 2 {
		writeErrorResponse(w, 401, "Invalid request")
		return
	}
	brand := params["brand"]
	year := params["year"]
	isDate, err := checkDate(year)
	if !isDate {
		writeErrorResponse(w, 401, "Date format YYYY-MM-DD or YYYY")
		return
	}
	tv := tv{
		Id: int64(id), Brand: brand, Year: year, Model: model, Manufacturer: manufacturer,
	}
	err = dbCreateTv(tv)
	if err != nil {
		writeErrorResponse(w, 501, "Error")
		return
	}
	writeOKResponse(w, "Success!")
}
