package api

import (
	"data-collector/api/handler"
	"data-collector/service/history"
	"github.com/GiairoZeppeli/utils/context"
	"github.com/GiairoZeppeli/utils/middleware"
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

func NewServer(ctx context.MyContext) *Server {
	router := mux.NewRouter()

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	wrappedRouter := middleware.RecoveryMiddleware(ctx, router)

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

func (s *Server) Shutdown(ctx context.MyContext) error {
	return s.httpServer.Shutdown(ctx.Ctx)
}

func (s *Server) HandlePing(ctx context.MyContext) {
	s.router.HandleFunc("/ping/", handler.Ping(ctx)).Methods(http.MethodGet)
}

func (s *Server) HandleAccount(ctx context.MyContext) {
	s.router.HandleFunc("/account/balance", handler.AccountBalance(ctx)).Methods(http.MethodGet)
}

func (s *Server) HandleHistory(ctx context.MyContext, service history.Service) {
	s.router.HandleFunc("/history/candle", handler.CandleHistory(ctx, service)).Methods(http.MethodGet)
}
