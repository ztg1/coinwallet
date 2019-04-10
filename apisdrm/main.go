package main

import (
	"apisdrm/models"
	_ "apisdrm/routers"
	_ "apisdrm/sysinit"
	"github.com/astaxie/beego"
	"time"
)



func main() {

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

    //go Sacann()     //扫描 交易转出是否完成
	go SannIn()    //监听区块 查找转入记录
	beego.Run()
}

func Sacann()  {

		for  {
			ts:= new(models.TradeOutScab)
			ts.List()
			time.Sleep(10*time.Second)
		}

}

// 监听最新区块  扫描转币记录 推送给客户端
func SannIn()  {

	for  {
		m:=new(models.TradeInScab)
		m.ListCoinBlock()
		time.Sleep(10*time.Second)
	}


}

