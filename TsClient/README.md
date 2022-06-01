# 客户端

## 1. 功能介绍
本部分包含使用typescript语言编写的跨链交易脚本。通过执行脚本模拟客户端来向区块链网络发送HTTP请求，从而完成对应的跨链交易流程。

## 2. 脚本的运行
执行以下命令，安装所需依赖并运行脚本。
```
./run.sh
```
run.sh最终运行testWF.ts文件，在一个for循环中共完成5次跨链资产兑换交易，并将每次资产兑换交易的每步所需的时长记录到./htlc-client/cc2.xlsx中。下图展示了前3次跨链交易的时间开销：

<img width="282" alt="截屏2022-06-02 上午6 33 12" src="https://user-images.githubusercontent.com/49592082/171512689-d9fe7fa3-b3e2-48fe-ac9e-fd92131bc81b.png">
