package main

import (
	"log"
	"os"
	"github.com/spf13/pflag"
	"net"
	"net/http"
	"fmt"
	"github.com/YaojunYu/xcloud-dashboard/src/app/backend/handler"
	"github.com/YaojunYu/xcloud-dashboard/src/app/backend/args"
	"github.com/YaojunYu/xcloud-dashboard/src/app/backend/client"
)

// args var
var (
	argInsecurePort        = pflag.Int("insecure-port", 9090, "The port to listen to for incoming HTTP requests.")
	argPort                = pflag.Int("port", 8443, "The secure port to listen to for incoming HTTPS requests.")
	argInsecureBindAddress = pflag.IP("insecure-bind-address", net.IPv4(127, 0, 0, 1), "The IP address on which to serve the --port (set to 0.0.0.0 for all interfaces).")
	argBindAddress         = pflag.IP("bind-address", net.IPv4(0, 0, 0, 0), "The IP address on which to serve the --secure-port (set to 0.0.0.0 for all interfaces).")

	argApiserverHost       = pflag.String("apiserver-host", "", "The address of the Kubernetes Apiserver "+
		"to connect to in the format of protocol://address:port, e.g., "+
		"http://localhost:8080. If not specified, the assumption is that the binary runs inside a "+
		"Kubernetes cluster and local discovery is attempted.")
	argKubeConfigFile     = pflag.String("kubeconfig", "", "Path to kubeconfig file with authorization and master location information.")

)

func main() {
	// Set logging output to standard console out
	log.SetOutput(os.Stdout)

	pflag.Parse()

	// Initializes dashboard arguments holder so we can read them in other packages
	initArgHolder()

	// Initial request to the apiserver
	clientManager := client.NewClientManager(args.Holder.GetKubeConfigFile(), args.Holder.GetApiServerHost())
	versionInfo, err := clientManager.InsecureClient().Discovery().ServerVersion()
	if err != nil {
		handleFatalInitError(err)
	}
	log.Printf("Successful initial request to the apiserver, version: %s", versionInfo.String())


	apiHandler, err := handler.CreateHTTPAPIHandler()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/api/", apiHandler)

	addr := fmt.Sprintf("%s:%d", *argInsecureBindAddress, *argInsecurePort)
	log.Printf("Serving insecurely on HTTP server url: %s", addr)
	go func() { log.Fatal(http.ListenAndServe(addr, nil)) }()

	// block main thread
	select {}
}

func initArgHolder() {
	builder := args.GetHolderBuilder()
	builder.SetInsecurePort(*argInsecurePort)
	builder.SetPort(*argPort)
	builder.SetInsecureBindAddress(*argInsecureBindAddress)
	builder.SetBindAddress(*argBindAddress)
}

/**
 * Handles fatal init error that prevents server from doing any work. Prints verbose error
 * message and quits the server.
 */
func handleFatalInitError(err error) {
	log.Fatalf("Error while initializing connection to Kubernetes apiserver. "+
		"This most likely means that the cluster is misconfigured (e.g., it has "+
		"invalid apiserver certificates or service account's configuration) or the "+
		"--apiserver-host param points to a server that does not exist. Reason: %s\n"+
		"Refer to our FAQ and wiki pages for more information: "+
		"https://github.com/kubernetes/dashboard/wiki/FAQ", err)
}


