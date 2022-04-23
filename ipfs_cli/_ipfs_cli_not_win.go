package main

import (
    "fmt"
    "log"
    "os/exec"
    "os"
    "io/ioutil"
    "strconv"
    "runtime"
    "path/filepath"
)

import (
    "syscall"
)

func ipfs_init(path string) bool {

    swarm_key := []byte("/key/swarm/psk/1.0.0/\n/base16/\n32fe3e79d4b4ace82e07d3217678a5a19b918da5208ae071dd0de89a65680905")
    err := ioutil.WriteFile(path+_SEPERATOR+".ipfs" + _SEPERATOR + "swarm.key", swarm_key, 0644)

    cmd := exec.Command(_IPFS_BINARY_LOCATION,"init")
    cmd.Env = os.Environ()
    cmd.Env = append(cmd.Env, "IPFS_PATH=.ipfs")
    out,err := cmd.Output()
    _ = out
    if err != nil {
        log.Fatal(err)
        return false
    }
    return true
}


func ipfs_bootstrap(path string) bool {

    cmd := exec.Command(_IPFS_BINARY_LOCATION,"bootstrap", "rm", "--all")
    cmd.Env = os.Environ()
    cmd.Env = append(cmd.Env, "IPFS_PATH=.ipfs")
    out,err := cmd.Output()
    _ = out
    if err != nil {
        log.Fatal(err)
        return false
    }

    cmd = exec.Command(_IPFS_BINARY_LOCATION,"bootstrap", "add", "/ip4/23.88.38.110/tcp/4001/p2p/12D3KooWA5Fz5qHMj3scoCktbPDcK4SPgzqACJhSwFqjgkmnvGkc")
    cmd.Env = os.Environ()
    cmd.Env = append(cmd.Env, "IPFS_PATH=.ipfs")
    out,err = cmd.Output()
    _ = out
    if err != nil {
        log.Fatal(err)
        return false
    }

    return true
}


func ipfs_daemon(path string) bool {

    cmd := exec.Command(_IPFS_BINARY_LOCATION,"daemon")
    cmd.Env = os.Environ()
    cmd.Env = append(cmd.Env, "IPFS_PATH=.ipfs")

    cmd.SysProcAttr = &syscall.SysProcAttr{Pdeathsig: syscall.SIGTERM}

    out,err := cmd.Output()
    _ = out
    if err != nil {
        log.Fatal(err)
        return false
    }
    return true
}

func ipfs_config_ports(path string, swarm int, http int, gateway int) bool {

    s := fmt.Sprintf("'[\"/ip4/0.0.0.0/tcp/%d\",\"/ip6/::/tcp/%d\",\"/ip4/0.0.0.0/udp/%d/quic\",\"/ip6/::/udp/%d/quic\"]'",swarm,swarm,swarm,swarm)

    cmd := exec.Command("sh","-c",_IPFS_BINARY_LOCATION + " config --json Addresses.Swarm "+s)
    cmd.Env = os.Environ()
    cmd.Env = append(cmd.Env, "IPFS_PATH=.ipfs")
    out,err := cmd.Output()
    _ = out
    if err != nil {
        log.Fatal(err)
        return false
    }


    cmd = exec.Command(_IPFS_BINARY_LOCATION,"config","Addresses.API",fmt.Sprintf("/ip4/0.0.0.0/tcp/%d",http))
    cmd.Env = os.Environ()
    cmd.Env = append(cmd.Env, "IPFS_PATH=.ipfs")
    out,err = cmd.Output()
    _ = out
    if err != nil {
        log.Fatal(err)
        return false
    }


    cmd = exec.Command(_IPFS_BINARY_LOCATION,"config","Addresses.Gateway",fmt.Sprintf("/ip4/0.0.0.0/tcp/%d",gateway))
    cmd.Env = os.Environ()
    cmd.Env = append(cmd.Env, "IPFS_PATH=.ipfs")
    out,err = cmd.Output()
    _ = out
    if err != nil {
        log.Fatal(err)
        return false
    }


    return true
}

func ipfs_experimental(path string) bool {

        cmd := exec.Command(_IPFS_BINARY_LOCATION,"config","--json","Pubsub.Enabled","true")
    cmd.Env = os.Environ()
    cmd.Env = append(cmd.Env, "IPFS_PATH=.ipfs")
    out,err := cmd.Output()
    _ = out
    if err != nil {
        log.Fatal(err)
        return false
    }


    cmd = exec.Command(_IPFS_BINARY_LOCATION,"config","--json","Experimental.Libp2pStreamMounting","true")
    cmd.Env = os.Environ()
    cmd.Env = append(cmd.Env, "IPFS_PATH=.ipfs")
    out,err = cmd.Output()
    _ = out
    if err != nil {
        log.Fatal(err)
        return false
    }


    return true
}


var _SEPERATOR = "/"
var _BINARY_NAME = "ipfs"
var _IPFS_BINARY_LOCATION = ""

func main() {
    if runtime.GOOS == "windows" {
        _SEPERATOR = "\\"
        _BINARY_NAME = "ipfs.exe"
    }

    args := os.Args[1:]

    port_swarm := 4001
    port_http := 5001
    port_gateway := 8080
    _IPFS_BINARY_LOCATION,_ = filepath.Abs("." + _SEPERATOR + _BINARY_NAME)

    if (len(args) == 4) {
        _IPFS_BINARY_LOCATION,_ = filepath.Abs(args[0])
        port_swarm,_ = strconv.Atoi(args[1])
        port_http,_ = strconv.Atoi(args[2])
        port_gateway,_ = strconv.Atoi(args[3])
    }

    path, err := os.Getwd()

    // Check if ipfs already initialized
    if _, err = os.Stat(path + _SEPERATOR + ".ipfs"); os.IsNotExist(err) {
        err = os.Mkdir(path + _SEPERATOR + ".ipfs", os.ModePerm);
        if (ipfs_init(path)) {
            fmt.Println("Init successfull")
            if (ipfs_bootstrap(path)) {
                fmt.Println("Bootstrap successfull")
            } else {
                fmt.Println("Bootstrap failed")
                os.Exit(1)
            }
        } else {
            fmt.Println("Init failed")
            os.Exit(1)
        }
    }
    if (ipfs_config_ports(path,port_swarm,port_http,port_gateway)) {
        fmt.Println("Port config Successfull")
    } else {
        fmt.Println("Port config failed")
        os.Exit(1)
    }

    if (ipfs_experimental (path)) {
        fmt.Println("Enabeling experimental features Successfull")
    } else {
            fmt.Println("Enabeling experimental features failed")
        os.Exit(1)
    }

    ipfs_daemon(path)
}
