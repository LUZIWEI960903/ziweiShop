package logic

import (
	"ziweiShop/models"
	"ziweiShop/pkg/cookie"

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

	return &models.PassCheckoutData{
		OrderList:    orderList,
		TotalPrice:   totalPrice,
		ShopBaseData: baseData,
	}, true
}
