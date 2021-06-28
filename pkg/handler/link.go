package handler

import (
	"encoding/json"
	"fmt"
	shorter "github.com/p12s/fintech-link-shorter"
	"io/ioutil"
	"log"
	"net/http"
)

func (h *Handler) Short(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		link := &shorter.UserLink{}
		err = json.Unmarshal(reqBody, link)
		if err != nil {
			log.Fatal(err)
		}

		shortenedLink, err := h.services.Create(link.Url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := json.NewEncoder(w).Encode(shortenedLink); err != nil {
			http.Error(w, "Error: an error occurred - the link could not be shortened.", http.StatusInternalServerError)
			return
		}

		fmt.Println("отправлено ок: %v\n", shortenedLink)

	default:
		http.Error(w, "Error: only POST method is supported.", http.StatusBadRequest)
		return
	}
}

func (h *Handler) Long(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "TEST TEST 1")
}

