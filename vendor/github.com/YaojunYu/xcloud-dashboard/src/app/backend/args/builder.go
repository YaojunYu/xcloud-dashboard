package args

import "net"

var builder = &holderBuilder{holder: Holder}

// Used to build argument holder structure. It is private to make sure that only 1 instance can be created
// that modifies singleton instance of argument holder.
type holderBuilder struct {
	holder *holder
}


// SetInsecurePort 'insecure-port' argument of Dashboard binary.
func (self *holderBuilder) SetInsecurePort(port int) *holderBuilder {
	self.holder.insecurePort = port
	return self
}

// SetPort 'port' argument of Dashboard binary.
func (self *holderBuilder) SetPort(port int) *holderBuilder {
	self.holder.port = port
	return self
}

// SetInsecureBindAddress 'insecure-bind-address' argument of Dashboard binary.
func (self *holderBuilder) SetInsecureBindAddress(ip net.IP) *holderBuilder {
	self.holder.insecureBindAddress = ip
	return self
}

// SetBindAddress 'bind-address' argument of Dashboard binary.
func (self *holderBuilder) SetBindAddress(ip net.IP) *holderBuilder {
	self.holder.bindAddress = ip
	return self
}

// GetHolderBuilder returns singletone instance of argument holder builder.
func GetHolderBuilder() *holderBuilder {
	return builder
}
