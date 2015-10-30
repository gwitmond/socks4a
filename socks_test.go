// Socks4a Client library
//
// Connects to Socks servers such as the Tor Socks port
//
// Specs: http://socks-relay.sourceforge.net/socks4.protocol.txt
//        http://socks-relay.sourceforge.net/socks4a.protocol.txt
//
// In violation of the specs, we don't send the username as identd servers have died out
// And we don't expect the Tor client to bother with the local username.
//
// Copyright 2015, Guido Witmond <guido@witmond.nl>
// Licensed under GPL v3 or later. See LICENSE

package socks4a

import (
	"testing"
	"io"
	"log"
	"os"
)

func TestSocksConnection(t *testing.T) {
	torsocks := "127.0.0.1:9050"
	destination := "eccentric-authentication.org:80"

	s := &Socks4a {
		Network: "tcp",
		Address: torsocks,
	}

	conn, err := s.Dial(destination)
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: eccentric-authentication.org\r\n\r\n"))
	n, err := io.Copy(os.Stdout, conn)
	if err != nil {
		panic(err)
	}
	log.Printf("copying %v bytes with error %v", n, err)
	conn.Close()
}
