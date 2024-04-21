package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"time"
)

type hotel struct {
	Name       string
	City       string
	Country    string
	Rating     uint8
	IsActive   bool
	LastUpdate dateYMD
}

type dateYMD struct {
	Year  int
	Month time.Month
	Day   int
}

func main() {
	fmt.Println("Hey there!")

	var hotels []hotel

	hotel1 := hotel{
		"Ibis",
		"Porto Alegre",
		"Brasil",
		3,
		true,
		dateYMD{
			2024,
			time.February,
			1,
		},
	}

	hotel2 := hotel{
		"Laghetto",
		"Gramado",
		"Brasil",
		4,
		true,
		dateYMD{
			2023,
			time.December,
			12,
		},
	}

	hotel3 := hotel{
		"Marriott",
		"Paris",
		"France",
		5,
		true,
		dateYMD{
			2024,
			time.January,
			15,
		},
	}

	hotel4 := hotel{
		"Ball 8",
		"Atlantic City",
		"United States of America",
		1,
		true,
		dateYMD{
			2022,
			time.April,
			30,
		},
	}

	hotel5 := hotel{
		"Boulevard Hotel",
		"Sidney",
		"Australia",
		5,
		true,
		dateYMD{
			2024,
			time.March,
			2,
		},
	}

	hotels = append(hotels, hotel1, hotel2, hotel3, hotel4, hotel5)

	for index, value := range hotels {
		fmt.Printf("%v - %+v\n", index+1, value)
	}

	hotelsJSON, err := json.Marshal(hotels)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n%v\n\n", string(hotelsJSON))

	var decodeJSON = []hotel{}

	if err = json.Unmarshal(hotelsJSON, &decodeJSON); err != nil {
		log.Fatal(err)
	}

	for index, value := range decodeJSON {
		fmt.Printf("%v - %v\n", index+1, value)
	}

	// fmt.Printf("\n%+v\n", decodeJSON)

	// sort.Slice(decodeJSON, func(i, j int) bool {
	// 	return decodeJSON[i].Rating < decodeJSON[j].Rating
	// })

	sort.Slice(decodeJSON, func(i, j int) bool {
		if decodeJSON[i].Rating != decodeJSON[j].Rating {
			return decodeJSON[i].Rating < decodeJSON[j].Rating
		}
		// if decodeJSON[i].Name != decodeJSON[j].Name {
		// 	return decodeJSON[i].Name < decodeJSON[j].Name
		// }
		return decodeJSON[i].Name > decodeJSON[j].Name
	})

	fmt.Println()
	for index, value := range decodeJSON {
		fmt.Printf("%v - %v\n", index+1, value)
	}

}
