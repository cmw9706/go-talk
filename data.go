package main

import "fmt"

//PopulateData loads all the data into memory
func PopulateData() {
	fmt.Println("Populating mock data...")

	people := []Person{
		{
			ID:          1,
			Name:        "Connor Williams",
			Age:         26,
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		},
		{
			ID:          2,
			Name:        "Gary Meyer",
			Age:         27,
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		},
		{
			ID:          3,
			Name:        "Connor Williams",
			Age:         26,
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		},
		{
			ID:          4,
			Name:        "Gary Meyer",
			Age:         27,
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		},
	}

	for _, p := range people {
		db.SavePerson(p)
	}

	fmt.Println("Mock data populated...")
}
