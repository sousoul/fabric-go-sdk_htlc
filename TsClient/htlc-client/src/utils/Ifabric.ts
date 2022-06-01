import { receiveMessageOnPort } from "worker_threads"
import { ResponseDto } from "./responseDto/response"

var rp = require('request-promise')
// let ip = "http://127.0.0.1:9191"
// let ip = "http://192.168.213.7:9191" // MacBook上的虚拟机2
let ip = "http://10.200.5.128:9191" // GPU75

// let ip = "http://202.120.39.3:9191" // connect ECONNREFUSED 202.120.39.3:9191
// let ip = "http://202.120.39.3:9076" // Error: Parse Error: Expected HTTP/

async function sendRequest(options:any):Promise<ResponseDto<any>> {
  let resp:ResponseDto<any>
  try {
    let res = await rp(options)
    return resp = {
      code: 200,
      data: res.data,
      err: ''
    }
  } catch (error) {
    return resp = {
      code: 400,
      data: error.stack,
      err: error.message
    }
  }
}

async function testInvoke(id:string, preImage:string):Promise<ResponseDto<any>> {
  let requestJsonBody = {
    "id": id,
    "pre_image": preImage
  }
  var options = {
    method: 'POST',
    uri: ip + "/htlc/testinvoke",  // 和sdk/main()中relativepath一致
    body: requestJsonBody,
    json: true // Automatically parses the JSON string in the response
  };
  return await sendRequest(options)
}

async function lock(hashValue:string, timeLock:string, value:string,
                    txKey:string, flag:string, spenderIdx:string, receiverIdx:string):Promise<ResponseDto<any>> {

  let requestJsonBody = {
    "hashValue": hashValue,
    "timeLock": timeLock,
    "value": value,
    "txKey": txKey,
    "flag": flag,
    "spenderIdx": spenderIdx,
    "receiverIdx": receiverIdx

  }
  var options = {
    method: 'POST',
    uri: ip + "/htlc/lock",  // 和sdk/main()中relativepath一致
    body: requestJsonBody,
    json: true // Automatically parses the JSON string in the response
  };
  return await sendRequest(options)
}

async function withdraw(preImage:string, id:string, txKey:string,
                        channel_idx:string, org_idx:string):Promise<ResponseDto<any>> {

  let requestJsonBody = {
    "preImage": preImage,
    "id": id,
    "txKey": txKey,
    "channel_idx": channel_idx,
    "org_idx": org_idx,
  }
  var options = {
    method: 'POST',
    uri: ip + "/htlc/withdraw",
    body: requestJsonBody,
    json: true // Automatically parses the JSON string in the response
  };
  return await sendRequest(options)
}

async function audit(balance:string, value:string,
                    txKey:string, channelIdx:string, spenderIdx:string, receiverIdx:string):Promise<ResponseDto<any>> {

  let requestJsonBody = {
    "balance": balance,
    "value": value,
    "txKey": txKey,
    "channelIdx": channelIdx,
    "spenderIdx": spenderIdx,
    "receiverIdx": receiverIdx
  }
  var options = {
    method: 'POST',
    uri: ip + "/htlc/audit",  // 和sdk/main()中relativepath一致
    body: requestJsonBody,
    json: true // Automatically parses the JSON string in the response
  };
  return await sendRequest(options)
}

async function verifytwo(txKey:string, channel_idx:string, org_idx:string, receiverIdx:string):Promise<ResponseDto<any>> {
  let requestJsonBody = {
    "txKey": txKey,
    "channel_idx": channel_idx,
    "org_idx": org_idx,
    "receiverIdx": receiverIdx
  }
  var options = {
    method: 'POST',
    uri: ip + "/htlc/verifytwo",  // 和sdk/main()中relativepath一致
    body: requestJsonBody,
    json: true // Automatically parses the JSON string in the response
  };
  return await sendRequest(options)
}

async function verifytwoall(txKey:string, channel_idx:string, receiverIdx:string):Promise<ResponseDto<any>> {
  let requestJsonBody = {
    "txKey": txKey,
    "channel_idx": channel_idx,
    "receiverIdx": receiverIdx
  }
  var options = {
    method: 'POST',
    uri: ip + "/htlc/verifytwoall",  // 和sdk/main()中relativepath一致
    body: requestJsonBody,
    json: true // Automatically parses the JSON string in the response
  };
  return await sendRequest(options)
}

async function createAccount(accountName:string, passwd:string):Promise<ResponseDto<any>> {
  let requestJsonBody = {
    "address": accountName,
    "passwd": passwd
  }
  var options = {
    method: 'POST',
    uri: ip + "/account/create",
    body: requestJsonBody,
    json: true // Automatically parses the JSON string in the response
  };
  return await sendRequest(options)
}


async function createMidAccount(sender:string, receiver:string, preImage:string, flag:string):Promise<ResponseDto<any>> {
  let requestJsonBody = {
    "sender": sender,
    "receiver": receiver,
    "pre_image": preImage,
    "flag": flag
  }
  var options = {
    method: 'POST',
    uri: ip + "/htlc/midaccount",
    body: requestJsonBody,
    json: true // Automatically parses the JSON string in the response
  };
  return await sendRequest(options)
}

async function createHTLC(sender:string, receiver:string, amount:string, ttl:string, hash:string, passwd:string, mid_address:string):Promise<ResponseDto<any>> {
  let requestJsonBody = {
    "sender": sender,
    "receiver": receiver,
    "amount": amount, 
    "ttl": ttl, 
    "hash": hash, 
    "passwd": passwd, 
    "mid_address": mid_address 
  }

  var options = {
    method: 'POST',
    uri: ip + "/htlc/createbyhash",
    body: requestJsonBody,
    json: true // Automatically parses the JSON string in the response
  };
  return await sendRequest(options)
}

async function withdrawFabricAssets(id:string, preImage:string):Promise<ResponseDto<any>> {
  let requestJsonBody = {
  	"id": id,
	  "pre_image": preImage
  }

  var options = {
    method: 'POST',
    uri: ip + "/htlc/withdraw",
    body: requestJsonBody,
    json: true // Automatically parses the JSON string in the response
  };
  return await sendRequest(options)
}

async function refund(id:string, preImage:string):Promise<ResponseDto<any>> {
  let requestJsonBody = {
  	"id": id,
	  "pre_image": preImage
  }

  var options = {
    method: 'POST',
    uri: ip + "/htlc/refund",
    body: requestJsonBody,
    json: true // Automatically parses the JSON string in the response
  };
  return await sendRequest(options)
}

async function QueryHTLC(id:string):Promise<ResponseDto<any>> {
  let requestJsonBody = {
  	"id": id
  }

  var options = {
    method: 'POST',
    uri: ip + "/htlc/query",
    body: requestJsonBody,
    json: true // Automatically parses the JSON string in the response
  };
  return await sendRequest(options)
}

async function faucet(to:string, amount:string):Promise<ResponseDto<any>> { 
  let requestJsonBody = {
    "from": "account-assert-genesis-account",
    "to": to,
    "amount": amount,
    "passwd": "12345678"
  }

  var options = {
    method: 'POST',
    uri: ip + "/account/transfer",
    body: requestJsonBody,
    json: true // Automatically parses the JSON string in the response
  };
  return await sendRequest(options)
}

// createAccount("asadasdazqb", "babc")
// queryAccount("zqa")
// queryAccount("zqb")
// faucet("zqa", "1000")
// createMidAccount("zqa", "0242c0436daa4c241ca8a793764b7dfb50c223121bb844cf49be670a3af4dd18", "hash")
// createHTLC("zqa", "zqb", "100", "2000", "0242c0436daa4c241ca8a793764b7dfb50c223121bb844cf49be670a3af4dd18", "aabc", "zqa0");
// withdraw("8659f8056bfb94b143bd3f0fb36236862614eb303434fe5322dd9ab6715fbc6e", "rootroot")
// queryAccount("zqa")
// queryAccount("zqb")


export { 
  createAccount,
  createMidAccount,
  createHTLC,
  withdrawFabricAssets,
  refund,
  QueryHTLC,
  faucet,

  testInvoke,
  lock,
  withdraw,
  audit,
  verifytwo,
  verifytwoall
}
