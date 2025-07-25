package main

import (
	"fmt"
	httpapi "the_chamber_of_keys/api/http"
	ck "the_chamber_of_keys/pkg/chamber_of_keys"
)

func main() {

	chamber, err := ck.NewChamber()
	if err != nil {
		fmt.Println("Failed to initialise app")
	}

	chamber.Start()
	defer chamber.Stop()

	r := httpapi.NewRouter(chamber)
	r.Run(":8080")

}
