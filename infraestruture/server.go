package infraestruture

import (
	
	"net/http"
	"time"
	"github.com/gorilla/mux"
	
)



func Service() {

    r := mux.NewRoute()

	srv := &http.Server{
		Addr: "0.0.0.0:8000",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}
	fmt.Println("server running on http://0.0.0.0:8000 ")
	
}