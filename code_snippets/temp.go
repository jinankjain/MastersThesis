package main

import (
	"fmt"

	"github.com/scionproto/scion/go/lib/addr"
	"github.com/scionproto/scion/go/lib/sciond"
	"github.com/scionproto/scion/go/lib/snet"
)

var sciondPath = sciond.GetDefaultSCIONDPath(nil)

func server(addr *snet.Addr) {
	err := snet.Init(addr.IA, sciondPath, "")
	if err != nil {
		panic(err)
	}
	conn, err := snet.ListenSCION("udp4", addr)
	if err != nil {
		panic(err)
	}
	for {
		buf := make([]byte, 1024)
		len, raddr, _ := conn.ReadFromSCION(buf)
		conn.WriteToSCION([]byte("Hello from Server!"), raddr)
	}
}

func client(laddr *snet.Addr, raddr *snet.Addr) {
	err := snet.Init(laddr.IA, sciondPath, "")
	if err != nil {
		panic(err)
	}
	conn, err := snet.DialSCION("udp4", laddr, raddr)
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("Hello from client!"))
	buf := make([]byte, 1024)
	len, _ := conn.Read(buf)
	fmt.Printf("Server response: %s\n", string(buf))
	conn.Close()
}

func main() {
	server_addr, _ = snet.AddrFromString("1-ff00:0:133,[127.0.0.1]:40001")
	client_addr, _ = snet.AddrFromString("2-ff00:0:222,[127.0.0.1]:40002")
}
