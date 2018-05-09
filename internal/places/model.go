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
	Country     *string `json:"country"`
	City        *string `json:"city"`
	Street      *string `json:"street"`
	Housenumber *string `json:"housenumber"`
	Name        *string `json:"name"`
	Fullname    *string `json:"fullname"`
	Coordinate  *Point  `json:"coordinate"`
}
