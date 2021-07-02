package handler

import (
	"encoding/json"
	"github.com/p12s/fintech-link-shorter"
	"io/ioutil"
	"log"
	"net/http"
)

// Short @Summary Short
// @Tags short
// @Description Getting a short link by a long one
// @ID get-short-link
// @Accept  json
// @Produce  json
// @Param input body shorter.UserLink true "descr long link"
// @Success 200 {string} string "https://p12s.ru/1b"

// @Router /short [post]
func (h *Handler) Short(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Error: only POST method is supported.", http.StatusBadRequest)
		return
	}

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
}

// Long @Summary Long
// @Tags long
// @Description Getting a long link from a short one
// @ID get-long-link
// @Accept  json
// @Produce  json
// @Param input body shorter.UserLink true "descr long link 2"
// @Success 200 {string} string "https://www.golangprograms.com/example-to-handle-get-and-post-request-in-golang.html"

// @Router /long [post]
func (h *Handler) Long(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Error: only POST method is supported.", http.StatusBadRequest)
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	link := &shorter.UserLink{}
	err = json.Unmarshal(reqBody, link)
	if err != nil {
		log.Fatal(err)
	}

	shortenedLink, err := h.services.Long(link.Url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(shortenedLink); err != nil {
		http.Error(w, "Error: an error occurred - the link could not be shortened.", http.StatusInternalServerError)
		return
	}
}
