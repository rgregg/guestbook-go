package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	// Imports the Google Cloud Datastore client package
	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"

	// Imports UUID generator
	"github.com/satori/go.uuid"
)

const ()

var (
	projectID = os.Getenv("PROJECT_ID")
	ctx       context.Context
	client    *datastore.Client
)

type GuestEntry struct {
	Message    string
	PostedTime time.Time
}

func writeEntryHandler(w http.ResponseWriter, r *http.Request) {
	// Write a new entry into the guest log

	kind := "GuestEntry"

	id, _ := uuid.NewV4()
	name := id.String()
	entryKey := datastore.NameKey(kind, name, nil)

	entry := GuestEntry{
		Message:    r.URL.Query().Get("message"),
		PostedTime: time.Now()}

	if _, err := client.Put(ctx, entryKey, &entry); err != nil {
		log.Printf("Failed to save the entry: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to save the entry"))
		return
	}

	log.Printf("Saved %v: %v\n", entryKey, entry.Message)
	w.Write([]byte(fmt.Sprintf("Saved %v: %v\n", name, entry.Message)))
}

func listEntriesHandler(w http.ResponseWriter, r *http.Request) {
	var entities []GuestEntry
	q := datastore.NewQuery("GuestEntry").Order("-PostedTime").Limit(20)
	_, err := client.GetAll(ctx, q, &entities)
	if err != nil {
		log.Printf("Failed to retrieve entries: %v\n", err)
	}

	w.Header().Set("Content-Type", "application/json")

	json, err := json.Marshal(entities)
	fmt.Fprintf(w, "%s", json)
}

func main() {

	// Initialize data store
	ctx = context.Background()

	// Create the datastore client
	client, _ = datastore.NewClient(ctx, projectID)

	http.Handle("/", http.FileServer(http.Dir("/static")))
	http.HandleFunc("/post", writeEntryHandler)
	http.HandleFunc("/list", listEntriesHandler)
	http.ListenAndServe(":8080", nil)
}
