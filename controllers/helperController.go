package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

//RespondWithJSON definition
func RespondWithJSON(w http.ResponseWriter, statusCode int, payLoad interface{}) {
	response, _ := json.Marshal(payLoad)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

//RespondWithError definition
func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	RespondWithJSON(w, statusCode, map[string]string{"An error ocurred": message})
}

// PerPage defines the default number of items per page
const PerPage = 15

func GetPageNum(r *http.Request) int {
	p := r.URL.Query().Get("page")

	if p == "" {
		return 1
	}
	pageNum, err := strconv.Atoi(p)
	if err != nil {
		log.Printf("%+v\n", err)

		pageNum = 1
	}
	return pageNum
}

type Error interface {
	error
	Status() int
}

type StatusError struct {
	Code int
	Err  error
}

func (se StatusError) Error() string {
	return se.Err.Error()
}

func (se StatusError) Status() int {
	return se.Code
}

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		switch e := err.(type) {
		case StatusError:
			log.Printf("HTTP %d - %+v\n", e.Status(), e.Err)
			RespondWithError(w,http.StatusInternalServerError,strconv.Itoa(e.Status()))
		default:
			log.Printf("%+v\n", err)
			RespondWithError(w,http.StatusBadRequest,"cannot be resolved")
		}
	}
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	RespondWithError(w,http.StatusNotFound,"the handler you are looking for was not found")
}
