package handler

import (
	"encoding/json"
	"github.com/p12s/fintech-link-shorter"
	"io/ioutil"
	"net/http"
)

// Short @Summary Getting a short link
// @Tags short
// @Description Getting a short link by a long one
// @ID get-short-link
// @Accept  json
// @Produce  json
// @Param url body shorter.UserLink true "long link"
// @Success 200 {object} shorter.UserLink
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /short [post]
func (h *Handler) Short(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		NewErrorResponse(w, http.StatusBadRequest, "Error: only POST method is supported.")
		return
	}
	// @Success 200 {string} string "https://p12s.ru/1b"

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	link := &shorter.UserLink{}
	err = json.Unmarshal(reqBody, link)
	if err != nil {
		NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	shortenedLink, err := h.services.Create(link.Url)
	if err != nil {
		NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(shortenedLink); err != nil {
		http.Error(w, "Error: an error occurred - the link could not be shortened.", http.StatusInternalServerError)
		return
	}
}

// Long @Summary Getting a long link
// @Tags long
// @Description Getting a long link from a short one
// @ID get-long-link
// @Accept  json
// @Produce  json
// @Param url body shorter.UserLink true "short link"
// @Success 200 {object} shorter.UserLink
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /long [post]
func (h *Handler) Long(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		NewErrorResponse(w, http.StatusBadRequest, "Error: only POST method is supported.")
		return
	}

	// @Success 200 {string} string "https://www.golangprograms.com/example-to-handle-get-and-post-request-in-golang.html"

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		NewErrorResponse(w, http.StatusInternalServerError, err.Error()) ///////// ??????? так работает или нет?
		return
	}

	link := &shorter.UserLink{}
	err = json.Unmarshal(reqBody, link)
	if err != nil {
		NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	shortenedLink, err := h.services.Long(link.Url)
	if err != nil {
		NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(shortenedLink); err != nil {
		http.Error(w, "Error: an error occurred - the link could not be shortened.", http.StatusInternalServerError)
		return
	}
}
