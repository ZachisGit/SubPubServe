import ipfsclipy
import requests

ipfsclipy.proxy.P2P_SERVER_ID = input ("P2P_SERVER_ID:")

ipfsclipy.proxy.tcpclient("falcon-server-0","http://127.0.0.1:12345")

while True:
    try:
        #ipfsclipy.await_init()
        resp = requests.get("http://127.0.0.1:12345/simple")
        print (resp.status_code,resp.text)
    except:
        import traceback
        traceback.print_exc()

    input("")

