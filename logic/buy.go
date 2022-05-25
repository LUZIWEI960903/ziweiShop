package logic

import (
	"strconv"
	"ziweiShop/dao/mysql"
	"ziweiShop/models"
	"ziweiShop/pkg/cookie"
	"ziweiShop/pkg/tools"

	"github.com/gin-gonic/gin"
)

type BuyLogic struct {
	BaseLogic
}

func (l BuyLogic) Checkout(c *gin.Context) (*models.PassCheckoutData, bool) {
	// 获取基础信息
	baseData, err := l.getBaseData(c)
	if err != nil {
		return nil, false
	}

	// 获取购物车信息
	cartList := []models.Cart{}
	if !cookie.Cookie.Get(c, "cartList", &cartList) {
		return nil, false
	}

	orderList := []models.Cart{}
	var totalPrice float64
	if len(cartList) > 0 {
		for i, l := 0, len(cartList); i < l; i++ {
			if cartList[i].Checked {
				orderList = append(orderList, cartList[i])
				totalPrice += cartList[i].Price * float64(cartList[i].Num)
			}
		}
	}

	// 获取收货地址
	user := models.User{}
	if !cookie.Cookie.Get(c, "userInfo", &user) {
		return nil, false
	}

	addressList, err2 := mysql.GetAddressByUid(user.Id)
	if err2 != nil {
		return nil, false
	}

	return &models.PassCheckoutData{
		OrderList:    orderList,
		TotalPrice:   totalPrice,
		ShopBaseData: baseData,
		AddressList:  addressList,
	}, true
}

func (BuyLogic) DoCheckout(c *gin.Context) error {
	// 获取用户信息
	user := models.User{}
	cookie.Cookie.Get(c, "userInfo", &user)

	// 获取收货地址
	addressInfo, err1 := mysql.GetDefaultAddressByUid(user.Id)
	if err1 != nil {
		return err1
	}

	// 获取购物信息 + 创建结算清单
	cartList := []models.Cart{}
	cookie.Cookie.Get(c, "cartList", &cartList)

	var totalPrice float64
	orderList := []models.Cart{}
	if len(cartList) > 0 {
		for i, l := 0, len(cartList); i < l; i++ {
			if cartList[i].Checked {
				totalPrice += cartList[i].Price * float64(cartList[i].Num)
				orderList = append(orderList, cartList[i])
			}
		}
	}

	// 把订单信息放进订单表（order）里面，把商品信息放进商品表（order_item）
	order := models.Order{
		OrderId:     strconv.Itoa(int(tools.GetUnixN())),
		Uid:         user.Id,
		AllPrice:    totalPrice,
		Phone:       addressInfo.Phone,
		Name:        addressInfo.Name,
		Address:     addressInfo.Address,
		PayStatus:   0,
		PayType:     0,
		OrderStatus: 0,
		AddTime:     int(tools.GetUnix()),
	}
	orderInfo, err2 := mysql.CreateOrder(&order)
	if err2 != nil {
		return err2
	}

	for i, l := 0, len(orderList); i < l; i++ {
		orderItem := models.OrderItem{
			OrderId:      orderInfo.Id,
			Uid:          user.Id,
			ProductId:    orderList[i].Id,
			ProductNum:   orderList[i].Num,
			AddTime:      int(tools.GetUnix()),
			ProductTitle: orderList[i].Title,
			GoodsVersion: orderList[i].GoodsVersion,
			GoodsColor:   orderList[i].GoodsColor,
			ProductImg:   orderList[i].GoodsImg,
			ProductPrice: orderList[i].Price,
		}
		mysql.CreateOrderItem(&orderItem)
	}

	// 删除购物车选中的数据
	newCartList := make([]models.Cart, 0)
	for i, l := 0, len(cartList); i < l; i++ {
		if !cartList[i].Checked {
			newCartList = append(newCartList, cartList[i])
		}
	}
	cookie.Cookie.Set(c, "cartList", &newCartList)
	return nil
}
