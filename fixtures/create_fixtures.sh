# ./bin/目录下有fabric二进制工具
# 生成组织的证书文件
../bin/cryptogen generate --config=crypto-config.yaml
# 生成通道配置文件
../bin/configtxgen -profile TwoOrgsOrdererGenesis -outputBlock  ./channel-artifacts/genesis.block -channelID fabric-channel  # 生成创世块
../bin/configtxgen -profile TwoOrgsChannel1 -outputCreateChannelTx  ./channel-artifacts/channel1.tx -channelID mychannel1 # 创建channel1.tx
../bin/configtxgen -profile TwoOrgsChannel2 -outputCreateChannelTx  ./channel-artifacts/channel2.tx -channelID mychannel2 # 创建channel2.tx
# 更新组织的锚节点
# org1
../bin/configtxgen -profile TwoOrgsChannel1 -outputAnchorPeersUpdate  ./channel-artifacts/Org1MSPanchors1.tx -channelID mychannel1 -asOrg Org1MSP # 为org1创建通道1上的锚节点文件
../bin/configtxgen -profile TwoOrgsChannel2 -outputAnchorPeersUpdate  ./channel-artifacts/Org1MSPanchors2.tx -channelID mychannel2 -asOrg Org1MSP # 为org1创建通道2上的锚节点文件
# org2
../bin/configtxgen -profile TwoOrgsChannel1 -outputAnchorPeersUpdate  ./channel-artifacts/Org2MSPanchors1.tx -channelID mychannel1 -asOrg Org2MSP # 为org2创建通道1上的锚节点文件
../bin/configtxgen -profile TwoOrgsChannel2 -outputAnchorPeersUpdate  ./channel-artifacts/Org2MSPanchors2.tx -channelID mychannel2 -asOrg Org2MSP # 为org2创建通道2上的锚节点文件
# org3
../bin/configtxgen -profile TwoOrgsChannel1 -outputAnchorPeersUpdate  ./channel-artifacts/Org3MSPanchors1.tx -channelID mychannel1 -asOrg Org3MSP # 为org3创建通道1上的锚节点文件
../bin/configtxgen -profile TwoOrgsChannel2 -outputAnchorPeersUpdate  ./channel-artifacts/Org3MSPanchors2.tx -channelID mychannel2 -asOrg Org3MSP # 为org3创建通道2上的锚节点文件
# org4
../bin/configtxgen -profile TwoOrgsChannel1 -outputAnchorPeersUpdate  ./channel-artifacts/Org4MSPanchors1.tx -channelID mychannel1 -asOrg Org4MSP # 为org4创建通道1上的锚节点文件
../bin/configtxgen -profile TwoOrgsChannel2 -outputAnchorPeersUpdate  ./channel-artifacts/Org4MSPanchors2.tx -channelID mychannel2 -asOrg Org4MSP # 为org4创建通道2上的锚节点文件
# org5
../bin/configtxgen -profile TwoOrgsChannel1 -outputAnchorPeersUpdate  ./channel-artifacts/Org5MSPanchors1.tx -channelID mychannel1 -asOrg Org5MSP # 为org5创建通道1上的锚节点文件
../bin/configtxgen -profile TwoOrgsChannel2 -outputAnchorPeersUpdate  ./channel-artifacts/Org5MSPanchors2.tx -channelID mychannel2 -asOrg Org5MSP # 为org5创建通道2上的锚节点文件
# org6
../bin/configtxgen -profile TwoOrgsChannel1 -outputAnchorPeersUpdate  ./channel-artifacts/Org6MSPanchors1.tx -channelID mychannel1 -asOrg Org6MSP # 为org6创建通道1上的锚节点文件
../bin/configtxgen -profile TwoOrgsChannel2 -outputAnchorPeersUpdate  ./channel-artifacts/Org6MSPanchors2.tx -channelID mychannel2 -asOrg Org6MSP # 为org6创建通道2上的锚节点文件
# org7
../bin/configtxgen -profile TwoOrgsChannel1 -outputAnchorPeersUpdate  ./channel-artifacts/Org7MSPanchors1.tx -channelID mychannel1 -asOrg Org7MSP # 为org7创建通道1上的锚节点文件
../bin/configtxgen -profile TwoOrgsChannel2 -outputAnchorPeersUpdate  ./channel-artifacts/Org7MSPanchors2.tx -channelID mychannel2 -asOrg Org7MSP # 为org7创建通道2上的锚节点文件
# org8
../bin/configtxgen -profile TwoOrgsChannel1 -outputAnchorPeersUpdate  ./channel-artifacts/Org8MSPanchors1.tx -channelID mychannel1 -asOrg Org8MSP # 为org8创建通道1上的锚节点文件
../bin/configtxgen -profile TwoOrgsChannel2 -outputAnchorPeersUpdate  ./channel-artifacts/Org8MSPanchors2.tx -channelID mychannel2 -asOrg Org8MSP # 为org8创建通道2上的锚节点文件
