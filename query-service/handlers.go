package main

import (
	"context"
	"cqrs-meows/db"
	"cqrs-meows/event"
	"cqrs-meows/schema"
	"cqrs-meows/search"
	"cqrs-meows/util"
	"log"
	"net/http"
	"strconv"
)

// onMeowCreated inserts a meow into Elasticsearch whenever the OnMeowCreated event is received
func onMeowCreated(m event.MeowCreatedMessage) {
	meow := schema.Meow{
		ID:        m.ID,
		Body:      m.Body,
		CreatedAt: m.CreatedAt,
	}
	if err := search.InsertMeow(context.Background(), meow); err != nil {
		log.Println(err)
	}
}

// searchMeowsHandlers performs full-text search and returns
// meows bounded with skip and take parameters
func searchMeowsHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()

	// Read parameters
	query := r.FormValue("query")
	if len(query) == 0 {
		util.ResponseError(w, http.StatusBadRequest, "Missing query parameter")
		return
	}
	skip := uint64(0)
	skipStr := r.FormValue("skip")
	take := uint64(100)
	takeStr := r.FormValue("take")
	if len(skipStr) != 0 {
		skip, err = strconv.ParseUint(skipStr, 10, 64)
		if err != nil {
			util.ResponseError(w, http.StatusBadRequest, "Invalid skip parameter")
			return
		}
	}
	if len(takeStr) != 0 {
		take, err = strconv.ParseUint(takeStr, 10, 64)
		if err != nil {
			util.ResponseError(w, http.StatusBadRequest, "Invliad take parameter")
			return
		}
	}

	// Search meows
	meows, err := search.SearchMeows(ctx, query, skip, take)
	if err != nil {
		log.Println(err)
		util.ResponseOk(w, []schema.Meow{})
		return
	}

	util.ResponseOk(w, meows)
}

// listMeowsHandler returns all meows ordered by creation time.
func listMeowsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error

	// Read parameters
	skip := uint64(0)
	skipStr := r.FormValue("skip")
	take := uint64(100)
	takeStr := r.FormValue("take")
	if len(skipStr) != 0 {
		skip, err = strconv.ParseUint(skipStr, 10, 64)
		if err != nil {
			util.ResponseError(w, http.StatusBadRequest, "Invalid skip parameter")
			return
		}
	}
	if len(takeStr) != 0 {
		take, err = strconv.ParseUint(takeStr, 10, 64)
		if err != nil {
			util.ResponseError(w, http.StatusBadRequest, "Invalid take parameter")
			return
		}
	}

	// Fetch meows
	meows, err := db.ListMeows(ctx, skip, take)
	if err != nil {
		log.Println(err)
		util.ResponseError(w, http.StatusInternalServerError, "Could not fetch meows")
		return
	}

	util.ResponseOk(w, meows)
}
