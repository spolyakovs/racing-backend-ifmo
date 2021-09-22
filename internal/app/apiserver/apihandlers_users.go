package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/spolyakovs/racing-backend-ifmo/internal/app/model"
)

func (server *server) handleRegistration() http.HandlerFunc {
	type request struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(writer http.ResponseWriter, req *http.Request) {
		requestStruct := &request{}
		if err := json.NewDecoder(req.Body).Decode(requestStruct); err != nil {
			server.error(writer, req, http.StatusBadRequest, err)
			return
		}

		user := &model.User{
			Username: requestStruct.Username,
			Email:    requestStruct.Email,
			Password: requestStruct.Password,
		}
		if err := server.store.User().Create(user); err != nil {
			server.error(writer, req, http.StatusUnprocessableEntity, err)
			return
		}
		server.respond(writer, req, http.StatusCreated, user)
	}
}

func (server *server) handleLogin() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(writer http.ResponseWriter, req *http.Request) {
		requestStruct := &request{}
		if err := json.NewDecoder(req.Body).Decode(requestStruct); err != nil {
			server.error(writer, req, http.StatusBadRequest, err)
			return
		}

		user, err := server.store.User().FindByEmail(requestStruct.Email)
		if err != nil || !user.ComparePassword(requestStruct.Password) {
			server.error(writer, req, http.StatusUnauthorized, errIncorrectEmailOrPassword)
			return
		}

		session, err := server.sessionStore.Get(req, sessionName)
		if err != nil {
			server.error(writer, req, http.StatusInternalServerError, err)
			return
		}

		session.Values["user_id"] = user.ID
		if err := server.sessionStore.Save(req, writer, session); err != nil {
			server.error(writer, req, http.StatusInternalServerError, err)
			return
		}

		server.respond(writer, req, http.StatusOK, nil)
	}
}

func (server *server) handleUsersMe() http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		server.respond(writer, req, http.StatusOK, req.Context().Value(ctxKeyUser).(*model.User))
	}
}
