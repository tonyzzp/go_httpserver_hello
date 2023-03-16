package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	var mux = &http.ServeMux{}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
	var ssl = os.Getenv("HTTP_SSL")
	var host = os.Getenv("HTTP_HOST")
	var email = os.Getenv("HTTP_EMAIL")
	log.Println(os.Environ())
	log.Println("ssl", ssl)
	log.Println("host", host)
	log.Println("email", email)
	var hosts = strings.Split(host, ",")
	log.Println("hosts", hosts)
	if strings.ToLower(ssl) == "true" && len(host) > 0 {
		var mgr = &autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			Cache:      autocert.DirCache("/cert-cache"),
			HostPolicy: autocert.HostWhitelist(hosts...),
			Email:      email,
		}
		var server = &http.Server{
			Addr:      ":https",
			Handler:   mux,
			TLSConfig: mgr.TLSConfig(),
		}
		var e = server.ListenAndServeTLS("", "")
		if e != nil {
			log.Fatalln(e)
		}
	} else {
		var server = &http.Server{
			Addr:    ":http",
			Handler: mux,
		}
		var e = server.ListenAndServe()
		if e != nil {
			log.Fatalln(e)
		}
	}
}
