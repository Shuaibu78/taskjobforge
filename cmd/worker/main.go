package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shuaibu78/taskjobforge/pkg/api"
)

func main() {
	fmt.Println("🚀 TaskJobForge Worker starting...")

	// Start REST API server
	router := api.NewRouter()
	log.Println("API server listening on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
