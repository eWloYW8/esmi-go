package main

import (
	"fmt"

	esmi "github.com/eWloYW8/esmi-go"
)

func main() {
	cli, err := esmi.NewClient()
	if err != nil {
		fmt.Printf("create client failed: %v\n", err)
		return
	}
	defer cli.Close()

	sockets, err := cli.NumberOfSockets()
	if err != nil {
		fmt.Printf("NumberOfSockets failed: %v\n", err)
		return
	}

	proto, err := cli.HSMPProtoVersion()
	if err != nil {
		fmt.Printf("HSMPProtoVersion failed: %v\n", err)
		return
	}

	fmt.Printf("sockets=%d hsmp_proto=%d\n", sockets, proto)
}
