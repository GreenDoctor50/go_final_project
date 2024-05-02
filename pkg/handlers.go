package hadlers

import (
	"fmt"
	"net/http"
)

func getIndexHTML(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}