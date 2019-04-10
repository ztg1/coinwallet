package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

func (m*TradeIn)TableName() string {

	 return TradeInTBName()
}

type TradeIn struct {
	Id              int
	CoinName        string   `valid:"Required;MaxSize(32)"`
	UserId          string   `valid:"Required;MaxSize(64)"`
	FromAddress     string   `valid:"Required;MaxSize(161)"`
	ToAddress       string   `valid:"Required;MaxSize(161)"`
	Tx              string
	Value           string
	GasPrice        string
	Gas             int64
	AddTime         int64
	BlockNumber     string
	Status          int
}

//查询看看有没有这个hash 值
func TradeOn(Tx string)(*TradeIn,error)  {
	query:=orm.NewOrm()
	m:=TradeIn{Tx:Tx}
	err:=query.Read(&m,"tx")
	if err == orm.ErrNoRows {
		//fmt.Println("查询不到")
		return nil,err
	} else if err == orm.ErrMissPK {
		//fmt.Println("找不到主键")
		return nil,err
	} else {
		return &m,nil
	}
}

func TradeSave(data *TradeIn)( n int64, err error)  {
	o := orm.NewOrm()
	var trade TradeIn
	trade = *data
	id, err := o.Insert(&trade)
	if err == nil {
		return  id,nil
	}else {
		return 0,err
	}
	
}


//查询转入记录推送出去

func TradeInList(status string,userid string) ([]*TradeIn,error) {
	query:=orm.NewOrm().QueryTable(TradeInTBName())
	data:=make([]*TradeIn,0)
	var err error
	if status=="-1"{
		_,err=query.All(&data)
	}else {
		_,err=query.Filter("status", status).Filter("user_id",userid).All(&data)
	}
	if err!=nil{

		return nil,err
	}

	return data,nil


}

/*
 修改对应的记录
*/

func TradeInSet(uid string,tx string) bool {
	o:=orm.NewOrm()
	n, err := o.QueryTable(TradeInTBName()).Filter("user_id", uid).Filter("tx",tx).Update(orm.Params{
		"status": "1",
	})
	if err !=nil{
		return false
	}
	fmt.Println("n=",n)
	if n==0{
		return false
	}
	return true
}