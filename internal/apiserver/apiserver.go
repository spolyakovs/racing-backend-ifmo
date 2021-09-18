package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spolyakovs/racing-backend-ifmo/internal/store"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (server *APIServer) Start(config *Config) error {
	server.logger.Info("Starting server")

	if err := server.configureLogger(); err != nil {
		return nil
	}

	server.configureRouter()

	if err := server.configureStore(); err != nil {
		return err
	}

	server.logger.Info("Server started")

	return http.ListenAndServe(server.config.BindAddr, server.router)
}

func (server *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(server.config.LogLevel)
	if err != nil {
		return err
	}

	server.logger.SetLevel(level)
	return nil
}

func (server *APIServer) configureRouter() {
	server.router.HandleFunc("/hello", server.handleHello())
}

func (server *APIServer) configureStore() error {
	st := store.New(server.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	server.store = st
	return nil
}

func (server *APIServer) handleHello() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "Hello")
	}
}
