package main

import (
	"fmt"
	"net/http"
	"github.com/AleksandraShmakova/backend/internal/parallel"
)

func main() {
	fmt.Println("Starting server at :8080")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Вызов функции Say и отправка результата в HTTP-ответ
		result := parallel.Say()
		fmt.Fprintln(w, result)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
