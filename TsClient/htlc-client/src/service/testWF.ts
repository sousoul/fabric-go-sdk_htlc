const fs = require('fs')

import {
  testInvoke,
  lock,
  withdraw,
  audit,
  verifytwo,
  verifytwoall
} from '../utils/Ifabric'

import { getCurrentTime, appendData } from './utils'

let hashLock = '0242c0436daa4c241ca8a793764b7dfb50c223121bb844cf49be670a3af4dd18'
let expireDuration:number = 20000000000
let expireTimestamp = Date.now() + expireDuration
let preimageBytesHex = "0x726f6f74726f6f74"
let preimageString = "rootroot"

let fabricHTLCId = ""

var preimage = "rootroot"

var myDate = new Date();

async function testCallingBackend() {
  console.log("------测试前后端调用------")
  let startTime = getCurrentTime()
  let res = await testInvoke(fabricHTLCId, preimage)
  let endTime = getCurrentTime()
  let consumeTime = endTime - startTime;
  await appendData(startTime, endTime, consumeTime, "TestCallingBackend")
  console.log("OUTPUT:", res.data)
  console.log("\n\n\n")
}

async function LockAsset(timelock:number, value:number, txKey:string, flag:string, spenderIdx:number, receiverIdx:number) {
  console.log("------锁定资产------")
  let startTime = getCurrentTime()

  if (flag=="hashValue"){
    var res = await lock(hashLock, timelock.toString(), value.toString(), txKey, flag, spenderIdx.toString(), receiverIdx.toString())
  } else {
    var res = await lock(preimage, timelock.toString(), value.toString(), txKey, flag, spenderIdx.toString(), receiverIdx.toString())
  }

  let endTime = getCurrentTime()
  let consumeTime = endTime - startTime;
  await appendData(startTime, endTime, consumeTime, "LockAsset")

  console.log("OUTPUT:", res.data)
  console.log("\n\n\n")

  return JSON.stringify(res.data)
}

async function WithdrawAsset(preImage:string, id:string, txKey:string, channel_idx:number, org_idx:number) {
  console.log("------领取资产------")
  let startTime = getCurrentTime()
  let res = await withdraw(preimage, id, txKey, channel_idx.toString(), org_idx.toString())
  let endTime = getCurrentTime()
  let consumeTime = endTime - startTime;
  await appendData(startTime, endTime, consumeTime, "WithdrawAsset")
  console.log("OUTPUT:", res.data)
  console.log("\n\n\n")
}

async function Audit(balance:number, value:number,
                     txKey:string, channelIdx:number, spenderIdx:number, receiverIdx:number){
  console.log("------响应监管------")
  let startTime = getCurrentTime()
  let res = await audit(balance.toString(), value.toString(), txKey, channelIdx.toString(), spenderIdx.toString(), receiverIdx.toString())
  let endTime = getCurrentTime()
  let consumeTime = endTime - startTime;
  await appendData(startTime, endTime, consumeTime, "Audit")
  console.log("OUTPUT:", res.data)
  console.log("\n\n\n")
}

async function VerifyTwo(txKey:string, channel_idx:number, org_idx:number, receiverIdx:number){
  console.log("------验证------")
  let startTime = getCurrentTime()
  let res = await verifytwo(txKey, channel_idx.toString(), org_idx.toString(), receiverIdx.toString())
  let endTime = getCurrentTime()
  let consumeTime = endTime - startTime;
  await appendData(startTime, endTime, consumeTime, "VerifyTwo")
  console.log("OUTPUT:", res.data)
  console.log("\n\n\n")
}

async function VerifyTwoAll(txKey:string, channel_idx:number, receiverIdx:number){
  console.log("------验证------")
  let startTime = getCurrentTime()
  let res = await verifytwoall(txKey, channel_idx.toString(), receiverIdx.toString())
  let endTime = getCurrentTime()
  let consumeTime = endTime - startTime;
  await appendData(startTime, endTime, consumeTime, "VerifyTwo")
  console.log("OUTPUT:", res.data)
  console.log("\n\n\n")
}

async function testWf() {
  for (let i=0; i<5; i++) {
    let startTime = getCurrentTime()

    // preimage = "root" + i.toString() // 随机选取一个字符串
    // TODO：导入web3库计算哈希值，把htlc结构体中的TxKey字段去掉
    var txKey = (i+1).toString()
    var value1 = 1
    var value2 = 1

    var id1Colon = await LockAsset(100+2*i, value1, txKey, "preimage", 0, 1) // 锁定链1
    var id2Colon = await LockAsset(50+i, value2, txKey, "hashValue", 1, 0) // 锁定链2

    var id1 = id1Colon.slice(1,-1)
    var id2 = id2Colon.slice(1,-1)
    console.log(id1)
    console.log(id2)

    await WithdrawAsset(preimage, id2, txKey, 1, 0)
    await WithdrawAsset(preimage, id1, txKey, 0, 1)

    await Audit(Asset1[0]-value1, value1, txKey, 0, 0, 1) // 监管链1
    await Audit(Asset2[1]-value2, value2, txKey, 1, 1, 0) // 监管链2

    // await VerifyTwo(txKey, 0, 1, 1) // org2验证chain1
    // await VerifyTwo(txKey, 1, 0, 0) // org1验证chain2
    await VerifyTwoAll(txKey, 0, 1) // 验证chain1
    await VerifyTwoAll(txKey, 1, 0) // 验证chain2

    Asset1[0] = Asset1[0]-value1
    Asset2[1] = Asset2[1]-value2
    let endTime = getCurrentTime()
    let consumeTime = endTime - startTime;
    await appendData(startTime, endTime, consumeTime, "一次跨链交易时间")
  }
}

async function testWF2() {
  for (let i=0; i<1; i++) {
    await testCallingBackend()
  }
}

// 这里和SDK中保持一致
var Asset1:number[] = [10000000, 10000000] // 链1上的初始余额
var Asset2:number[] = [10000000, 10000000] // 链2上的初始余额

testWf()
// testWF2()