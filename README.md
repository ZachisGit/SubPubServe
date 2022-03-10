# SubPubServe
Build a decentralised server infrastructure with a sub/pub queue based API instead of a REST based API server.
The main difference between a centralised REST api and the sub/pub queue based api is that your server applications can be completely stateless while the client makes sure it gets the processed information it wants from the Peer-to-Peer network.

The goal is to get rid of the centralised control server (like the kubernetes api server), solve the storage allocation problem natively without having to setup volume mounts or configuring minio tenants and have the system be completely decentralised.

## Design Policy
All software built on top of the SubPubServe stack will include the SubPubServe library to interact with other nodes running microservices.
Internally you can interact with the library in a similar way to a normal rest endpoint.
All stored data will be open and accessible to the entire network, security and access control is natively handled purely through encryption.
Node and data Identity is guaranteed, everybody in the network will know exactly which piece of data was written by which client or node.


## Features

Some if the features that you will get running right out of the box:
* IPFS based databases for long term redundant storage
* Automatic node discovery
* Automatic load balancing between all your nodes
* Edge devices can act as extra processing/storage nodes in the existing network
* Automatic routing
* Fully decentralised

# Sub/Pub Queue API

Decentralized Sub/Pub Queue based infrastructure setup

![sub_pub_queue](/imgs/subpubqueue.svg)


REST based server infrastructure in a Kubernetes Cluster

![rest based api](/imgs/kbc_rest.svg)


# Interesting Links
https://skyhub.site/Skynet/skfs/raw/branch/master/docs/experimental-features.md


