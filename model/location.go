package model

// Location represents a nested object of the API reponse from https://data.sfgov.org/resource/jjew-r69b
type Location struct {
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	HumanAddress string `json:"human_address"`
}
