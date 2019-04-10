package models

import (
	"apisdrm/utils"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"time"
)


type TradeOutScab struct {

}

func (m *TradeOutScab)List()  {
	 TradeOutTBName()
	 data,err:=TradeOutList("0","")    //查询没有交易的完成的记录
	 if err!=nil{
		 utils.LogInfo(fmt.Sprintf("没有记录要扫描:%s,time=%s\n",err,time.Now().Unix()))//日志
		 return
	 }
	if len(data)>0{
		for _,i:=range data{
			
			switch i.CoinName {
			case "eth":
               go  OutEth(i.Tx)
			case "btc":
				fmt.Println("btc .... 开发中")
			case "usdt":
				fmt.Println("usdt .... 开发中")
			}
		}
		}

}

//查找交易情况
func OutEth(TX string)  {

	client, err := ethclient.Dial(ETHURL)
	defer client.Close()

	if err != nil {
		utils.LogInfo(fmt.Sprintf("client链接失败:%s,time=%s\n",err,time.Now().Unix()))//日志
		return
	}
	txHash := common.HexToHash(TX)
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {

		utils.LogInfo(fmt.Sprintf("没有找到该has=%s错误内容=:%s,time=%s\n",txHash,err,time.Now().Unix()))//日志
		return
	}
	if isPending==false{
		//已经交易完成了
		 out:=TradeOut{}
	 	UserId,Tx,err:= out.TradeUpdate(tx.Hash().Hex())
		if err!=nil{
			utils.LogInfo(fmt.Sprintf("修改错误TX=%s错误内容=:%s,time=%s\n",tx.Hash().Hex(),err,time.Now().Unix()))//日志
			return
		}
		utils.LogInfo(fmt.Sprintf("转出扫描成功UserId=%s TX=%s,time=%s\n",UserId,Tx,time.Now().Unix()))//日志
	}
	return
}




