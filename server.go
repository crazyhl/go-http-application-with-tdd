package go_http_application_with_tdd

import (
	"fmt"
	"net/http"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "20")
}
