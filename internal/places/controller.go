package places

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/tetafro/geocoding/internal/api"
)

// Controller handles HTTP API requests.
type Controller struct {
	service *Service
	log     *logrus.Logger
}

// NewController creates new controller.
func NewController(service *Service, log *logrus.Logger) *Controller {
	return &Controller{service, log}
}

// Get handles request for finding place.
func (c *Controller) Get(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	name := req.URL.Query().Get("name")
	if len(name) < 3 {
		api.Respond(w, http.StatusBadRequest, api.Error("Name is too short"))
		return
	}

	places, err := c.service.GetByName(name)
	if err != nil {
		c.log.Errorf("Failed to get places: %v", err)
		api.RespondInternalServerError(w)
		return
	}
	api.Respond(w, http.StatusOK, places)
}
