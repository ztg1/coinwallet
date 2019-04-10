package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {
    beego.GlobalControllerRouter["apisdrm/controllers:UserController"] = append(beego.GlobalControllerRouter["apisdrm/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apisdrm/controllers:UserController"] = append(beego.GlobalControllerRouter["apisdrm/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})



    beego.GlobalControllerRouter["apisdrm/controllers:UserController"] = append(beego.GlobalControllerRouter["apisdrm/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apisdrm/controllers:UserController"] = append(beego.GlobalControllerRouter["apisdrm/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apisdrm/controllers:UserController"] = append(beego.GlobalControllerRouter["apisdrm/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apisdrm/controllers:UserController"] = append(beego.GlobalControllerRouter["apisdrm/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apisdrm/controllers:UserController"] = append(beego.GlobalControllerRouter["apisdrm/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

//获取所有的币种
    beego.GlobalControllerRouter["apisdrm/controllers:EthCoinController"] = append(beego.GlobalControllerRouter["apisdrm/controllers:EthCoinController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

// 创建私钥
    beego.GlobalControllerRouter["apisdrm/controllers:EthCoinController"] = append(beego.GlobalControllerRouter["apisdrm/controllers:EthCoinController"],
        beego.ControllerComments{
            Method: "PrivateKey",
            Router: `/privatekey`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    //创建账户
    beego.GlobalControllerRouter["apisdrm/controllers:EthCoinController"] = append(beego.GlobalControllerRouter["apisdrm/controllers:EthCoinController"],
            beego.ControllerComments{
                Method: "CreatAccount",
                Router: `/creataccount`,
                AllowHTTPMethods: []string{"post"},
                MethodParams: param.Make(),
                Filters: nil,
                Params: nil})

    //账户余额
    beego.GlobalControllerRouter["apisdrm/controllers:EthCoinController"] = append(beego.GlobalControllerRouter["apisdrm/controllers:EthCoinController"],
                beego.ControllerComments{
                    Method: "AccountBalance",
                    Router: `/accountbalance`,
                    AllowHTTPMethods: []string{"post"},
                    MethodParams: param.Make(),
                    Filters: nil,
                    Params: nil})

    //
    beego.GlobalControllerRouter["apisdrm/controllers:EthCoinController"] = append(beego.GlobalControllerRouter["apisdrm/controllers:EthCoinController"],
                    beego.ControllerComments{
                        Method: "RecordAccount",
                        Router: `/recordaccount`,
                        AllowHTTPMethods: []string{"post"},
                        MethodParams: param.Make(),
                        Filters: nil,
                        Params: nil})

    beego.GlobalControllerRouter["apisdrm/controllers:EthCoinController"] = append(beego.GlobalControllerRouter["apisdrm/controllers:EthCoinController"],
                        beego.ControllerComments{
                            Method: "TradeOut",
                            Router: `/tradeout`,
                            AllowHTTPMethods: []string{"post"},
                            MethodParams: param.Make(),
                            Filters: nil,
                            Params: nil})
    beego.GlobalControllerRouter["apisdrm/controllers:EthCoinController"] = append(beego.GlobalControllerRouter["apisdrm/controllers:EthCoinController"],
                        beego.ControllerComments{
                            Method: "SetReturnIn",
                            Router: `/setreturnin`,
                            AllowHTTPMethods: []string{"post"},
                            MethodParams: param.Make(),
                            Filters: nil,
                            Params: nil})

    beego.GlobalControllerRouter["apisdrm/controllers:EthCoinController"] = append(beego.GlobalControllerRouter["apisdrm/controllers:EthCoinController"],
                            beego.ControllerComments{
                                Method: "SetReturnOut",
                                Router: `/setreturnout`,
                                AllowHTTPMethods: []string{"post"},
                                MethodParams: param.Make(),
                                Filters: nil,
                                Params: nil})




    //ws

    beego.GlobalControllerRouter["apisdrm/controllers:WebSocketController"] = append(beego.GlobalControllerRouter["apisdrm/controllers:WebSocketController"],
        beego.ControllerComments{
            Method: "WsHandler",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})
}
