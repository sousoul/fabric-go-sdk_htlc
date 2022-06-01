# fabric-go-sdk_htlc
本仓库为区块链可信交易构件代码。代码基于Hyperledger Fabric进行开发，实现了对跨链资产兑换交易的隐私保护、交易监管。

项目代码包含两个分支：
- main分支下的代码搭建了2个组织的fabric网络，并进行功能测试
- 8orgs分支下的代码搭建了8个组织的fabric网络，并进行功能测试

## 1. 代码结构
项目的开发工作主要分以下三部分：
- SDK代码的开发：./main.go中包含了使用fabric-sdk-go来搭建区块链网络的主要代码。

- Fabric链码的开发：./chaincode/chaincode.go中包含了Withdraw、Lock等链码方法，这些链码方法会进一步调用密码学API实现对应的功能：

  | 名称      | 功能                                                         |
  | --------- | ------------------------------------------------------------ |
  | Lock      | 锁定资产：调用 ZkPutState 完成一笔转账                       |
  | Withdraw  | 领取资产：判断哈西和时间锁是否满足条件，以保证 跨链事务的原子性 |
  | Audit     | 响应监管：调用 ZkAudit                                       |
  | VerifyOne | 验证交易：调用 ZkVerifyOne                                   |
  | VerifyTwo | 验证证明：调用 ZkVerifyTwo                                   |

- 密码学API的开发：./chaincode/vendor/github.com/hyperledger/fabric-chaincode-go/shim/stub.go 中包含了ZkPutState，ZKAudit等密码学API：

  | 名称        | 功能                                                         |
  | ----------- | ------------------------------------------------------------ |
  | ZkPutState  | 为所有组织计算交易金额的承诺值 Com，令牌 Token， 并保存到二维账本 |
  | ZkAudit     | 为所有组织生成范围证明和析取证明，并保存到二维账本           |
  | ZkVerifyOne | 验证一笔转账的金额平衡性;验证某个组织的 Pedersen 承诺的正确性 |
  | ZkVerifyTwo | 验证二维账本上某个组织的范围证明和析取证明                   |

  项目整体的目录结构如下所示：
```

├── README.md

├── chaincode //链码目录

│  ├── chaincode.go // 链码文件，包含Lock、Withdraw等链码方法

│  ├── go.mod

│  ├── go.sum

│  ├── htlc.go 

│  ├── protobuf // 二维账本

│  └── vendor // 链码依赖

├── config.yaml // fabric-sdk-go配置文件（包含2个组织的fabric网络）

├── config_8orgs.yaml // fabric-sdk-go配置文件（包含8个组织的fabric网络）

├── fabric-go-sdk // 可执行文件

├── fixtures // 区块链网络配置文件目录（包含2个组织的fabric网络）

│  ├── channel-artifacts

│  ├── configtx.yaml

│  ├── create_fixtures.sh

│  ├── crypto-config

│  ├── crypto-config.yaml

│  └── docker-compose.yaml

├── fixtures_8orgs（包含8个组织的fabric网络）

│  ├── channel-artifacts

│  ├── configtx.yaml

│  ├── create_fixtures.sh

│  ├── crypto-config

│  ├── crypto-config.yaml

│  └── docker-compose.yaml

├── go.mod

├── go.sum

├── main.go // fabric-sdk-go代码，包含搭建区块链网络、跨链资产兑换交易脚本

├── monitordocker.sh // 用于查看各个容器日志的脚本

├── restart_network.sh // 用于启动各个容器的脚本

└── sdkInit // fabric-go-sdk代码，用于搭建区块链网络，调用链码等

  ├── generateSkPk.go // 生成各组织加密用到的公私钥

  ├── get.go 

  ├── set.go 

  ├── sdkInfo.go

  ├── sdkSetting.go

  └── integration.go
```

## 2. 代码使用
### 2.1 测试代码运行
克隆代码仓库
```
git clone git@github.com:sousoul/fabric-go-sdk_htlc.git
```
进入项目根目录
```
cd fabric-go-sdk_htlc/
```
执行脚本，启动各节点容器并编译sdk代码生成可执行文件：
```bash
./restart_network.sh
```
执行可执行文件，搭建区块链网络并模拟执行一笔跨链交易：

```shell
./fabric-go-sdk 
```

此外，可在另一终端执行monitordocker.sh，监听各个容器输出的日志：

```
./monitordocker.sh fixtures_test
```

