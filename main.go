package main

import (
	"BANK_TRANSACTION_APPLICATION/db"
	"BANK_TRANSACTION_APPLICATION/router"
	"log"
	"net/http"
)

func main() {
	// Initialize DB first
	err := db.InitDB("postgresql://db1:JlKOppScP00WdzqAvoxLq0mEPIqFshRv@dpg-d294snggjchc73cc6pu0-a.oregon-postgres.render.com/db1_91el")
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Conn.Close()

	// Setup router and start server
	r := router.SetupRoutes()
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
