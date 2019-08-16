package modules

import (
	"fmt"
	"log"
	"net"
)

//LocalInfo structure
type LocalInfo struct {
	IPs []string
}

// Run is the interface of cron doing tasks.
func (l *LocalInfo) Run() {
	var err error
	l.IPs, err = l.GetIP()
	if err != nil {
		log.Println(" Failed to get IP.")
	}
	log.Println(l.IPs)
}

//GetIP is used to get the IP list of  local computer.
func (l *LocalInfo) GetIP() (ips []string, err error) {
	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("fail to get net interface addrs: %v", err)
		return ips, err
	}

	for _, address := range interfaceAddr {
		ipNet, isValidIPNet := address.(*net.IPNet)
		if isValidIPNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	return ips, err
}
