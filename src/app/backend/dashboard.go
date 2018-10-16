package main

import (
	"log"
	"os"
	"github.com/spf13/pflag"
	"net"
	"net/http"
	"fmt"
	"github.com/YaojunYu/xcloud-dashboard/src/app/backend/handler"
)

// args var
var (
	argInsecurePort        = pflag.Int("insecure-port", 9090, "The port to listen to for incoming HTTP requests.")
	argPort                = pflag.Int("port", 8443, "The secure port to listen to for incoming HTTPS requests.")
	argInsecureBindAddress = pflag.IP("insecure-bind-address", net.IPv4(127, 0, 0, 1), "The IP address on which to serve the --port (set to 0.0.0.0 for all interfaces).")
	argBindAddress         = pflag.IP("bind-address", net.IPv4(0, 0, 0, 0), "The IP address on which to serve the --secure-port (set to 0.0.0.0 for all interfaces).")

)

func main() {
	log.SetOutput(os.Stdout)
	pflag.Parse()

	apiHandler, err := handler.CreateHTTPAPIHandler()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/api/", apiHandler)

	addr := fmt.Sprintf("%s:%d", *argInsecureBindAddress, *argPort)
	go func() { log.Fatal(http.ListenAndServe(addr, nil)) }()

	// block main thread
	select {}
}


