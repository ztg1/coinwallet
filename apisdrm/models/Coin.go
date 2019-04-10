package models

import "github.com/astaxie/beego/orm"

// TableName 设置BackendUser表名
func (a *Coin) TableName() string {
	return CoinTBName()
}
type Coin struct {
	Id                    int
	Fid                    int    //负极id
	CoinName              string `orm:"size(32)"`
	CoinTitle             string `orm:"size(32)"`
	ContractAddr          string     //合约地址
	Decimal               int
	AddTime               int64
	EndTime               int64
	Status                int
}


func CoinList()([]*Coin,error)  {
	query:=orm.NewOrm().QueryTable(CoinTBName())
	data:=make([]*Coin,0)
	_,err:=query.Filter("status", "1").All(&data)
	if err!=nil{
		return nil,err
	}
	return data,nil
}


func CoinOn(id int)(*Coin,error)  {
	query:=orm.NewOrm()
	m:=Coin{Id:id}
	err:=query.Read(&m)
	if err !=nil{
		return nil,err
	}

	return &m,nil
}

//通过名字查找
func CoinNa(name string)(*Coin ,error)  {
	query:=orm.NewOrm()
	m:=Coin{CoinName:name}
	err:=query.Read(&m,"coin_name")
	if err !=nil{
		return nil,err
	}
	return &m,nil

}


//查找对应币种 的代币
func CoinListD(name string)([]*Coin,error)  {
	query:=orm.NewOrm().QueryTable(CoinTBName())
	data:=make([]*Coin,0)

    first,err:=CoinNa(name)
    if err !=nil{
    	return nil,err
    }
	_,errs:=query.Filter("status", "1").Filter("fid",first.Id).All(&data)
	if errs!=nil{
		return nil,errs
	}
	return data,nil

}

//通过合约查找
func CoinContract(hax string)(*Coin ,error)  {
	query:=orm.NewOrm()
	m:=Coin{ContractAddr:hax}
	err:=query.Read(&m,"contract_addr")
	if err !=nil{
		return nil,err
	}
	return &m,nil

}
