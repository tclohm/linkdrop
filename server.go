package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"log"
	"io/ioutil"
	"github.com/gorilla/mux"
	"github.com/google/uuid"
)


type Link struct {
	ID   string `json:"Id"`
	Url string `json:"url"`
	Hash string `json:"hash"`
}

var hashed []Link

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api", Up)
	router.HandleFunc("/new", createNewLink)
	router.HandleFunc("/{hash}", redirect)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Up(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("api endpoint hit: return 'Up'")
	json.NewEncoder(w).Encode("Up")
}

func createNewLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal("error!", err)
	}

	var link Link

	json.Unmarshal(body, &link)

	i := uuid.New()
	link.ID = i.String()
	link.Hash = i.String()[:6]

	hashed = append(hashed, link)

	json.NewEncoder(w).Encode(link)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["hash"]

	for _, link := range hashed {
		if link.Hash == key {
			fmt.Println("found!")
			http.Redirect(w, r, link.Url, 307)
		}
	}
}

func main() {
	// MARK: -- cookie drop
	// expiration := time.Now().add(365 * 25 * time.Hour)
	// cookie := http.Cookie{Name: "ghostdrop", Value: "peekaboo", Expires: expiration}
	// http.SetCookie(w, &cookie)
	fmt.Println("starting up")
	fmt.Println("port :8080")
	handleRequests()
}
