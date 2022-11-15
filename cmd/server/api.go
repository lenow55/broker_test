package main

import (
	"log"
	"net/http"
    "strconv"
	"encoding/json"
)

// handleAPICache handles API calls for cached messages.
//func handleAPICache(cr *redis.Client) func(w http.ResponseWriter, r *http.Request) {
//	return func(w http.ResponseWriter, r *http.Request) {
//		data, err := cache.GetCacheJSON(cr)
//		if err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			w.Write([]byte("500 – Something went wrong"))
//
//			log.Printf("get cache json: %s", err)
//			return
//		}
//
//		w.Header().Set("Content-Type", "application/json")
//		fmt.Fprint(w, data)
//	}
//}

// Recursive
func fibonacciRecursive(n int) int {
    if n < 2 {
        return n
    }
    return fibonacciRecursive(n - 1) + fibonacciRecursive(n - 2)
}

// handleAPICache handles API calls for cached messages.
func handleFibonachi(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
    query := r.URL.Query()
    number, present := query["number"] //значение, которое надо вычислить фибоначи
    if !present || len(number) != 1 {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("301 – bad request params"))
        log.Printf("Error server: bad request params")
        return
    }
    n, err := strconv.Atoi(number[0])
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("500 – Something went wrong"))

        log.Printf("get cache json: %s", err)
        return
    }
    if n < 0 || n > 1000 {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("301 – bad request params"))
        log.Printf("Error server: bad range")
        return
    }

    var result int = fibonacciRecursive(n)
    renderResponse(w, result, 200)
}

func renderResponse(w http.ResponseWriter, res interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")

	content, err := json.Marshal(res)
	if err != nil {
		// XXX Do something with the error ;)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(status)

	if _, err = w.Write(content); err != nil { //nolint: staticcheck
		// XXX Do something with the error ;)
	}
}
