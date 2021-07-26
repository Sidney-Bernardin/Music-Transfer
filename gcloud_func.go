package musictransfer

import (
	"fmt"
	"net/http"
)

func GCloudFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi")
}
