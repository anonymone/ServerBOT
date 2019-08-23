package modules

import (
	"fmt"
	"log"
	"net"
)

//LocalInfo structure
type LocalInfo struct {
	IPs     []string
	changed bool
}

// Run is the interface of cron doing tasks.
func (l *LocalInfo) Run() {
	// var err error
	newIPs, err := l.GetIP()
	if err != nil {
		log.Println(" Failed to get IP.")
	}
	changed := l.isChanged(newIPs)
	if l.changed {
		l.IPs = newIPs
		l.changed = changed
	}
	log.Println(l.IPs)
}

// isChanged is used to check whether the local data is changed.
func (l *LocalInfo) isChanged(newData []string) bool {
	var find bool
	for _, newS := range newData {
		find = false
		for _, s := range l.IPs {
			if newS == s {
				find = true
			}
		}
		if !find {
			return true
		}
	}
	return false
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
