import requests

from .__main__ import http_port


class p2p:

    @staticmethod
    def listener(protocol,target_address,allow_custom_protocol):
        endpoint = "http://127.0.0.1:"+str(http_port())+"/api/v0/p2p/listen?arg="+protocol+"&arg="+target_address+"&allow-custom-protocol="+("true" if allow_custom_protocol else "false")
        resp = requests.post(endpoint)

        if resp.status_code != 200:
            #print ("None:","p2p.listener")
            #print (resp.text)
            return None
        #print("succ")
        return resp.text


    @staticmethod
    def ls():
        endpoint = "http://127.0.0.1:"+str(http_port())+"/api/v0/p2p/ls"
        #print ("Pre-request")
        resp = requests.post(endpoint)
        #print ("123")
        if resp.status_code != 200:
            #print ("None:","p2p.listener")
            return None
   
        return resp.text


    def forward(protocol,listen_address,target_address,allow_custom_protocol):
        endpoint = "http://127.0.0.1:"+str(http_port())+"/api/v0/p2p/forward?arg="+protocol+"&arg="+listen_address+"&arg="+target_address+"&allow-custom-protocol="+("true" if allow_custom_protocol else "false")
        resp = requests.post(endpoint)

        if resp.status_code != 200:
            #print ("None:","p2p.listener")
            return None

        return resp.text



