package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"food-truck/model"
	"food-truck/repository"
	"food-truck/service"

	// This is an example of aliasing a pkg, can be used when the pkg name contains special characters
	// of like in this case I want to override the use of the default log library to use a more enhanced one
	log "github.com/sirupsen/logrus"
)

func main() {
	// With Golang what you see is what you get, which can mean if you don't set it up
	// yourself then it won't be there. No Spring magic to see here. But I like this
	// as it can enfore simplicity of design and easily elevate smells of when a microservice
	// is becoming too complicated
	repo := repository.NewFoodTruckRepository("https://data.sfgov.org/resource/jjew-r69b")
	svc := service.NewFoodTruckService(repo)

	foodTrucks, err := svc.FindOpenFoodTrucks()
	if err != nil {
		log.Error(err)
	}

	// print all of the food trucks in a formatted list
	for k, v := range foodTrucks {
		if v == nil || len(v) == 0 {
			continue
		}
		print(v)

		// check if this was the last page to print
		if k == (len(foodTrucks) - 1) {
			fmt.Println("\nTh-th-th-that's all folks!")
			break
		}

		// check if the user wants to proceed to the next page
		if !proceed() {
			fmt.Println("\nThank you please come again!")
			break
		}
	}
}

func print(foodTrucks []model.FoodTruck) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	fmt.Fprintln(writer, "NAME\tADDRESS")
	for _, v := range foodTrucks {
		fmt.Fprintln(writer, v.Applicant+"\t"+v.Location)
	}
	writer.Flush()
}

func proceed() bool {
	var nextPage *bool
	for ok := true; ok; ok = (nextPage == nil) {
		nextPage = wouldYouLikeToKnowMore()
		if nextPage == nil {
			fmt.Println("I am sorry we didn't understand that.")
		}
	}
	return *nextPage
}

func wouldYouLikeToKnowMore() *bool {
	fmt.Println("Would you like to know more? (Yes/No)")
	var input string
	fmt.Scanln(&input)
	yes := (strings.ToLower(input) == "yes" || strings.ToLower(input) == "y")
	no := (strings.ToLower(input) == "no" || strings.ToLower(input) == "n")
	switch true {
	case yes:
		return &yes
	case no:
		// this may seem like an error but it's intentional with the way
		// pointers work in go, can't get an address of &(!no) so at this point
		// yes will be it's opposite so we can pass it by pointer instead
		return &yes
	default:
		return nil
	}
}
