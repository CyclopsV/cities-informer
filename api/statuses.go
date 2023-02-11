package api

import "net/http"

func statusOK(w *http.ResponseWriter, answer []byte) {
	(*w).WriteHeader(http.StatusOK)
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Write(answer)
}

func statusCreated(w *http.ResponseWriter, answer []byte) {
	(*w).WriteHeader(http.StatusCreated)
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Write(answer)
}

func statusBadRequest(w *http.ResponseWriter, info string) {
	(*w).WriteHeader(http.StatusBadRequest)
	(*w).Write([]byte(info))
}

func statusInternalServerError(w *http.ResponseWriter, info string) {
	(*w).WriteHeader(http.StatusInternalServerError)
	(*w).Write([]byte(info))
}
