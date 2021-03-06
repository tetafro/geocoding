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

	// Assemble search criterea
	criterea := &Place{}

	fullname := req.URL.Query().Get("fullname")
	if len(fullname) < 3 {
		api.Respond(w, http.StatusBadRequest, api.Error("Fullname is too short"))
		return
	}
	criterea.Fullname = &fullname

	places, err := c.service.Get(criterea)
	if err != nil {
		c.log.Errorf("Failed to get places: %v", err)
		api.RespondInternalServerError(w)
		return
	}
	api.Respond(w, http.StatusOK, places)
}
