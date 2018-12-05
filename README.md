# ShareLedger #

ShareLedger is a custom-designed distributed blockchain with [Tendermint](https://github.com/tendermint/tendermint) as a Byzantine-Fault Tolerant State Machine. ShareLedger provides essential building blocks for any rental/sharing services on top of it.

ShareLedger currently ultilizes a dual-token mechanism, SharePay (SHRP) and ShareToken (SHR). The former provides a stable currency for any additional services running on top of ShareLedger while the latter acts as an ultility token.



## Requirements ##

* [`go`](https://golang.org/doc/install) - compile tool. Version >=1.10
* [`dep`](https://github.com/golang/dep) - package management tool
* [`tendermint`](https://github.com/tendermint/tendermint) - consensus algorithm v.0.21.0

## Build ##

* `make get_vendor_deps` to install neccesary dependant packace to `vendor` folder
* `make build` to build ShareLedger 

## Install ShareLedger ##

The fatest and easiest way to install `ShareLedger` is to run our os-specfic application which guides you through four steps to setup and run a MasterNode. [TODO](https://todo.com)


### From Binary

To download pre-built binaries, see the [releases page](https://github.com/sharering/shareledger/releases).


### From Source

You'll need `go`[installed](https://golang.org/doc/install)


#### Get Source Code

```
git clone https://github.com/sharering/shareledger.git
cd shareledger
```


#### Get Tools & Dependencies

```
make get_vendor_deps
```


#### Compile
```
make build
```

This will compile ShareLedger and put the binary in `./build`


#### Run

To start a one-node ShareLedger
```
shareledger init
shareledger node
```

Type `shareledger -h` to get more detailed information on how to execute ShareLedger


#### Upgrade
```
git pull origin master
make get_vendor_deps
make build
```

