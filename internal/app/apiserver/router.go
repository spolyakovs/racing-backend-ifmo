package apiserver

import "github.com/gorilla/handlers"

func (server *server) configureRouter() {
	server.router.Use(server.setRequestID)
	server.router.Use(server.logRequest)
	server.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	server.router.HandleFunc("/registration", server.handleRegistration()).Methods("POST")
	server.router.HandleFunc("/login", server.handleLogin()).Methods("POST")

	private := server.router.PathPrefix("/private").Subrouter()
	private.Use(server.authenticateUser)
	private.HandleFunc("/users/me", server.handleUsersMe())
	private.HandleFunc("/users/update", server.handleUsersMe())
	private.HandleFunc("/users/delete", server.handleUsersMe())

	private.HandleFunc("/teams/create", server.handleUsersMe())
	private.HandleFunc("/teams/update", server.handleUsersMe())
	private.HandleFunc("/teams/delete", server.handleUsersMe())
	private.HandleFunc("/teams/{id}", server.handleUsersMe())
	private.HandleFunc("/teams/{id}/drivers", server.handleUsersMe())
	private.HandleFunc("/teams/{id}/history", server.handleUsersMe())

	private.HandleFunc("/drivers/create", server.handleUsersMe())
	private.HandleFunc("/drivers/update", server.handleUsersMe())
	private.HandleFunc("/drivers/delete", server.handleUsersMe())
	private.HandleFunc("/drivers/{id}", server.handleUsersMe())
	private.HandleFunc("/drivers/{id}/team", server.handleUsersMe())
	private.HandleFunc("/drivers/{id}/career", server.handleUsersMe())

	private.HandleFunc("/races/create", server.handleUsersMe())
	private.HandleFunc("/races/update", server.handleUsersMe())
	private.HandleFunc("/races/delete", server.handleUsersMe())
	private.HandleFunc("/races/{id}", server.handleUsersMe())
	private.HandleFunc("/races/{id}/results", server.handleUsersMe())
}
