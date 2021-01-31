package types

// Paired struct for wine data
type Paired struct {
	PairedWines    []string `json:"pairedWines"`
	PairingText    string   `json:"pairingText"`
	ProductMatches []struct {
		ID            int         `json:"id"`
		Title         string      `json:"title"`
		AverageRating float64     `json:"averageRating"`
		Description   interface{} `json:"description"`
		ImageURL      string      `json:"imageUrl"`
		Link          string      `json:"link"`
		Price         string      `json:"price"`
		RatingCount   float64     `json:"ratingCount"`
		Score         float64     `json:"score"`
	} `json:"productMatches"`
}
