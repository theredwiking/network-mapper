package utils

import (
	"errors"
	"net"
	"strings"
)

func LocalIp() (string, string, error) {
	ifaces, err := net.Interfaces()

	if err != nil {
		return "", "", err
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 || iface.Name == "virbr0" {
			continue // interface down & loopback interface
		}

		addr, err := iface.Addrs()

		if err != nil {
			return "", "", err
		}

		info := strings.Split(addr[0].String(), "/")

		return info[0], info[1], nil
	}
	return "", "", errors.New("are you even connected to a network?")
}
