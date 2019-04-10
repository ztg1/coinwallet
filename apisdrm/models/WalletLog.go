package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"strings"
)

func (a *WalletLog) TableName() string {
	return WalletLogTBName()
}
type WalletLog struct {
  Id            int
  UserName      string `valid:"Required"`
  UserId        string `valid:"Required;"`
  FromAddress   string `valid:"Required;"`
  CoinName      string ` valid:"Required;"`
  ContractAddr  string
  Num           float64
  AddTime       int64
  Status        int
}

func (w *WalletLog)Valid(v *validation.Validation)  {
	if strings.Index(w.UserName, "admin") != -1 {
		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
		 v.SetError("Name", "名称里不能含有 admin")
	}
}



func WalletLogOn(userid string)(*WalletLog,error)  {
	query:=orm.NewOrm()
	m:=WalletLog{UserId:userid}
	err:=query.Read(&m,"user_id")

	if err == orm.ErrNoRows {
		return nil,nil
	} else if err == orm.ErrMissPK {
		return nil,nil
	} else {
		return &m,nil
	}
}

func WallerFirset(address string)(*WalletLog, error)  {

	query:=orm.NewOrm()
	m:=WalletLog{FromAddress:address}
	err:=query.Read(&m,"from_address")

	if err == orm.ErrNoRows {
		return nil,nil
	} else if err == orm.ErrMissPK {
		return nil,nil
	} else {
		return &m,nil
	}

}


