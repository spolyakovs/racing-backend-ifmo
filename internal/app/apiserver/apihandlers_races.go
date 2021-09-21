package apiserver

import "net/http"

// TODO: create, update, delete, createResults
func (server *server) handleRacesCreate() http.HandlerFunc {
	// type request struct {
	// 	Username string `json:"username"`
	// 	Email    string `json:"email"`
	// 	Password string `json:"password"`
	// }

	return func(writer http.ResponseWriter, req *http.Request) {
		server.respond(writer, req, http.StatusOK, nil)
	}
}
