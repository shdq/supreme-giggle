package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func fib(n int) int {
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func handleFib(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	s := q.Get("n")
	if s == "" {
		http.Error(w, "Query parameter 'n' is required", http.StatusBadRequest)
		return
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		http.Error(w, "Query parameter 'n' has to be a positive integer", http.StatusBadRequest)
		return
	}

	if n < 1 {
		http.Error(w, "Query parameter 'n' has to be a positive integer", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")

	fmt.Fprintf(w, "The Fibonacci number at position %d is %d", n, fib(n))
}

func main() {
	http.HandleFunc("/fib", handleFib)

	fmt.Println("Server is running on port 8090...")
	http.ListenAndServe(":8090", nil)
}
