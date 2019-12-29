package handle

import (
	"fmt"
	"net/http"
)

// HomeLink shows welcome message
func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
