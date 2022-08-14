package main

import (
	"errors"
	"strings"
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	content, err := os.ReadFile("src/index.html")
	if err != nil {
		io.WriteString(w, "Failed to load the index file!\n")
		return
	}

	domain_restrict_maps_key := "AIzaSyB3AzwS5rXDm3vcid4d1-Up_ujFlkOQaC4"
	local_development_key := os.Getenv("LOCAL_DEV_MAPS_KEY")
	
	io.WriteString(w, strings.Replace(string(content), domain_restrict_maps_key, local_development_key, 1))
}

func load_env(){
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file")
		return
	}
}

func main() {
	load_env()

	http.HandleFunc("/", getRoot)

	err := http.ListenAndServe(":2121", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}