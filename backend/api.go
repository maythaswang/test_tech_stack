package main

import (
	"log"
	"net/http"
	"backend/messagepost"
)

type APIServer struct {
	addr string
}

// Create new API Server
func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

// Run Server
func (s *APIServer) run() error {
	var router *http.ServeMux = http.NewServeMux()

	router.HandleFunc("GET /users/{userID}", func(w http.ResponseWriter, r *http.Request) {
		var userID string = r.PathValue("userID")
		w.Write([]byte("User ID: " + userID))
	})

	var messagePostController *messagepost.MessagePostController = &messagepost.MessagePostController{}

	router.HandleFunc("POST /api/post_message", func(w http.ResponseWriter, r *http.Request) {
		messagePostController.PostMessage(w, r)
	})

	var mwChain Middleware = middlewareChain(
		requestLoggerMiddleware,
		requireAuthMiddleware,
	)

	var server http.Server = http.Server{
		Addr:    s.addr,
		Handler: mwChain(router),
	}

	log.Printf("Server has started %s", s.addr)

	return server.ListenAndServe()
}

func requestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method %s, path: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

func requireAuthMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// check if the user is authenticated
		var token string = r.Header.Get("Authorization")
		if token != "Bearer token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

type Middleware func(http.Handler) http.HandlerFunc

func middlewareChain(middleware ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for i := len(middleware) - 1; i >= 0; i-- {
			next = middleware[i](next)
		}
		return next.ServeHTTP
	}
}
