package main

import (
	"fmt"
	"net"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s: %s\n", os.Args[0], "<iface>")
}

func exitUsage(rt int, message string, args ...string) {
	if message != "" {
		fmt.Fprintln(os.Stderr, message, args)
	}

	usage()

	os.Exit(rt)
}

const (
	IPv4 = iota
	IPv6
)

func main() {
	if len(os.Args) != 2 {
		exitUsage(2, "")
	}

	name := os.Args[1]
	if name == "" {
		exitUsage(2, "Argument <iface> cannot be empty")
	}

	iface, err := net.InterfaceByName(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching interface %s\n%s\n", name, err.Error())
		os.Exit(1)
	}

	addrs, err := iface.Addrs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching address on interface %s\n%s\n", name, err.Error())
		os.Exit(1)
	}

	printAddrs(addrs, IPv4)
}

func printAddrs(addrs []net.Addr, flags uint) {
	for i := range addrs {
		switch addr := addrs[i].(type) {
			case *net.IPNet:
				if flags & IPv4 == IPv4 && assertIPv4(addr.IP) {
					fmt.Println(addr.String())
				} else if flags & IPv6 == IPv6 && assertIPv6(addr.IP) {
					fmt.Println(addr.String())
				}
			default:
				fmt.Fprintf(os.Stderr, "Unsupported type %T for address %s\n", addr, addr.String())
		}
	}
}

func assertIPv4(ip net.IP) (bool) {
	ok := ip.To4()
	if ok != nil {
		if len(ok) == net.IPv4len {
			return true
		}
	}

	return false
}

func assertIPv6(ip net.IP) (bool) {
	ok := ip.To16()
	if ok != nil {
		if len(ok) == net.IPv6len {
			return true
		}
	}

	return false
}
