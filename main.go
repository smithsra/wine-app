package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"./config"
)

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

func handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/matchwine", http.StatusSeeOther)
}

func MatchWine(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "meal.gohtml", nil)
}

func MatchWineProcess(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("Meal or cuisine")
	u, err := url.Parse("https://api.spoonacular.com/food/wine/pairing")
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Set("food", v)
	q.Set("apiKey", "c2c78ceb6d784338ab7e7cb1c359f9ed")
	u.RawQuery = q.Encode()

	var data Paired

	res, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		decoder := json.NewDecoder(res.Body)
		err = decoder.Decode(&data)
		config.TPL.ExecuteTemplate(w, "wines.gohtml", data)
	}
}

func winePic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("./images/wine.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/matchwine", MatchWine)
	http.HandleFunc("/matchwine/process", MatchWineProcess)
	http.HandleFunc("/wine.jpg", winePic)
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
