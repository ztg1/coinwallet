package controllers

import (
	"apisdrm/extend"
	"apisdrm/models"
	"apisdrm/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var(
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin:func(r *http.Request) bool{
			return true
		},
	}
)

type Userlogin struct {
	Type         string `json:"type"`
	Uid          string `json:"uid"`
	Usernme      string `json:"usernme"`
	Password     string `json:"password"`
}

type WebSocketController struct {

	WsConn *websocket.Conn
	Err error
	Conn *impl.Connection
	Datas []byte

	beego.Controller
}




// @Title ws websockt
// @Description websocket  接口 请求地址 ws://192.168.1.223:8080/v1/ws 第一次发送数据 {"type":"login","uid":"Mdsdf12323?sdf2222rerses","usernme":"ztg123","password":"yEQfTMX9xqt66zZdWOtc"}
// @Success 1 {} { ogin ok}
// @router / [get]
func (c *WebSocketController) WsHandler(){


	// 完成ws协议的握手操作
	// Upgrade:websocket
	if c.WsConn , c.Err = upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request,nil); c.Err != nil{
		return
	}

	if c.Conn , c.Err = impl.InitConnection(c.WsConn); c.Err != nil{

		c.Conn.Close()
		goto ERR
	}


	for {
		if c.Datas , c.Err = c.Conn.ReadMessage();c.Err != nil{

			goto ERR
		}
		if len(c.Datas)==0{
			if c.Err = c.Conn.WriteMessage([]byte("hello ok"));c.Err != nil{

				return
			}
		}

		//把数据格式化
		p := &Userlogin{}
		errs := json.Unmarshal([]byte(c.Datas), p)
		if errs !=nil{

			goto ERR         //数据格式错误 就马上干掉 客户端链接
		}

		//得到数据做相应的判断

		if p.Usernme !="ztg123"{               //用户名错误
			if c.Err = c.Conn.WriteMessage([]byte("username err "));c.Err != nil{
				goto ERR
			}
		}

		if c.Err = c.Conn.WriteMessage([]byte("login ok "));c.Err != nil{

			goto ERR
		}
		//如果想通过这里做其他事情可以用 switch  分支执行

		//扫描
		go c.Broadcast(p.Uid)      //充币提示
		go c.BroadcastOut(p.Uid)    //提币成功提示

	}

ERR:
	c.Conn.Close()

}



//广播 推送消息只是推送当前用户(充币提醒)
func (c *WebSocketController) Broadcast(uid string)  {

	for{
		 data,err:=models.TradeInList("0",uid)
		 if err !=nil{
			 utils.LogError(fmt.Sprintf("查询区块记录失败:%s\n",err))//日志
		 	break
		 }



		 //把数据加入到data里面去

		 if len(data)>0{



		 	//数据重组
			b, err := json.Marshal(data)
			if err !=nil{
				break
			}
			 map1:= []byte(`{"type":"in","uid":"`+uid+`","data":`+string(b)+`}`)
			 if err!=nil{
				 break
			 }
			 if c.Err = c.Conn.WriteMessage(map1);c.Err != nil{     //推送出去
				 c.Conn.Close()
			 }


		 }
		time.Sleep(5*time.Second)
	}
}


//广播 推送消息只是推送当前用户(转出成功提醒)
func (c *WebSocketController) BroadcastOut(uid string)  {

	for{
		data1,err:=models.TradeOutList("1",uid)
		if err !=nil{
			utils.LogError(fmt.Sprintf("查询区块记录失败:%s\n",err))//日志
			break
		}
		if len(data1)>0{

			b1, err := json.Marshal(data1)
			if err !=nil{
				break
			}


			str1:= []byte(`{"type":"out","uid":"`+uid+`","data":`+string(b1)+`}`)
			if c.Err = c.Conn.WriteMessage(str1);c.Err != nil{     //推送出去

				c.Conn.Close()
			}
		}
		time.Sleep(5*time.Second)
	}
}