package logic

import (
	"fmt"
	"time"
	"ziweiShop/config"
	"ziweiShop/models"

	"github.com/smartwalle/alipay/v3"

	"github.com/gin-gonic/gin"
)

type AlipayLogic struct {
	BaseLogic
}

var (
	privateKey = "xxx" // 必须，上一步中使用 RSA签名验签工具 生成的私钥
	appId      = "xxx"
)

func (l AlipayLogic) Alipay(c *gin.Context, orderId int) (*models.AlipayData, error) {
	// 校验该订单是否属于该用户
	order, err1 := l.verifyOrder(c, orderId)
	if err1 != nil {
		return nil, err1
	}

	// 调用支付宝接口
	var client, err2 = alipay.New(appId, privateKey, true)

	client.LoadAppPublicCertFromFile("crt/appCertPublicKey_2017011104995404.crt") // 加载应用公钥证书
	client.LoadAliPayRootCertFromFile("crt/alipayRootCert.crt")                   // 加载支付宝根证书
	client.LoadAliPayPublicCertFromFile("crt/alipayCertPublicKey_RSA2.crt")       // 加载支付宝公钥证书

	// 将 key 的验证调整到初始化阶段
	if err2 != nil {
		fmt.Println(err1)
		return nil, err2
	}

	var p = alipay.TradePagePay{}
	p.NotifyURL = fmt.Sprintf("%v%v/v3/alipayNotify", config.Conf.Host, config.Conf.Port)
	p.ReturnURL = fmt.Sprintf("%v%v/v3/alipayReturn", config.Conf.Host, config.Conf.Port)
	p.Subject = order.Name
	template := "2006-01-02 15:04:05"
	p.OutTradeNo = time.Now().Format(template) // 传递一个唯一单号
	p.TotalAmount = fmt.Sprintf("%v", order.AllPrice)
	p.ProductCode = "FAST_INSTANT_TRADE_PAY" // https://opendocs.alipay.com/open/028r8t?scene=22

	var url, err4 = client.TradePagePay(p)
	if err4 != nil {
		fmt.Println(err4)
	}

	var payURL = url.String()
	fmt.Println(payURL)
	// 这个 payURL 即是用于支付的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	return &models.AlipayData{
		PayUrl: payURL,
	}, nil
}

func (AlipayLogic) AlipayNotify(c *gin.Context) error {
	var client, err1 = alipay.New(appId, privateKey, true)

	client.LoadAppPublicCertFromFile("crt/appCertPublicKey_2017011104995404.crt") // 加载应用公钥证书
	client.LoadAliPayRootCertFromFile("crt/alipayRootCert.crt")                   // 加载支付宝根证书
	client.LoadAliPayPublicCertFromFile("crt/alipayCertPublicKey_RSA2.crt")       // 加载支付宝公钥证书

	// 将 key 的验证调整到初始化阶段
	if err1 != nil {
		fmt.Println(err1)
		return err1
	}

	req := c.Request
	req.ParseForm()
	ok, err := client.VerifySign(req.Form)
	fmt.Println(ok, err)
	return err
}
