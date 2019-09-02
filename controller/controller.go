package controller

import (
	"encoding/json"
	"log"
	"metrics-collector/model"
	"metrics-collector/service"
	"net/http"
	"os"
)

var (
	logger = log.New(os.Stdout, "controller: ", log.Lshortfile)
)

type Controller struct {
	service *service.Service
}

// Creates a new controller instance
func New(s *service.Service) *Controller {
	logger.Println("New controller instance was initialized")
	return &Controller{
		service: s,
	}
}

// Handles request to submit metric about happen event. Marshals request body
// into object and sends it to storage
func (c *Controller) SubmitMetric(w http.ResponseWriter, r *http.Request) {
	logger.Println("Received request to submit metric")

	defer r.Body.Close()
	var metric model.Metric
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&metric); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		logger.Println("Error occurred while decoding request body")
		return
	}

	if err := metric.Validate(); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		logger.Println("Request body is invalid")
		return
	}

	c.service.SubmitMetric(metric)
	w.WriteHeader(http.StatusCreated)
	logger.Println("Metric was submitted successfully")
}

// Writes json with error field to response with specified code
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

// Writes json to response with specified code
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	logger.Printf("Send response code: %v, body: %v", code, string(response))
	w.Write(response)
}
