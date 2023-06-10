package health

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type Health struct {
	Status string `json:"status"`
}

var isHealthy bool = false

func Init(ctx context.Context, ready <-chan bool) <-chan interface{} {
	done := make(chan interface{}, 1)
	go func(threadCtx context.Context, done chan<- interface{}) {
	loop:
		for {
			select {
			case rdy := <-ready:
				isHealthy = rdy
				continue loop
			case <-threadCtx.Done():
				log.Println("stopping health")
				break loop
			}
		}
		done <- nil
	}(ctx, done)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		var body []byte
		var err error

		w.Header().Add("Content-Type", "application/json")
		if isHealthy {
			body, err = json.Marshal(Health{Status: "healthy"})
			if err != nil {
				log.Printf("healthcheck success marshal error: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		} else {
			body, err = json.Marshal(Health{Status: "unhealthy"})
			if err != nil {
				log.Printf("healthcheck failure marshal error: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusTeapot)
			}
		}
		_, err = w.Write(body)

		if err != nil {
			log.Printf("healthcheck write error: %v", err)
		}
	})

	return done
}
