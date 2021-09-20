package main

import (
	"encoding/json"
	"net/http"
	"log"
	"fmt"
)


type Website struct {
	ID   string `json:"id"`
	url string `json:"URL"`
	hash string `json:"hash"`
}

var hashed []Website

func main() {

}
