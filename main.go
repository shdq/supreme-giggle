package main

import (
	"fmt"
	"net/http"
)

func fib(n int) int {
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func handleFib(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your fibonacci number is", fib(6))
}

func main() {
	http.HandleFunc("/fib", handleFib)

	fmt.Println("Server is running on port 8090...")
	http.ListenAndServe(":8090", nil)
}
