package server

import (
	"fmt"
	"net/http"

	service "github.com/IanPedroV/dad-jokes/service"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	joke := service.Get()
	fmt.Fprintf(w, "%s", service.GetImage(joke.Id))
}
