package impl

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"regexp"
)
const  ETHURL string ="https://mainnet.infura.io"

func init()  {


	client, err := ethclient.Dial(ETHURL)
	defer client.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("链接类型=%T",client)
    _=new(Eth)
	_ = client // we'll use this in the upcoming sections
}
type Eth struct {

}

/**
    创建钱包,私钥 ,公钥，地址
 */
func (e*Eth) StatusKey() (string ,error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
		return "",err
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privkey:=hexutil.Encode(privateKeyBytes)[2:]


	fmt.Println("私钥=",privkey)
//公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("公钥=",hexutil.Encode(publicKeyBytes)[4:]) // 9a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05


	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("地址=：",address) // 0x96216849c49358B10257cb55b28eA603c874b05E

	return privkey,nil
}

/*
 创建加密钱包
*/
func (e * Eth) NewAccount(password string)(string,error) {

	ks := keystore.NewKeyStore("/home/wwwroot/default/wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.URL) // 0x20F8D42FB0F667F2E53930fed426f225752453b3
	fmt.Println(account.URL.Path) // 0x20F8D42FB0F667F2E53930fed426f225752453b3

	//上传到Swarm 仓库
	return account.Address.Hex(),nil
}


/*
  读取 文件解码钱包 地址

*/
func (e *Eth)NewKeyStore()  {
	file := "/home/wwwroot/default/wallets/UTC--2019-03-27T08-46-59.108304496Z--c68c1e5c2663b0bd8019420701cf6e8645cc5d15"

	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("读取错误")

		log.Fatal(err)
	}

	password := "akdd"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("新的地址=",account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3



	key, err :=keystore.DecryptKey(jsonBytes,password)
	address := key.Address.Hex()
	privateKey := hex.EncodeToString(crypto.FromECDSA(key.PrivateKey))

	fmt.Printf("Address:\t%s\nPrivateKey:\t%s\n",
		address,
		privateKey,
	)
	/*if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}*/
}

/*
 地址检查
*/
func (c * Eth)AddressCheck(address string)(bool) {
	 re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(address)
}


/*
  获取账户最新的余额
*/
func (e *Eth)AccountBalance(address string) (balanc *big.Float,err error)  {
     //0x93c96789bC0679e4897C55EE16472a18B2c70908
	client, err := ethclient.Dial(ETHURL)
	defer client.Close()
	if err != nil {
		log.Fatal(err)

		return nil,err
	}
	account := common.HexToAddress(address)

	//balance, err := client.BalanceAt(context.Background(), account, nil)
	balance, err := client.BalanceAt(context.Background(), account, nil)

	if err != nil {
		log.Fatal(err)
		return nil,err
	}
	//fmt.Printf("add=%d",balance.Int64()) // 25893180161173005034

   //处理数据
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Println(ethValue) // 25.729324269165216041
	fmt.Printf("类型=%T",ethValue)

	return ethValue,nil
}


/*
  转账
*/
func (e * Eth) TradeAccount(signed string)(txs string ,err error) {
	client, err := ethclient.Dial(ETHURL)
	defer client.Close()
	if err != nil {
		//log.Fatal(err)
		return "",err
	}

	//私钥
	privatekey:="27225580f1fc750f31a2846468c75f7913c4b00ddcb7e66a647f2945255a0bc7"
	privateKey, err := crypto.HexToECDSA(privatekey)
	if err != nil {
		//log.Fatal(err)
		return "" ,err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		//log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return "",errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		//log.Fatal(err)
		return "",err
	}

	value := big.NewInt(10000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		//log.Fatal(err)
		return "",err
	}

	toAddress := common.HexToAddress("0xc68C1e5c2663b0Bd8019420701cF6E8645cC5D15")  // 转 币
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		//log.Fatal(err)
		return "",err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		//log.Fatal(err)
		return "",err
	}

	//fmt.Println("signedTx=",signedTx)

	//来将已签名的事务广播到整个网络。
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		//log.Fatal(err)
		return "",err
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
	fmt.Printf("tx 类型: %T", signedTx.Hash().Hex())

     return signedTx.Hash().Hex(),nil
}


//查询交易
func (e *Eth)TransactionCount(Tx string)(txs string,prep bool,err error)  {

	client, err := ethclient.Dial(ETHURL)
	defer client.Close()
	if err != nil {
		//log.Fatal(err)
		return "",false,err
	}

	txHash := common.HexToHash("0x612a1fcafee57feb202974f01101bf5ac86897985c5b7d714a1be9b787a21754")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("txdata=%s\n",tx.To())

	return "1111",isPending,nil


}


//监听新区块
func (e* Eth) ListenBlock()(*ethclient.Client,error) {

	client, err := ethclient.Dial(ETHURL)
	if err != nil {

		return nil,err
	}
	return client,nil
}




