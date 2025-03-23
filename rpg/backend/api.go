package main

import (
	"backend/db"
	"backend/messagepost"
	"log"
	"net/http"
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
	// Setup new Router
	var router *http.ServeMux = http.NewServeMux()

	// Setup connection with DB
	dbHandler, err := db.NewDBHandler()
	if err != nil {
		log.Fatalf("failed to establish connection with the database")
	}
	log.Printf("successfully established connection with the database")
	dbHandler.CreateTables()

	// Setup Middleware chain
	var mwChain Middleware = middlewareChain(
		requestLoggerMiddleware,
		// requireAuthMiddleware,
	)

	// Setup Server
	var server http.Server = http.Server{
		Addr:    s.addr,
		Handler: mwChain(router),
	}

	log.Printf("Server has started %s", s.addr)

	// create controller and service for MessagePost
	var messagePostService *messagepost.MessagePostService = messagepost.NewMessagePostService(dbHandler.GetDB())
	var messagePostController *messagepost.MessagePostController = messagepost.NewMessagePostController(messagePostService)

	// /api/post_message
	router.HandleFunc("POST /api/post_message", func(w http.ResponseWriter, r *http.Request) {
		var err error = messagePostController.PostMessage(w, r)
		if err != nil {
			log.Printf(err.Error())
			log.Printf("ERROR | method %s, path: %s, %v", r.Method, r.URL.Path, err.Error())
		}
	})

	// /api/get_message/{message_id}
	router.HandleFunc("GET /api/get_message/{message_id}", func(w http.ResponseWriter, r *http.Request) {
		var err error = messagePostController.GetMessage(w, r)
		if err != nil {
			log.Printf("ERROR | method %s, path: %s, %v", r.Method, r.URL.Path, err.Error())
		}
	})

	// /api/delete_message/{message_id}
	router.HandleFunc("DELETE /api/delete_message/{message_id}", func(w http.ResponseWriter, r *http.Request) {
		var err error = messagePostController.DeleteMessage(w, r)
		if err != nil {
			log.Printf("ERROR | method %s, path: %s, %v", r.Method, r.URL.Path, err.Error())
		}
	})

	return server.ListenAndServe()
}

// ------------------------ MIDDLEWARES ------------------------
type Middleware func(http.Handler) http.HandlerFunc

func middlewareChain(middleware ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for i := len(middleware) - 1; i >= 0; i-- {
			next = middleware[i](next)
		}
		return next.ServeHTTP
	}
}

func requestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method %s, path: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

// For the sake of learning purposes if we want multiple middleware
// func requireAuthMiddleware(next http.Handler) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// check if the user is authenticated
// 		var token string = r.Header.Get("Authorization")
// 		if token != "Bearer token" {
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	}
// }
