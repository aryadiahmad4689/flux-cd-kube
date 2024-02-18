package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Mendefinisikan handler untuk root URL
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "test")
	})

	// Menjalankan server pada port 8080 dan menangani error jika ada
	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
