from subprocess import Popen, PIPE
import inspect
import os.path
import platform
import random
import threading

from . import port


_BINARY_FORMAT = ""
if platform.system() == "windows":
    _BINARY_FORMAT = ".exe"

_IPFS_CLI_PATH = os.path.dirname(os.path.abspath(inspect.getframeinfo(inspect.currentframe()).filename)) + os.path.sep + "resources" + os.path.sep
_PORT_SWARM = None
_PORT_HTTP = None
_PORT_GATEWAY = None

def _daemon_init():
    global _PORT_SWARM,_PORT_HTTP,_PORT_GATEWAY,_BINARY_FORMAT,_IPFS_CLI_PATH

    _PORT_SWARM = _random_port()
    _PORT_HTTP = _random_port()
    _PORT_GATEWAY = _random_port()
    print ([_IPFS_CLI_PATH + "ipfs-cli" + _BINARY_FORMAT,_IPFS_CLI_PATH+"ipfs"+_BINARY_FORMAT,str(_PORT_SWARM),str(_PORT_HTTP),str(_PORT_GATEWAY)])
    p = Popen([_IPFS_CLI_PATH + "ipfs-cli" + _BINARY_FORMAT,_IPFS_CLI_PATH+"ipfs"+_BINARY_FORMAT,str(_PORT_SWARM),str(_PORT_HTTP),str(_PORT_GATEWAY)], stdout=PIPE, bufsize=1, universal_newlines=True)
    p.communicate()

    _PORT_SWARM, _PORT_HTTP, _PORT_GATEWAY = None,None,None
    print ("[IPFS-CLI] ERROR: DAEMON FAILURE")

def swarm_port():
    global _PORT_SWARM
    return _PORT_SWARM

def http_port():
    global _PORT_HTTP
    return _PORT_HTTP

def gateway_port():
    global _PORT_GATEWAY
    return _PORT_GATEWAY

def _random_port():
    c_port = random.randint(4001,65535)
    while True:
        if not port.is_open(c_port):
            return c_port

def _th_start():
    th = threading.Thread(target=_daemon_init,args=())
    th.daemon = True
    th.start()

_th_start()
