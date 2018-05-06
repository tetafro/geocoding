package places

// Point represent a single point on a map.
type Point struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// Place represents geographical place.
type Place struct {
	ID         *int    `json:"id"`
	OSMID      *int    `json:"osm_id"`
	OSMType    *string `json:"osm_type"`
	Name       *string `json:"name"`
	Coordinate *Point  `json:"coordinate"`
}
