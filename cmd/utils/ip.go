package utils

import (
	"errors"
	"net"
	"strconv"
	"strings"
)

type Interface struct {
	Ip     string
	Subnet int
}

func LocalIp() (Interface, error) {
	ifaces, err := net.Interfaces()

	if err != nil {
		return Interface{"", 0}, err
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 || iface.Name == "virbr0" {
			continue // interface down & loopback interface
		}

		addr, err := iface.Addrs()

		if err != nil {
			return Interface{"", 0}, err
		}

		info := strings.Split(addr[0].String(), "/")

		sub, err := strconv.Atoi(info[1])

		if err != nil {
			return Interface{"", 0}, err
		}

		return Interface{info[0], sub}, nil
	}
	return Interface{"", 0}, errors.New("no network interface found matching requirements")
}
