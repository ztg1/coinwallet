package models

import (
	"apisdrm/extend"
	"apisdrm/utils"
	"context"
	"fmt"
	"github.com/ethereum-development-with-go-book/code/contracts_erc20"
	"github.com/ethereum-development-with-go-book/code/util"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strconv"
	"strings"
	"time"
)
type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}
const  ETHURL string ="https://mainnet.infura.io"

// LogApproval ..
type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}
type TradeInScab struct {

}

//查询监听最新的币种区块
func (m *TradeInScab) ListCoinBlock()   {
	data,err:= CoinList()
	if err !=nil{
		utils.LogError(fmt.Sprintf("查询区块记录失败:%s\n",err))//日志
		return
	}

	if len(data)>0{
		for _,i:=range data{
			switch i.CoinName {
			case "eth":
				go  ListEthd()    //以太坊代币
				//go  ListEth()      //以太坊
			case "btc":
				fmt.Println("btc .... 待开发中")
			case "usdt":
				fmt.Println("usdt .... 待开发中")
			}
		}
	}


}

/*
  以太坊
*/
func ListEth()  {
       ethcoin:= new(impl.Eth)
	   client,err:= ethcoin.ListenBlock()
	   defer  client.Close()     //关闭链接
	   if err!=nil{

		   utils.LogError(fmt.Sprintf("链接节点错误:%s\n",err))//日志
		   return
	   }
		data,err:= 	BlockeFrirst("eth")  //找到上一次查询到的区块
		if err !=nil{

			utils.LogError(fmt.Sprintf("查找上一次节点错误:%s\n",err))//日志
			return
		}

		num,_:=strconv.ParseInt(data.New_num,0,64)
		blockNumber := big.NewInt(num+1)
		block, err := client.BlockByNumber(context.Background(), blockNumber)
		if err != nil {
			utils.LogError(fmt.Sprintf("获取区块失败:%s,时间=%s\n",err,time.Now().Unix()))//日志
			return
		}
		//查找交易记录
		for _, tx := range block.Transactions() {
			//查看事务
			receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				fmt.Println("tx====",tx.Hash().Hex())

				utils.LogTrace(fmt.Sprintf("查看事务receipt err =%s,hx=%s,时间=%s\n",err,tx.Hash(),time.Now().Unix()))//日志
				continue
			}
			if receipt.Status !=1{
				utils.LogTrace(fmt.Sprintf("查看事务receipt.Status =%s,时间=%s\n",0,time.Now().Unix()))//日志

				continue
			}

			if tx.To()==nil{

				continue
			}
			//查该地址是否在有人用
			dataAddress,err:= WallerFirset(tx.To().Hex())                  //这里会报错
			if dataAddress == nil{
				utils.LogTrace(fmt.Sprintf("没有人用该地址 =%s,时间=%s\n",tx.To().Hex(),time.Now().Unix()))//日志
				continue
			}
			//找到发送 地址
			chainID, err := client.NetworkID(context.Background())
			if err != nil {

				utils.LogTrace(fmt.Sprintf("高度=%s查找发送地址失败NetworkIDerr=%s,时间=%s\n",blockNumber,err,time.Now().Unix()))//日志
				continue
			}
			var frome string
			msg, err := tx.AsMessage(types.NewEIP155Signer(chainID))
			if err != nil {
				utils.LogTrace(fmt.Sprintf("查找发送地址失败 发送地址设置为 空 =%s,时间=%s\n",err,time.Now().Unix()))//日志
				frome=""
			}else {
				frome=string(msg.From().Hex())
			}

			//查询转入记录表是否存在 不存在就插入
			tradeon,_:=TradeOn(tx.Hash().Hex())

			if tradeon != nil{                           //存在退出
				utils.LogTrace(fmt.Sprintf("存在重复的的记录 TX=%s,时间=%s\n",tx.Hash().Hex(),time.Now().Unix()))//日志
				continue
			}

		   // gas 转化  //插入数据库
			wei := new(big.Int)
			wei.SetString(tx.GasPrice().String(), 10)
			GasPrice := util.ToDecimal(wei, 18)

			wei.SetString(tx.Value().String(),10)
			value:=util.ToDecimal(wei,18)

		  m:=TradeIn{CoinName:"eth",UserId:dataAddress.UserId,FromAddress:frome,ToAddress:tx.To().Hex(),Tx:tx.Hash().Hex(),AddTime:int64(block.Time().Uint64()),Value:value.String(),Gas:int64(tx.Gas()),GasPrice:GasPrice.String(),BlockNumber:block.Number().String()}
		  _,err1:=TradeSave(&m)
		  if err1!=nil{
			  //fmt.Println("err TradeSave =",errr)
			  utils.LogTrace(fmt.Sprintf("保存记录失败 TX=%s,高度=%s,时间=%s\n",tx.Hash().Hex(),block.Number().String(),time.Now().Unix()))//日志
			  continue
		  }


		}
	     //更新扫描的区块记录
		_,errs:=BlockSet("eth",block.Number().String())
		if errs !=nil {
			utils.LogTrace(fmt.Sprintf("更新扫描高度失败 =%s,时间=%s\n",errs,time.Now().Unix()))//日志
			return
		}


	return

}

/*
     以太坊代币转入扫描
*/
func ListEthd(){
	client, err := ethclient.Dial(ETHURL)
	defer client.Close()
	if err != nil {
		utils.LogError(fmt.Sprintf("扫描代币链接错误 =%s,时间=%s\n",err,time.Now().Unix()))//日志
		return
	}

	ethd,err:=CoinListD("eth")   //找到以太坊代币
	if err !=nil{
		utils.LogTrace(fmt.Sprintf("ETH代币不存在 =%s,时间=%s\n",err,time.Now().Unix()))//日志
		return
	}

	fmt.Printf("Log Block Number: %d\n", "代币扫描")

	data,err:= 	BlockeFrirst("eth")  //找到上一次查询到的区块
	if err !=nil{

		utils.LogError(fmt.Sprintf("查找上一次节点错误:%s\n",err))//日志
		return
	}
	AddNum,_:=strconv.ParseInt(data.AddNum,0,64)
	EndNum,_:=strconv.ParseInt(data.New_num,0,64)
	tokenaddress:=[]common.Address{}      //合约地址
	for _,i:=range ethd {
		 address:= common.HexToAddress(i.ContractAddr)
		 tokenaddress=append(tokenaddress,address)
	}
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(AddNum-1),
		//FromBlock: big.NewInt(7534933),
		ToBlock:   big.NewInt(EndNum),
		//ToBlock:   big.NewInt(7534935),
		Addresses: tokenaddress,
	}

	logs, err := client.FilterLogs(context.Background(), query)     //来过滤日志：
	if err != nil {
		utils.LogError(fmt.Sprintf("来过滤日志错误 =%s,时间=%s\n",err,time.Now().Unix()))//日志
		return
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(token.TokenABI)))
	if err != nil {
		utils.LogError(fmt.Sprintf("解析JSON abi错误 =%s,时间=%s\n",err,time.Now().Unix()))//日志
		return
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	LogApprovalSig := []byte("Approval(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)

	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		//fmt.Printf("Log Index: %d\n", vLog.Index)

		//fmt.Printf("Log TX: %s\n", string(vLog.TxHash.Hex()))
		switch vLog.Topics[0].Hex() {
			case logTransferSigHash.Hex():
				var transferEvent LogTransfer
				err := contractAbi.Unpack(&transferEvent, "Transfer", vLog.Data)
				if err != nil {
					utils.LogError(fmt.Sprintf("代币contractAbi.Unpack错误 =%s,时间=%s\n",err,time.Now().Unix()))//日志
					continue
				}
				transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
				transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())
				//查找是否有人在用这个地址
				dataAddress,_:= WallerFirset(transferEvent.From.Hex())
				if dataAddress == nil{
					utils.LogTrace(fmt.Sprintf("没有人用该代币地址 =%s,时间=%s\n",0,time.Now().Unix()))//日志
					continue
				}

				//fmt.Println("TX=",string(vLog.TxHash.Hex()))
				//查询转入记录表是否存在 不存在就插入
				tradeon,_:=TradeOn(string(vLog.TxHash.Hex()))
				if tradeon != nil{                           //存在退出
					utils.LogTrace(fmt.Sprintf("代币存在重复的的记录 TX=%s,时间=%s\n",string(vLog.TxHash.Hex()),time.Now().Unix()))//日志
					continue
				}

                var decimal int=18
			   //fmt.Println("合约地址=",vLog.Address.Hex())  //从合约地址读取币种的小数位数
                coindatas,err:=CoinContract(vLog.Address.Hex())
                if err!=nil{
					utils.LogTrace(fmt.Sprintf("代币存查找合约错误 合约=%s,时间=%s\n",vLog.Address.Hex(),time.Now().Unix()))//日志
					continue
				}
                decimal=coindatas.Decimal
				//格式化数字
				wei := new(big.Int)
				wei.SetString(transferEvent.Tokens.String(), 10)
				value := util.ToDecimal(wei, decimal)

				mm:=TradeIn{CoinName:"eth",UserId:dataAddress.UserId,FromAddress:transferEvent.To.Hex(),ToAddress:transferEvent.From.Hex(),Tx:string(vLog.TxHash.Hex()),AddTime:time.Now().Unix(),Value:value.String(),BlockNumber:strconv.Itoa(int(vLog.BlockNumber))}
				_,err1:=TradeSave(&mm)
				if err1!=nil{
					utils.LogTrace(fmt.Sprintf("代币保存记录失败 TX=%s,高度=%s,时间=%s\n",transferEvent.From.Hex(),string(vLog.BlockNumber),time.Now().Unix()))//日志
					continue
				}

			case logApprovalSigHash.Hex():
				utils.LogTrace(fmt.Sprintf("logApprovalSigHash.Hex()有执行,时间=%s\n",time.Now().Unix()))//日志
				continue
			// 不知道干什么用的
		}
	}

	//更新开始扫描的区块
	_,errs:=BlockSetAddNum("eth",data.New_num)
	if errs !=nil {
		utils.LogTrace(fmt.Sprintf("代币更新扫描区块失败 =%s,时间=%s\n",errs,time.Now().Unix()))//日志
		return
	}
	return

}








