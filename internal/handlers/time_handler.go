package handlers

import (
	"encoding/json"
	"net/http"
	"relogio_mundial/internal/errors"
	"relogio_mundial/internal/models"
	"relogio_mundial/internal/services"
)

type TimeResponse struct {
	City      string `json:"city"`
	LocalTime string `json:"local_time"`
}

func HandleTimeForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Acess-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Content-Type", "applicaiton/json")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var form models.TimeForm
	err := json.NewDecoder(r.Body).Decode(&form)

	if err != nil {
		errors.SendErrorResponse(w, http.StatusBadRequest, "Invalid Request Body")
		return
	}

	if form.City == "" {
		errors.SendErrorResponse(w, http.StatusBadRequest, "The City or Timezone is empty")
		return
	}
	localtime, err := services.GetLocalTime(form.City)
	if err != nil {
		errors.SendErrorResponse(w, http.StatusBadRequest, "City or TImezone not found")
		return
	}
	response := TimeResponse{
		City:      form.City,
		LocalTime: localtime,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
