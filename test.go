package main

import (
    "fmt"
)
func main() {
    v := 394
    fmt.Println("./ipfs","config","--json","Addresses.Swarm",fmt.Sprintf("'[\"/ip4/0.0.0.0/tcp/%d\",\"/ip6/::/tcp/%d\",\"/ip4/0.0.0.0/udp/%d/quic\",\"/ip6/::/udp/%d/quic\"]'",v,v,v,v))
}
