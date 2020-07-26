package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
)

func translate(q string) string {

	q = url.QueryEscape(q)
	ur := "https://api.mymemory.translated.net/get?q=" + q + "&langpair=pt-br|en"
	fmt.Println(ur)
	req, err := http.NewRequest("GET", ur, nil)

	if err != nil {
		log.Fatal("Error on translate\n", err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return "Error on translate"
	}

	defer resp.Body.Close()

	var j Response

	if err := json.NewDecoder(resp.Body).Decode(&j); err != nil {
		log.Println(err)
	}

	fmt.Println("Accuracy ", j.ResponseData.Match)
	return j.ResponseData.TranslatedText
}

func getEmoji(w http.ResponseWriter, r *http.Request) {

	emojiName := r.URL.Query()

	translatedVar := translate(emojiName["s"][0])

	fmt.Println("traduzida ", translatedVar)

	if strings.Contains(translatedVar, " ") {
		translatedVar = strings.ReplaceAll(translatedVar, " ", "-")
	}

	translatedVar = strings.ToLower(url.QueryEscape(translatedVar))

	response, err := http.Get("https://emojipedia.org/" + translatedVar + "/")

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	value, exists := document.Find("input#emoji-copy").Attr("value")

	if exists {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Emoji{value, translatedVar})
	} else {
		json.NewEncoder(w).Encode(NotFoundEmoji{"Emoji not found 🏜️"})
	}
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/emoji", getEmoji).Methods("GET")

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, router))

}
