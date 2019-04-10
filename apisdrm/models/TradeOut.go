package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

func (a *TradeOut) TableName() string {
	return TradeOutTBName()
}

type TradeOut struct {
	Id              int
	CoinName        string   `valid:"Required;MaxSize(32)"`
	UserId          string   `valid:"Required;MaxSize(64)"`
	FromAddress    string   `valid:"Required;MaxSize(161)"`
	ToAddress       string   `valid:"Required;MaxSize(161)"`
	Tx              string
	AddTime         int64
	EndTime         int64
	Status          int
}

/*
 查询所有没有
*/

func TradeOutList(status string,userid string)([]*TradeOut,error)  {
	query:=orm.NewOrm().QueryTable(TradeOutTBName())
	data:=make([]*TradeOut,0)
	var err error

	if status=="-1"{
		_,err=query.All(&data)
	}else if len(userid)>0 {
		_,err=query.Filter("status", status).Filter("user_id",userid).All(&data)
	}else {
		_,err=query.Filter("status", status).All(&data)
	}
	if err!=nil{

		return nil,err
	}

	return data,nil
}

//通过查找tx 更新状态
func (m *TradeOut)TradeUpdate(tx string)(userid string,Tx string,err error)  {
	  o := orm.NewOrm()
	  tradeout:=TradeOut{Tx:tx}
	 if err:=o.Read(&tradeout,"tx"); err== nil{
	 	tradeout.Status= 1
	 	tradeout.EndTime=time.Now().Unix()
	 	if num,err:= o.Update(&tradeout,"status","end_time");err ==nil{
	 		fmt.Println("ok=",num)
	 		fmt.Println("uer_id=",tradeout.UserId)
	 		return tradeout.UserId,tradeout.Tx,nil
		}
	 }
	  return tradeout.UserId,tradeout.Tx,err
}


func TradeOutSet(uid string,tx string) bool {
	o:=orm.NewOrm()
	n, err := o.QueryTable(TradeOutTBName()).Filter("user_id", uid).Filter("tx",tx).Update(orm.Params{
		"status": "2",
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

