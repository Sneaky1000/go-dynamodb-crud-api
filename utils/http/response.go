package http

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Status int         `json:"status"`
	Result interface{} `json:"result"`
}

func newResponse(data interface{}, status int) *response {
	return &response{
		Status: status,
		Result: data,
	}
}

func (resp *response) bytes() []byte {
	data, _ := json.Marshal(resp)

	return data
}

func (resp *response) string() string {
	return string(resp.bytes())
}

func (resp *response) sendResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(resp.Status)
	_, _ = w.Write(resp.bytes())

	log.Println(resp.string())
}

// 200 Status Code
func StatusOK(w http.ResponseWriter, r *http.Request, data interface{}) {
	newResponse(data, http.StatusOK).sendResponse(w, r)
}

// 204 Status Code
func StatusNoContent(w http.ResponseWriter, r *http.Request) {
	newResponse(nil, http.StatusNoContent).sendResponse(w, r)
}

// 400 Status Code
func StatusBadRequest(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{
		"error": err.Error(),
	}

	newResponse(data, http.StatusBadRequest).sendResponse(w, r)
}

// 404 Status Code
func StatusNotFound(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{
		"error": err.Error(),
	}

	newResponse(data, http.StatusNotFound).sendResponse(w, r)
}

// 405 Status Code
func StatusMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	newResponse(nil, http.StatusMethodNotAllowed).sendResponse(w, r)
}

// 409 Status Code
func StatusConflict(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{
		"error": err.Error(),
	}

	newResponse(data, http.StatusNotFound).sendResponse(w, r)
}

// 500 Status Code
func StatusInternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{
		"error": err.Error(),
	}

	newResponse(data, http.StatusInternalServerError).sendResponse(w, r)
}
