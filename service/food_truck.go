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
type FoodTruckServiceImpl struct {
	repo repository.FoodTruckRepository
}

// NewFoodTruckService creates a new instance of the FoodTruckService
func NewFoodTruckService(repo repository.FoodTruckRepository) FoodTruckServiceImpl {
	return FoodTruckServiceImpl{
		repo: repo,
	}
}

// GetFoodTrucks will return all of the FoodTrucks
func (svc FoodTruckServiceImpl) GetFoodTrucks() ([]model.FoodTruck, error) {
	return svc.repo.GetFoodTrucks()
}

// FindOpenFoodTrucks will only return the FoodTrucks that are currently open
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
