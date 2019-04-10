package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// init 初始化
func init() {
	orm.RegisterModel(new(Coin),new(WalletLog),new(TradeOut),new(TradeIn),new(Block))
}

// TableName 下面是统一的表名管理
func TableName(name string) string {
	prefix := beego.AppConfig.String("db_dt_prefix")
	return prefix + name
}

func CoinTBName() string  {

	return  TableName("coin")
}

func WalletLogTBName() string  {

	return  TableName("wallet_log")
}

func TradeOutTBName()string  {

	return  TableName("trade_out")
}

func TradeInTBName()string  {

	return TableName("trade_in")

}

func BlockeTBName()string  {

	return TableName("block")

}

