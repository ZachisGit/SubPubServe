export IPFS_SWARM_KEY="/key/swarm/psk/1.0.0/\n/base16/\n32fe3e79d4b4ace82e07d3217678a5a19b918da5208ae071dd0de89a65680905\n"
export IPFS_PATH=.ipfs

mkdir $IPFS_PATH

#echo $IPFS_SWAMR_KEY > "$(echo $IPFS_PATH)/swarm.key"

./ipfs daemon
