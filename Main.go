package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	r := NewRootRouter()

	// Start Router
	var err error
	addr := fmt.Sprintf(":%v", PORT)
	log.Info("Listening on ", addr)
	if DEV {
		log.Warn("Server running in development mode. Ensure that tls is setup before being used in production.")
		err = http.ListenAndServe(addr, r)
	} else {
		tlsCert := os.Getenv("TLS_CERT")
		tlsKey := os.Getenv("TLS_KEY")

		if len(tlsCert) == 0 || len(tlsKey) == 0 {
			log.Fatal("Missing TLS_CERT or TLS_KEY")
		}

		err = http.ListenAndServeTLS(addr, tlsCert, tlsKey, r)
	}
	if err != nil {
		log.Error("Error listening at ", addr, ". ", err)
	}
}
