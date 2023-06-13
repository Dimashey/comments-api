package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router  *mux.Router
	Service CommentService
	Server  *http.Server
}

func NewHandler(service CommentService) *Handler {
	handler := &Handler{
		Service: service,
	}

	handler.Router = mux.NewRouter()
	handler.mapRoutes()
	handler.Router.Use(JSONMiddleware)
	handler.Router.Use(LoggingMiddleware)

	handler.Server = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: handler.Router,
	}

	return handler
}

func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/api/v1/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/v1/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/v1/comment/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/v1/comment/{id}", h.DeleteComment).Methods("DELETE")
}

func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	chanel := make(chan os.Signal, 1)

	signal.Notify(chanel, os.Interrupt)

	<-chanel

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	h.Server.Shutdown(ctx)

	log.Println("shot down gracefully")
	return nil
}
