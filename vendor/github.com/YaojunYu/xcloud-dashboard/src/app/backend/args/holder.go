package args

import "net"

var Holder = &holder{}

// Argument holder structure. It is private to make sure that only 1 instance can be created. It holds all
// arguments values passed to Dashboard binary.
type holder struct {
	insecurePort int
	port int

	insecureBindAddress net.IP
	bindAddress net.IP

	apiServerHost        string
	kubeConfigFile       string
}

func (self *holder) GetInsecurePort() int {
	return self.insecurePort
}

func (self *holder) GetPort() int {
	return self.port
}

func (self *holder) GetInsecureBindAddress() net.IP {
	return self.insecureBindAddress
}

func (self *holder) GetBindAddress() net.IP {
	return self.bindAddress
}

// GetApiServerHost 'apiserver-host' argument of Dashboard binary.
func (self *holder) GetApiServerHost() string {
	return self.apiServerHost
}

// GetKubeConfigFile 'kubeconfig' argument of Dashboard binary.
func (self *holder) GetKubeConfigFile() string {
	return self.kubeConfigFile
}