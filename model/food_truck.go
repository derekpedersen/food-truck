package model

// FoodTruck represents the API reponse from https://data.sfgov.org/resource/jjew-r69b
// Golang is a strongly typed language so ideally these fields
// would typed defined but the SODA api appears to only return strings however
// there is a well-known solution to this problem, it's too simply overwrite
// the json Marshall and Unmarshall funcs https://golang.org/pkg/encoding/json/ I chose not
// to implement that here due to time limitations
type FoodTruck struct {
	DayOrder         string   `json:"dayorder"`
	DayOfWeekStr     string   `json:"dayofweekstr"`
	StartTime        string   `json:"starttime"`
	EndTime          string   `json:"endtime"`
	Permit           string   `json:"permit"`
	Location         string   `json:"location"`
	LocationDesc     string   `json:"locationdesc"`
	OptionalText     string   `json:"optionaltext"`
	LocationID       string   `json:"locationid"`
	Start24          string   `json:"start24"`
	End24            string   `json:"end24"`
	CNN              string   `json:"cnn"`
	AddrDateCreate   string   `json:"addr_date_create"`
	AddrDateModified string   `json:"addr_date_modified"`
	Block            string   `json:"block"`
	Lot              string   `json:"lot"`
	ColdTruck        string   `json:"coldtruck"`
	Applicant        string   `json:"applicant"`
	X                string   `json:"x"`
	Y                string   `json:"y"`
	Latitude         string   `json:"latitude"`
	Longitude        string   `json:"longitude"`
	Location2        Location `json:"location_2"`
}
