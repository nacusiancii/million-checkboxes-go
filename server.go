package main

import (
    "log"
    "net/http"
    "time"
	"os"
	"os/signal"
	"syscall"
	"context"
)

func main() {

    mux := http.NewServeMux()

    // Register handlers
	mux.HandleFunc("GET /", handleRoot)
	mux.HandleFunc("GET /hello/{name}", handleHello)
	mux.HandleFunc("GET /area", handleAreaOfRectangle)

    	// Create a server with some reasonable defaults
	srv := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
	}
    
    // Start the server in a goroutine
	go func() {
		log.Println("Starting server on :8080")
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

    // Wait for interrupt signal to gracefully shutdown the server
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-sigChannel
	log.Println("Server is shutting down. Received signal: %v\n",sig)

    // Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer calls are executed LIFO at end of function
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}
