package main

import (
    "fmt"
    "log"
    "os/exec"
    "os"
    "io/ioutil"
    "strconv"
)

func ipfs_init(path string) bool {

    swarm_key := []byte("/key/swarm/psk/1.0.0/\n/base16/\n32fe3e79d4b4ace82e07d3217678a5a19b918da5208ae071dd0de89a65680905")
    err := ioutil.WriteFile(path+"/.ipfs/swarm.key", swarm_key, 0644)

    cmd := exec.Command(path+"/ipfs","init")
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

    cmd := exec.Command(path+"/ipfs","bootstrap", "rm", "--all")
    cmd.Env = os.Environ()
    cmd.Env = append(cmd.Env, "IPFS_PATH=.ipfs")
    out,err := cmd.Output()
    _ = out
    if err != nil {
        log.Fatal(err)
        return false
    }

    cmd = exec.Command(path+"/ipfs","bootstrap", "add", "/ip4/23.88.38.110/tcp/4001/p2p/12D3KooWA5Fz5qHMj3scoCktbPDcK4SPgzqACJhSwFqjgkmnvGkc")
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

    cmd := exec.Command(path+"/ipfs","daemon")
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

func ipfs_config_ports(path string, swarm int, http int, gateway int) bool {

    s := fmt.Sprintf("'[\"/ip4/0.0.0.0/tcp/%d\",\"/ip6/::/tcp/%d\",\"/ip4/0.0.0.0/udp/%d/quic\",\"/ip6/::/udp/%d/quic\"]'",swarm,swarm,swarm,swarm)

    cmd := exec.Command("sh","-c",path+"/ipfs config --json Addresses.Swarm "+s)
    cmd.Env = os.Environ()
    cmd.Env = append(cmd.Env, "IPFS_PATH=.ipfs")
    out,err := cmd.Output()
    _ = out
    if err != nil {
        log.Fatal(err)
        return false
    }


    cmd = exec.Command(path+"/ipfs","config","Addresses.API",fmt.Sprintf("/ip4/0.0.0.0/tcp/%d",http))
    cmd.Env = os.Environ()
    cmd.Env = append(cmd.Env, "IPFS_PATH=.ipfs")
    out,err = cmd.Output()
    _ = out
    if err != nil {
        log.Fatal(err)
        return false
    }


    cmd = exec.Command(path+"/ipfs","config","Addresses.Gateway",fmt.Sprintf("/ip4/0.0.0.0/tcp/%d",gateway))
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


func main() {

    args := os.Args[1:]

    port_swarm := 4001
    port_http := 5001
    port_gateway := 8080

    if (len(args) == 3) {
        port_swarm,_ = strconv.Atoi(args[0])
        port_http,_ = strconv.Atoi(args[1])
        port_gateway,_ = strconv.Atoi(args[2])
    }

    path, err := os.Getwd()

    // Check if ipfs already initialized
    if _, err = os.Stat(path+"/.ipfs"); os.IsNotExist(err) {
        err = os.Mkdir(path + "/.ipfs", os.ModePerm);
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

    ipfs_daemon(path)
}
