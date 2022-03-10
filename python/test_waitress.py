import ipfsclipy
from waitress import serve
import falcon_server

ipfsclipy.proxy.tcpserver("falcon-server-0",8081)

serve(falcon_server.App, host="0.0.0.0", port=8081)
