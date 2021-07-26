package musictransfer

import (
	"flag"
	"log"
	"net/http"
	"testing"
)

var port = flag.String("p", "8080", "Port for the test server to listen and serve on.")

func TestGCloudFunc(t *testing.T) {
	http.HandleFunc("/", GCloudFunc)
	log.Println("Listening on :" + *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
