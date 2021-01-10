package repository

import (
	"encoding/json"
	"strconv"
	"time"

	"food-truck/model"

	// This is an example of aliasing a pkg, can be used when the pkg name contains special characters
	// or like in this case I want to override the use of the default log library to use a more enhanced one
	soda "github.com/SebastiaanKlippert/go-soda"
	log "github.com/sirupsen/logrus"
)

// FoodTruckRepository is repsonsilbe to interacting with the
// the data store that holds the Food Truck information
type FoodTruckRepository interface {
	GetFoodTrucks() ([]model.FoodTruck, error)
	FindOpenFoodTrucks(currentTime time.Time) ([]model.FoodTruck, error)
}

// FoodTruckRepositoryImpl is the implementation of the FoodTruckRepository interface
// that interacts with a SODA API as the form of data store using this pattern we can
// easily swap out a future implementation that uses a relational database for example
type FoodTruckRepositoryImpl struct {
	url string
}

// NewFoodTruckRepository creates a new instance of the FoodTruckRepository
func NewFoodTruckRepository(url string) FoodTruckRepositoryImpl {
	return FoodTruckRepositoryImpl{
		url: url,
	}
}

// GetFoodTrucks will return all of the FoodTrucks from the repository
// I implemented this mainly as a starting but opted to leave it
func (repo FoodTruckRepositoryImpl) GetFoodTrucks() ([]model.FoodTruck, error) {
	// Why reinvent the wheel when someone else has solved the problem? Thus I opted to use
	// a pre-existing pkg that interacts with SODA APIs
	sodareq := soda.NewGetRequest(repo.url, "")
	sodareq.Format = "json"
	sodareq.Query.Limit = 10

	// make the request
	resp, err := sodareq.Get()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()

	// decode the results
	results := []model.FoodTruck{}
	err = json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		log.Fatal(err)
	}

	return results, err
}

// FindOpenFoodTrucks will return only FoodTrucks from the repository that are currently open
// in the future I would like to expand upon this func by sending it parameter Object that can
// consist of a set of different search criteria that are then dynamically applied
func (repo FoodTruckRepositoryImpl) FindOpenFoodTrucks(currentTime time.Time) ([]model.FoodTruck, error) {
	// Why reinvent the wheel when someone else has solved the problem? Thus I opted to use
	// a pre-existing pkg that interacts with SODA APIs
	sodareq := soda.NewGetRequest(repo.url, "")

	// set result format to json
	sodareq.Format = "json"

	// filter to return only open food trucks for the current day and hour
	sodareq.Query.Where = `
	start24 <= '` + getFormattedTime(currentTime) + `'
		AND dayorder = '` + strconv.Itoa(int(currentTime.Weekday())) + `'
		AND end24 >='` + getFormattedTime(currentTime) + `'`

	// sort the results in alphabetically ascending order
	sodareq.Query.AddOrder("applicant", soda.DirAsc)

	// get the results
	resp, err := sodareq.Get()
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// now need to deserialize json into []FoodTruck
	results := []model.FoodTruck{}
	err = json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		log.Fatal(err)
	}

	return results, err
}

// Golang's simplicty can feel like a double edged sword some times as you are forced to write some rudimentary boilerplate
// but this is more of a small adjustment one makes when coming to the language as the benefits of the simplicty are still considerable
func getFormattedTime(currentTime time.Time) string {
	layout := "15:04"
	return currentTime.Format(layout)
}
