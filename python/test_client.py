import ipfsclipy
import requests
ipfsclipy.await_init()

input("READY?")
port = ipfsclipy.__main__._random_port()
print (ipfsclipy.http_api.p2p.forward("test-proto2", "/ip4/0.0.0.0/tcp/"+str(port)+"/", "/p2p/12D3KooWRqGv8zQ29TFPw4QvLxPdurUbYxFcaMYrduvG5dMTjnRq", True))

while True:
    input("")
    resp = requests.get("http://127.0.0.1:"+str(port)+"/simple")
    print (resp.status_code,resp.text)
