package handle

import (
	"encoding/json"
	"net/http"

	"github.com/doanvanvinhtho/simple-rest-api-by-gorilla-mux/model"
	"github.com/doanvanvinhtho/simple-rest-api-by-gorilla-mux/repository"
	"github.com/gorilla/mux"
)

// GetOneEvent helps to get an event from repository
func GetOneEvent(w http.ResponseWriter, r *http.Request) {
	var response model.Response

	// Get the ID from the url
	eventID := mux.Vars(r)["id"]

	if event, err := repository.GetOneEvent(eventID); err != nil {
		response = model.Response{
			Code:    http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		}
	} else {
		response = model.Response{
			Code:    http.StatusOK,
			Message: "",
			Data:    *event,
		}
	}

	json.NewEncoder(w).Encode(response)
}
