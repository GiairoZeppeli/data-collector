package api

import (
	"context"
	"data-collector/api/handler"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/swaggo/http-swagger"
	"net/http"
	"time"
)

const (
	maxHeaderBytes = 1 << 20 // 1 MB
	readTimeout    = 10 * time.Second
	writeTimeout   = 10 * time.Second
)

type Server struct {
	httpServer *http.Server
	router     *mux.Router
}

func NewServer(ctx utils.MyContext) *Server {
	router := mux.NewRouter()

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	wrappedRouter := utils.RecoveryMiddleware(ctx, router)

	return &Server{
		httpServer: &http.Server{
			Addr:           viper.GetString("db"),
			MaxHeaderBytes: maxHeaderBytes,
			ReadTimeout:    readTimeout,
			WriteTimeout:   writeTimeout,
			Handler:        wrappedRouter,
		},
		router: router,
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) HandlePing(ctx utils.MyContext) {
	s.router.HandleFunc("/ping/", handler.Ping(ctx)).Methods(http.MethodGet)
}
