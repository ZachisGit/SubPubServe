import socket
import re
from urllib.parse import urlparse
import sys
import _socket
import socket
import os

from .__main__ import await_init, random_port
from . import http_api


P2P_SERVER_ID = None


# Service name gets registered either based on domain and port 
# or ip and port. If no port is specified, all connections to 
# the domain/ip get forwarded to the service_name
#
# Example: tcpclient("myservice","test.example.com")
# Example: tcpclient("myservice","test.example.com:8080")
# Example: tcpclient("myservice","https://test.example:8080/url/path?some=parameters")
# Example: tcpclient("myservice","168.38.48.23:80")
# Example: tcpclient("myservice","http://168.38.48.23:80/url/path?some=parameters")
#
# Anti-Example: tcpclient("myservice","168.38.48.23")    # for ips port must be specified
def tcpclient(service_name,endpoint=None):
    if not "://" in endpoint:
        endpoint = "http://" + endpoint

    endpoint = urlparse(endpoint).netloc
    
    ip_split = re.findall(r"^(?:[0-9]{1,3}\.){3}[0-9]{1,3}", endpoint)
    
    # ip
    if len(ip_split) > 0:
        _add_service_name_ip(service_name, endpoint)
    # domain
    else:
        _add_service_name_domain(service_name,endpoint)
    return True   


# Service name gets registered based on listening port on the
# local server
def tcpserver(service_name, port):
    return _add_service_name_port(service_name,port)


# ((([a-z0-9]|[a-z0-9][a-z0-9\-]*[a-z0-9])\.)*([a-z0-9]|[a-z0-9][a-z0-9\-]*[a-z0-9]))?(:[0-9]+)





""" ServiceName - IP:Port Register """
_service_name_register = {}

def _add_service_name_domain(service_name,domain):
    global _service_name_register

    if not service_name in _service_name_register.keys():
        _service_name_register[service_name] = {"hosts":[], "ports":[]}

    _service_name_register[service_name]["hosts"].append(domain)

    return True
       



def _add_service_name_ip(service_name,ip):
    global _service_name_register

    if not service_name in _service_name_register.keys():
        _service_name_register[service_name] = {"hosts":[], "ports":[]}

    _service_name_register[service_name]["hosts"].append(ip)

    return True




def _add_service_name_port(service_name,port):
    global _service_name_register

    if not service_name in _service_name_register.keys():
        _service_name_register[service_name] = {"hosts":[], "ports":[]}

    _service_name_register[service_name]["ports"].append(port)

    return True


# Check if its a registered port
# if it is in the listener register,
# create a listener on the node
def _port_listen(port):

    service_name = None
    for k,v in _service_name_register.items():
        if port in v["ports"]:
            service_name = k
            break

    if service_name is None:
        return False

    await_init()

    print (http_api.p2p.listener(service_name, "/ip4/127.0.0.1/tcp/"+str(port)+"/", True))
    return True

# Check if host:port is in the host register,
# if so reroute the connection an create a forward
# endpoint on another port
def _host_forward(host,port):
    
    service_name = None
    for k,v in _service_name_register.items():
        if host + ":"+str(port) in v["hosts"]:
            service_name = k
            break

    if service_name is None:
        return host,port
    
    if len(_service_name_register[k]["ports"]) == 0:
        await_init()

        port = _random_port()
        print (http_api.p2p.forward(service_name, "/ip4/127.0.0.1/tcp/"+str(port)+"/", P2P_SERVER_ID, True))
        _service_name_register[k]["ports"].append(port)

    
    return "127.0.0.1",_service_name_register[k]["ports"][-1]


""" socket connection handler """
def getaddrinfo(host, port, family=0, type=0, proto=0, flags=0):
    """Resolve host and port into list of address info entries.
    Translate the host/port argument into a sequence of 5-tuples that contain
    all the necessary arguments for creating a socket connected to that service.
    host is a domain name, a string representation of an IPv4/v6 address or
    None. port is a string service name such as 'http', a numeric port number or
    None. By passing None as the value of host and port, you can pass NULL to
    the underlying C API.
    The family, type and proto arguments can be optionally specified in order to
    narrow the list of addresses returned. Passing zero as a value for each of
    these arguments selects the full range of results.
    """

    host,port = _host_forward(host, port)

    # We override this function since we want to translate the numeric family
    # and socket type values to enum constants.
    addrlist = []
    for res in _socket.getaddrinfo(host, port, family, type, proto, flags):
        af, socktype, proto, canonname, sa = res
        addrlist.append((socket._intenum_converter(af, socket.AddressFamily),
                         socket._intenum_converter(socktype, socket.SocketKind),
                         proto, canonname, sa))
    print (host,port,addrlist)
    return addrlist


def _bind_wrapper(self,a=None):
    #print ("Bind:",a)
    ret = self._bind(a)
    _port_listen(a[1])


_is_initialized = False

def _init(): 
    global _is_initialized

    if _is_initialized:
        return
    socket = sys.modules['socket']
    socket.getaddrinfo = getaddrinfo
    #socket.create_server = create_server
    socket.socket._bind = socket.socket.bind
    socket.socket.bind = _bind_wrapper

    sys.modules['socket'] = socket

_init()
