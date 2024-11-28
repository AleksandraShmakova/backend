package main

import (
	"fmt"
	"github.com/AleksandraShmakova/backend/internal/parallel"
	"github.com/AleksandraShmakova/backend/internal/server"
)

func main() {
	fmt.Println("Running parallel example:")
	parallel.Say()

	fmt.Println("Starting server at :8080")
	srv := server.New(":8080")
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
