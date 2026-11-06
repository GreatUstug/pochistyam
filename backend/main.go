package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
	return c.value
}

func (c *Counter) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {

	counter := &Counter{}

	http.HandleFunc("POST /increment", func(w http.ResponseWriter, r *http.Request) {
		n := counter.Increment()
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(map[string]int{"count": n})
		if err != nil {
			return
		}
	})

	http.HandleFunc("GET /count", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(map[string]int{"count": counter.Get()})
		if err != nil {
			return
		}
	})
	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
