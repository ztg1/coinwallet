package models

import (
	"github.com/astaxie/beego/orm"
)

/*
 扫描区块的记录
*/


func (b *Block) TableName() string {
	return BlockeTBName()
}

type Block struct {
	Id                    int
	CoinName              string `orm:"size(32)"`
	AddNum                string
	EndNum                string
	New_num               string
}


//查询
func BlockeFrirst(CoinName string) (*Block,error)  {

	query:=orm.NewOrm()
	m:=Block{CoinName:CoinName}
	err:=query.Read(&m,"coin_name")
	if err !=nil{
		return nil,err
	}

	return &m,nil
	
}

//设置扫描的区块

func BlockSet(CoinName string,num string)(n int64,err error)  {

	o := orm.NewOrm()
	blocke := Block{CoinName: CoinName}
	if o.Read(&blocke,"coin_name") == nil {
		blocke.New_num = num
		if n, err := o.Update(&blocke); err == nil {

			return n,nil
		}
	}
  return 0,err
}

//设置扫描的区块

func BlockSetAddNum(CoinName string,num string)(n int64,err error)  {

	o := orm.NewOrm()
	blocke := Block{CoinName: CoinName}
	if o.Read(&blocke,"coin_name") == nil {
		blocke.AddNum = num
		if n, err := o.Update(&blocke); err == nil {

			return n,nil
		}
	}
	return 0,err
}




