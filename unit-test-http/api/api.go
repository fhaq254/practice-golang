package api

import (
	"fmt"
	"net/http"
)

func CheckHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "health check - OK")
}
