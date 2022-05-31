# fabric-go-sdk_htlc
本仓库为区块链可信交易构件代码。代码基于Hyperledger Fabric进行开发，实现了对跨链资产兑换交易的隐私保护、交易监管。

## 代码结构
chaincode目录下为链码，其中链码方法在





```
.

├── README.md

├── chaincode

│  ├── chaincode.go

│  ├── go.mod

│  ├── go.sum

│  ├── htlc.go

│  ├── htlc_backup.go

│  ├── protobuf

│  ├── protobuf_wo_bulletproof

│  └── vendor

├── config.yaml

├── config_8orgs.yaml

├── fabric-go-sdk

├── fixtures

│  ├── channel-artifacts

│  ├── configtx.yaml

│  ├── create_fixtures.sh

│  ├── crypto-config

│  ├── crypto-config.yaml

│  └── docker-compose.yaml

├── fixtures_8orgs

│  ├── channel-artifacts

│  ├── configtx.yaml

│  ├── create_fixtures.sh

│  ├── crypto-config

│  ├── crypto-config.yaml

│  └── docker-compose.yaml

├── go.mod

├── go.sum

├── main.go

├── monitordocker.sh

├── restart_network.sh

└── sdkInit

  ├── generateSkPk.go

  ├── get.go

  ├── integration.go

  ├── sdkInfo.go

  ├── sdkSetting.go

  └── set.go
```



## 代码使用
./restart_network.sh 
