package places

// Place represents geographical place.
type Place struct {
	ID               *int     `json:"id"`
	Name             *string  `json:"name"`
	AlternativeNames *string  `json:"alternative_names"`
	OSMType          *string  `json:"osm_type"`
	OSMID            *int     `json:"osm_id"`
	Class            *string  `json:"class"`
	Type             *string  `json:"type"`
	Lon              *float64 `json:"lon"`
	Lat              *float64 `json:"lat"`
	PlaceRank        *int     `json:"place_rank"`
	Importance       *float64 `json:"importance"`
	Street           *string  `json:"street"`
	City             *string  `json:"city"`
	County           *string  `json:"county"`
	State            *string  `json:"state"`
	Country          *string  `json:"country"`
	CountryCode      *string  `json:"country_code"`
	DisplayName      *string  `json:"display_name"`
	West             *float64 `json:"west"`
	South            *float64 `json:"south"`
	East             *float64 `json:"east"`
	North            *float64 `json:"north"`
	Wikidata         *string  `json:"wikidata"`
	Wikipedia        *string  `json:"wikipedia"`
	Housenumbers     *string  `json:"housenumbers"`
}
