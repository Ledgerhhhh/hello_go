package util

import (
	"fmt"
	"net"
)

func GetIp() (ipv4, ipv6 string, err error) {
	interfaces, err := net.Interfaces()

	if err != nil {
		fmt.Println("Error:", err)
		return "", "", err
	}

	// 遍历每个网络接口
	for _, iface := range interfaces {
		// 排除一些特殊的接口，如 loopback 接口
		if iface.Flags&net.FlagUp != 0 && iface.Flags&net.FlagLoopback == 0 {
			addrs, err := iface.Addrs()
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			// 遍历每个地址
			for _, addr := range addrs {
				switch v := addr.(type) {
				case *net.IPNet:
					// 如果是 IPv4 地址
					if v.IP.To4() != nil {
						fmt.Printf("IPv4 Address: %s\n", v.IP)
						ipv4 = v.IP.String()
					}
				case *net.IPAddr:
					// 如果是 IPv6 地址
					if v.IP.To4() == nil {
						fmt.Printf("IPv6 Address: %s\n", v.IP)
						ipv6 = v.IP.String()
					}
				}
			}
		}
	}
	return ipv4, ipv6, nil
}
