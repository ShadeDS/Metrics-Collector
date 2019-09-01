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
	logger *log.Logger
)

func init() {
	logger = log.New(os.Stdout, "controller: ", log.Lshortfile)
}

type Controller struct {
	service *service.Service
}

func New(s *service.Service) *Controller {
	logger.Println("New controller instance was initialized")
	return &Controller{
		service: s,
	}
}

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

	c.service.SubmitMetric(metric)
	w.WriteHeader(http.StatusCreated)
	logger.Println("Metric was submitted successfully")
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	logger.Printf("Send response code: %v, body: %v", code, string(response))
	w.Write(response)
}
