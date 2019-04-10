package controllers

import (
	"apisdrm/extend"
	"apisdrm/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"log"
	"time"
)

//格式出书
type Msg struct {
	Status string         `json:"status"`
	Err   error           `json:"err"`
	Data  []*models.Coin  `json:"data"`
}

type EthCoinController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all Coin    查看所有币种  不需要什么参数
// @Success 1 {object} models.Coin
// @router / [get]
func (c *EthCoinController) GetAll() {

	data,err:=   models.CoinList()
 	if err != nil{
		c.Data["json"] = map[string]string{"status": "0","err":err.Error()}
	}else {
		msg:=Msg{Status:"1",Err:nil,Data:data}
		c.Data["json"] =msg
	}
	c.ServeJSON()

	return
}


// @Title CreatePrivateKey 创建私钥
// @Description create Private Key  address
// @Param	body		body 	models.Coin	true		"body for Coin content"
// @Success 200 {string}  Private Key
// @Failure 403 body is empty
// @router /privatekey [post]
func (c *EthCoinController) PrivateKey() {

       CoinNmae:= c.GetString("CoinName")
       if CoinNmae ==""{
		   c.Data["json"] = "coin name err"
		   c.ServeJSON()
	   }

      _,err:=  models.CoinNa(CoinNmae)
      if err !=nil{
		  c.Data["json"] = err
		  c.ServeJSON()
	  }

    eth:=impl.Eth{}
	privkey,err1:=  eth.StatusKey()
      if err1!=nil{
		  c.Data["json"] = err1
		  c.ServeJSON()
	  }

	//创建钱包 记录到数据库

	c.Data["json"] = privkey
	c.ServeJSON()

}

//密钥库创建钱包
// @Title Creat Account    密钥库创建钱包
// @Description Creat Account   密钥库创建钱包
// @Param	body		body 	models.CreatAccount	true		"body for coin content"
// @Success 200 {int} models.Coin.Id
// @Failure 403 body is empty
// @router /creataccount [post]
func (c *EthCoinController) CreatAccount() {
	PassWord:= c.GetString("PassWord")
	if PassWord ==""{
		c.Data["json"] = map[string]string{"status": "0","err":"PassWord err"}
		c.ServeJSON()
	}
	e:=impl.Eth{}
	accut,err:=e.NewAccount(PassWord)
	if err !=nil{
		c.Data["json"] = map[string]string{"status": "0","err":err.Error()}
		c.ServeJSON()
	}

	c.Data["json"] = accut

	c.ServeJSON()
}


//查看账户余额
// @Title   AccountBalance   查看账户余额
// @Description AccountBalance
// @Param	body		body 	models.AccountBalance	true		"body for AccountBalance content"
// @Success 200 {int} models.AccountBalance.Id
// @Failure 403 body is empty
// @router /accountbalance [post]
func (c *EthCoinController) AccountBalance() {
	Address:= c.GetString("Address")
	if Address ==""{
		c.Data["json"] =map[string]string{"status": "0","err":"coin Address err"}
		c.ServeJSON()
	}
	e:=impl.Eth{}


	//地址检查
	if res:=e.AddressCheck(Address);res !=true{

		c.Data["json"] =map[string]string{"status": "0","err":"address err"}
		c.ServeJSON()
	}

	balance, err :=e.AccountBalance(Address)
	if err !=nil{
		c.Data["json"] =map[string]string{"status": "0","err":err.Error()}
		c.ServeJSON()
	}

	c.Data["json"] =balance

	c.ServeJSON()
}




// @Title Creat RecordAccount
// @Description RecordAccount         记录用户的钱包
// @Param	UserName		query 	string	true		"用户名"
// @Param	UserId		   query 	string	true		"用户标示符"
// @Param	FromAddress		query 	string	true		"钱包地址"
// @Param	CoinName		query 	string	true		"币种名字"
// @Param	ContractAddr	query 	string	true		"合约地址  不是代币就给一个空值"
// @Success 1 {} {"Id": "\u0014","status": "1","success": "操作成功"}
// @Failure 403 body is empty
// @Failure 0 未知的错误
// @router /recordaccount [post]
func (c *EthCoinController) RecordAccount() {
	//交易发送成功插入数据库
	wallet:=models.WalletLog{}
	wallet.UserName=c.GetString("UserName")          //用户身份
	wallet.UserId=c.GetString("UserId")             //用户ID（唯一标识符）
	wallet.FromAddress=c.GetString("FromAddress")   //用户钱包地址
	wallet.CoinName=c.GetString("CoinName")         //币种名字
	wallet.ContractAddr=c.GetString("ContractAddr")         //合约地址  没有就为空
	wallet.AddTime=time.Now().Unix()
	wallet.Status=0

	valid := validation.Validation{}
	b, err := valid.Valid(&wallet)
	if err != nil {
		// handle error
		c.Data["json"] =map[string]string{"status": "402","err":err.Error()}
		c.ServeJSON()
		return
	}

	if !b {
		// validation does not pass
		// blabla...
		for _, err := range valid.Errors {
			c.Data["json"] =map[string]string{"status": "403","err":err.Key+err.Message}
			c.ServeJSON()
			return
		}
	}


	query:=orm.NewOrm()
	if wallet.Id==0{
		 first,errs:=models.WalletLogOn(wallet.UserId)
		 if errs!=nil{
			 c.Data["json"] =map[string]string{"status": "0","err":errs.Error()}
			 c.ServeJSON()
			 return
		 }

		 if first!=nil{
			 c.Data["json"] =map[string]string{"status": "0","err":"该记录已经存在了"}
			 c.ServeJSON()
		 }

			if _, err := query.Insert(&wallet); err != nil {
				c.Data["json"] =map[string]string{"status": "0","err":err.Error()}
				c.ServeJSON()
				return
			}

	}else {
			c.Data["json"] =map[string]string{"status": "0","err":"非法操作"}
			c.ServeJSON()
		    return
	}

	c.Data["json"] =map[string]string{"status": "1","Id":string(wallet.Id),"success":"操作成功"}

	c.ServeJSON()

	return
}




// @Title Creat TradeOut
// @Description TradeOut    转账
// @Param	UserId		   query 	string	true		"用户标示符"
// @Param	FromAddress		query 	string	true		"钱包地址"
// @Param	ToAddress		query 	string	true		"发送地址"
// @Param	CoinName		query 	string	true		"币种名字"
// @Param	SignedTx		query 	string	true		"交易 SignedTx"
// @Success 1 {}  {"Id": "\u0014","status": "1","success": "操作成功"}
// @Failure 0  未知的错误
// @Failure 403 body is empty
// @router /tradeout [post]
func (c *EthCoinController) TradeOut() {
	signedTx:=c.GetString("SignedTx")
	if signedTx ==""{
		c.Data["json"] =map[string]string{"status": "0","err":"SignedTx empty "}
		c.ServeJSON()
		return
	}

	valid:=validation.Validation{}
	trad:=models.TradeOut{}
	trad.UserId=c.GetString("UserId")
	trad.CoinName=c.GetString("CoinName")
	trad.FromAddress=c.GetString("FromAddress")
	trad.ToAddress=c.GetString("ToAddress")
	trad.AddTime=time.Now().Unix()
	trad.Status=0
	b,err:=valid.Valid(&trad)
	if err !=nil{
		c.Data["json"] =map[string]string{"status": "0","err1":err.Error()}

		c.ServeJSON()
		return
	}

	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			c.Data["json"] =map[string]string{"status": "403","err":err.Key+string(err.Message)}
			c.ServeJSON()
			return
		}
	}


	e:=impl.Eth{}
	e.NewKeyStore()  //

	//
	tx,err:= e.TradeAccount(signedTx)  //发送交易

	//
	if err !=nil{
		c.Data["json"] =map[string]string{"status": "0","err1":err.Error()}
		c.ServeJSON()
		return
	}
	trad.Tx=tx
    c.Save(&trad)
	return
}



//转账添加/修改
func (c *EthCoinController)Save(trad *models.TradeOut)  {
	//插入数据库
	o:=orm.NewOrm()  //modle 实例化
	if trad.Id==0{
		if _,err:=o.Insert(trad);err!=nil{
			c.Data["json"] =map[string]string{"status": "0","err":err.Error()}
		}
	}else {
		if _, err := o.Update(trad); err != nil {
			c.Data["json"] =map[string]string{"status": "0","err":err.Error()+string(trad.Id)}
		}
	}

	c.Data["json"] =map[string]string{"status": "1","success":"操作成功"}
	c.ServeJSON()
}

// @Title SetReturnIn
// @Description SetReturnIn    设置 转入成功的记录
// @Param	uid		query 	string	true		"用户身份id "
// @Param	tx		query 	string	true		"交易哈希 Tx"
// @Success 1 {status:1} {"status": "1","err":"操作成功"} {"status": "0","err":"某个参数错误 "}
// @router /setreturnin [post]
func (c *EthCoinController) SetReturnIn() {

	Uid:=c.GetString("uid")
	Tx:=c.GetString("tx")
	if Uid =="" &&  Tx==""{
		c.Data["json"] =map[string]string{"status": "0","err":"某个参数错误 "}
		c.ServeJSON()
		return
	}

   if bools:=models.TradeInSet(Uid,Tx);bools == false{
	   c.Data["json"] =map[string]string{"status": "0","err":"操作失败"}
	   c.ServeJSON()
	   return
   }
	c.Data["json"] =map[string]string{"status": "1","err":"操作成功"}
	c.ServeJSON()
	return
}

// @Title SetReturnOut
// @Description SetReturnOut    设置 转出成功的记录
// @Param	uid		query 	string	true		"用户身份id "
// @Param	tx		query 	string	true		"交易哈希 Tx"
// @Success 1 {status:1} {"status": "1","err":"操作成功"} {"status": "0","err":"某个参数错误 "}
// @router /setreturnout [post]
func (c *EthCoinController) SetReturnOut() {

	Uid:=c.GetString("uid")
	Tx:=c.GetString("tx")
	if Uid =="" &&  Tx==""{
		c.Data["json"] =map[string]string{"status": "0","err":"某个参数错误 "}
		c.ServeJSON()
		return
	}

	if bools:=models.TradeOutSet(Uid,Tx);bools == false{
		c.Data["json"] =map[string]string{"status": "0","err":"操作失败"}
		c.ServeJSON()
		return
	}
	c.Data["json"] =map[string]string{"status": "1","err":"操作成功"}
	c.ServeJSON()
	return
}




