import ipfsclipy

ipfsclipy.await_init()

input("123")
print (ipfsclipy.http_api.p2p.listener("test-proto2", "/ip4/127.0.0.1/tcp/8081/", True))

while True:
    print(ipfsclipy.swarm_port(), ipfsclipy.http_port(),ipfsclipy.gateway_port())
    input("")
