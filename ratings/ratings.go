package ratings

import "fmt"

func GetAllRatings(book string) (map[string]float32, error) {

	ratings := map[string]float32{}
	for site, ratingFunc := range ratingFuncs {

		rating, err := ratingFunc(book)
		if err != nil {
			fmt.Println("Could not get", site, "rating")
			continue
		}
		ratings[site] = rating

	}
	return ratings, nil
}
