# dns-proxy-poc

## How to use

1. Build and run a consul dev server from the [`ishustava/dns-proxy-poc` branch](https://github.com/hashicorp/consul/tree/ishustava/dns-proxy-poc):
    ```bash
    make dev
    ./bin/consul agent -server -dev
    ```
2. Register a service
   ```bash
   ./bin/consul services register -name foo -address 1.1.1.1 -port 8000
   ```
3. Run this proxy in another terminal:
   ```bash
   go run main.go
   ```
4. Try to resolve DNS:
   ```
   dig @localhost -p 8053 foo.service.consul
   ```
