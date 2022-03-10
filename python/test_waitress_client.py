import ipfsclipy
import requests

ipfsclipy.proxy.P2P_SERVER_ID = input ("P2P_SERVER_ID:")

ipfsclipy.proxy.tcpclient("falcon-server-0","http://127.0.0.1:12345")

while True:
    resp = requests.get("http://127.0.0.1:12345/simple")
    print (resp.status_code,resp.text)
    input("")

