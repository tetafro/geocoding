package places

// Point represent a single point on a map.
type Point struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// Place represents geographical place.
type Place struct {
	ID          *int    `json:"id"`
	OSMID       *int    `json:"osm_id"`
	OSMType     *string `json:"osm_type"`
	Coordinate  *Point  `json:"coordinate"`
	Priority    *int    `json:"priority"`
	Name        *string `json:"name"`
	Country     *string `json:"country"`
	City        *string `json:"city"`
	Street      *string `json:"street"`
	Housenumber *string `json:"housenumber"`
	Fullname    *string `json:"fullname"`
}
