package server

import (
	"fmt"
	"net/http"

	service "github.com/IanPedroV/dad-jokes/services"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	jokeId := service.GetJokeId()
	fmt.Fprintf(w, "%s", service.GetDadJokeImage(jokeId))
}
