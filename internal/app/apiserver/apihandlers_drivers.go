	package apiserver

import "net/http"

// TODO: create, update, delete, createContract, updateContract
func (server *server) handleDriversCreate() http.HandlerFunc {
	// type request struct {
	// 	Username string `json:"username"`
	// 	Email    string `json:"email"`
	// 	Password string `json:"password"`
	// }

	return func(writer http.ResponseWriter, req *http.Request) {
		server.respond(writer, req, http.StatusOK, nil)
	}
}
