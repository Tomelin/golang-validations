package validations

import (
	"fmt"
	"net"
)

func IsDomainValid(domain string) bool {

	ns, err := net.LookupNS(domain)

	if err != nil || len(ns) == 0 {
		fmt.Println("NS")
		fmt.Println(ns)
		return false
	}

	return true
}

func IsDomainAddrValid(domain string) bool {

	if !IsDomainValid(domain) {
		return false
	}

	a, err := net.LookupAddr(domain)
	if err != nil || len(a) == 0 {
		fmt.Println("A")
		fmt.Println(a)
		return false
	}
}
