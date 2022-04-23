mkdir -p build/tmp/
mkdir build/win
mkdir build/linux

wget "https://github.com/ipfs/go-ipfs/releases/download/v0.12.0/go-ipfs_v0.12.0_linux-386.tar.gz" -O build/tmp/linux.tar.gz
wget "https://github.com/ipfs/go-ipfs/releases/download/v0.12.0/go-ipfs_v0.12.0_windows-386.zip" -O build/tmp/win.zip

sudo apt -y install unzip

mkdir build/tmp/linux
mkdir build/tmp/win

tar -xvf build/tmp/linux.tar.gz -C build/tmp/linux
unzip build/tmp/win.zip -d build/tmp/win/

cp build/tmp/linux/go-ipfs/ipfs build/linux/
cp build/tmp/win/go-ipfs/ipfs.exe build/win


rm build/tmp -r

echo "linux..."
sed "s/\/\/\!win://" ipfs-cli.go > ipfs_cli_not_win.go
env GOOS=linux GOARCH=386 go build -o build/linux/ipfs-cli ipfs_cli_not_win.go
#rm _ipfs_cli_not_win.go

echo "win..."
env GOOS=windows GOARCH=386 go build -o build/win/ipfs-cli.exe ipfs-cli.go
