package main

import (
	"context"
	"fmt"
	"net"

	pbconsuldns "dns-proxy-poc/consul-proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Based on https://github.com/mkeeler/dns-proxy/blob/main/cmd/proxy-dns/main.go.
func main() {
	addr := net.UDPAddr{
		Port: 8053,
		IP:   net.ParseIP("127.0.0.1"),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	grpcConn, err := grpc.Dial("localhost:8502", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	proxyDNS(conn, pbconsuldns.NewDNSServiceClient(grpcConn))
}

func proxyDNS(conn *net.UDPConn, client pbconsuldns.DNSServiceClient) {
	buf := make([]byte, 1520)

	for {
		_, addr, err := conn.ReadFrom(buf)
		if err != nil {
			fmt.Printf("error reading from conn: %v\n", err)
			continue
		}

		req := &pbconsuldns.QueryRequest{
			Msg: buf,
		}

		resp, err := client.Query(context.Background(), req)
		if err != nil {
			fmt.Printf("error resolving request: %v\n", err)
			continue
		}

		_, err = conn.WriteTo(resp.Msg, addr)
		if err != nil {
			fmt.Printf("error sending response: %v\n", err)
		}
	}
}
