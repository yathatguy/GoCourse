package main

import (
	"errors"
	"fmt"
	"log"
	"net"
)

func main() {
	ips := mirroredQuery()
	fmt.Printf("The nearest Ubuntu mirrow for you in Russian Federartion has the following IP(s): %v", ips)
}

func mirroredQuery() []string {
	responses := make(chan []string, 3)

	for _, fqdn := range []string{"mirrors.powernet.com.ru", "mirror.yandex.ru", "mirror.timeweb.ru", "mirror.corbina.net", "mirror.beget.ru", "mirror.truenetwork.ru"} {
		go func(fqdn string) {
			ips, err := request(fqdn)
			if err != nil {
				log.Printf("can't get IP address for %s", fqdn)
			}
			responses <- ips
		}(fqdn)
	}
	return <-responses
}

func request(hostname string) ([]string, error) {
	ips, err := net.LookupIP(hostname)
	if err != nil {
		return nil, errors.New("can't get IP address")
	}
	var ipSlice []string
	for _, ip := range ips {
		ipSlice = append(ipSlice, ip.String())
	}
	return ipSlice, nil
}
