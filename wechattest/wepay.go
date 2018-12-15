package main

import (
	"fmt"
	"net/http"
	"github.com/objcoding/wxpay"
)

const (
	wxAppId           = "ww5a14157e2dc2c153"                           // 填上自己的参数
	mchiId       = "1520839051"                       // 填上自己的参数
	apiKey = "7iRv80sZNKuwIgQGsakHE42QMhNTfKSgU3wcvy1FjTa-sth16ZAVxEK0QBuZhQ6K" // 填上自己的参数
)


func init() {
	http.HandleFunc("/", Page1Handler)
}

// 建立必要的 session, 然后跳转到授权页面
func Page1Handler(w http.ResponseWriter, r *http.Request) {
	// 创建支付账户
	account := wxpay.NewAccount(wxAppId, mchiId, apiKey, false)

	// 新建微信支付客户端
	client := wxpay.NewClient(account)

	// 设置证书
	//account.SetCertData("证书地址")


	// 设置http请求超时时间
	client.SetHttpConnectTimeoutMs(2000)

	// 设置http读取信息流超时时间
	client.SetHttpReadTimeoutMs(1000)

	// 更改签名类型
	client.SetSignType(wxpay.HMACSHA256)

	// 统一下单
	params := make(wxpay.Params)
	params.SetString("body", "test").
		SetString("out_trade_no", "436577857").
		SetInt64("total_fee", 1).
		SetString("spbill_create_ip", "127.0.0.1").
		SetString("notify_url", "http://notify.objcoding.com/notify").
		SetString("trade_type", "NATIVE")
	p, err := client.UnifiedOrder(params)
	fmt.Println("%v %v", p, err)
}


func main() {
	fmt.Println(http.ListenAndServe(":8899", nil))
}