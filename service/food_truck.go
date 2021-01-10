package service

import (
	"food-truck/model"
	"food-truck/repository"
	"time"
)

// FoodTruckService is respoonsible for applying business logic to any queries or results
type FoodTruckService interface {
	GetFoodTrucks() ([]model.FoodTruck, error)
	FindOpenFoodTrucks() ([][]model.FoodTruck, error)
}

// FoodTruckServiceImpl is the implementation of the FoodTruckService interface
// with it only requiring an implementation of a defined interface we are taking
// advantage of dependency injection by injecting the type of repository we need/use
type FoodTruckServiceImpl struct {
	repo repository.FoodTruckRepository
}

// NewFoodTruckService creates a new instance of the FoodTruckService
// some more go boiler plate but it took just as long to write the comment
// explaining it, so no harm no foul
func NewFoodTruckService(repo repository.FoodTruckRepository) FoodTruckServiceImpl {
	return FoodTruckServiceImpl{
		repo: repo,
	}
}

// GetFoodTrucks will return all of the FoodTrucks
// I used this as a starting point but opted to keep it for my own debugging purposes
// to compare against the result in FindOpenFoodTrucks
func (svc FoodTruckServiceImpl) GetFoodTrucks() ([]model.FoodTruck, error) {
	return svc.repo.GetFoodTrucks()
}

// FindOpenFoodTrucks will only return the FoodTrucks that are currently open
// using page blocks of 10 results, ideally in the future I would to implement
// a way to page back and forth through the set
func (svc FoodTruckServiceImpl) FindOpenFoodTrucks() ([][]model.FoodTruck, error) {
	// get the set of open food trucks from the repository
	foodTrucks, err := svc.repo.FindOpenFoodTrucks(time.Now())
	if err != nil {
		return nil, err
	}

	// now let's sort the trucks into pages of 10
	results := [][]model.FoodTruck{}
	tempTruckSet := []model.FoodTruck{}
	for k, v := range foodTrucks {
		// check if we are at position 10
		if k%10 == 0 {
			// add this page to the array
			results = append(results, tempTruckSet)

			// reset so we have another page
			tempTruckSet = []model.FoodTruck{}
		}

		// add the FoodTruck to placeholder page
		tempTruckSet = append(tempTruckSet, v)
	}

	// check if we have a leftover or incomplete set
	// e.g. a last page with less than 10
	if tempTruckSet != nil && len(tempTruckSet) > 0 {
		results = append(results, tempTruckSet)
	}

	return results, nil
}
