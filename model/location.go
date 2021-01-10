package model

// Location represents a nested object of the API reponse from https://data.sfgov.org/resource/jjew-r69b
type Location struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	// User beware this is actually a json object but it's returned as string
	// but I'm not using in this project so I just left it as is
	HumanAddress string `json:"human_address"`
}
