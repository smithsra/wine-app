package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/smithsra/wine-app/config"
	"github.com/smithsra/wine-app/types"
)

// Handler redirects to the match wine form
func Handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/matchwine", http.StatusSeeOther)
}

// MatchWine handles the request to find matching wines
func MatchWine(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "meal.gohtml", nil)
}

// MatchWineProcess makes a request from the API
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

	var data types.Paired

	res, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		decoder := json.NewDecoder(res.Body)
		err = decoder.Decode(&data)
		config.TPL.ExecuteTemplate(w, "wines.gohtml", data)
	}
}

// WinePic serves up the background image
func WinePic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("./images/wine.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}
