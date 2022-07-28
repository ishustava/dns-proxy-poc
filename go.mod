module dns-proxy-poc

go 1.18

require (
	github.com/hashicorp/consul v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.37.1
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/golang/protobuf v1.5.0 // indirect
	golang.org/x/net v0.0.0-20211216030914-fe4d6282115f // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/genproto v0.0.0-20200623002339-fbb79eadd5eb // indirect
)

replace github.com/hashicorp/consul => ../consul/
