package ratings

func GetAllRatings(book string) (map[string]float32, error) {

	ratings := map[string]float32{}
	for site, ratingFunc := range ratingFuncs {

		rating, err := ratingFunc(book)
		if err != nil {
			return nil, err
		}
		ratings[site] = rating
	}
	return ratings, nil
}
