package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/terachanple/redis-ranking-demo/service"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rankings, err := service.GetCurrentDailyRanking()
	if err != nil {
		log.Printf("err: %+v", err)
		InternalServerError(w, r)
		return
	}

	json.NewEncoder(w).Encode(rankings)
}

func Show(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	if id == "" {
		BadRequest(w, r)
		return
	}

	ranking, err := service.IncrementDailyCountByID(params.ByName("id"))
	if err != nil {
		log.Printf("err: %+v", err)
		InternalServerError(w, r)
		return
	}
	json.NewEncoder(w).Encode(ranking)
}

func CreateWeeklyRanking(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := service.GenrateWeeklyRanking(); err != nil {
		log.Printf("err: %+v", err)
		InternalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func BadRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	if _, err := w.Write([]byte("Bad Request")); err != nil {
		return
	}
}

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	if _, err := w.Write([]byte("Not Found")); err != nil {
		return
	}
}
