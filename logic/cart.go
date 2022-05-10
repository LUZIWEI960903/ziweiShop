package logic

import (
	"ziweiShop/dao/mysql"
	"ziweiShop/models"
	"ziweiShop/pkg/cookie"

	"github.com/gin-gonic/gin"
)

type CartLogic struct {
}

func (CartLogic) GetCart(c *gin.Context) *models.CartData {
	cartList := []models.Cart{}
	cookie.Cookie.Get(c, "cartList", &cartList)
	return &models.CartData{
		CartList: cartList,
	}
}

func (CartLogic) AddCart(c *gin.Context, goodsId, colorId int) error {
	// 查询商品信息
	goodsInfo, err1 := mysql.GetGoodsById2(goodsId)
	if err1 != nil {
		return err1
	}

	// 查看商品颜色信息
	goodsColor, err2 := mysql.GetGoodsColorById2(colorId)
	if err2 != nil {
		return err2
	}

	// 购物车目前的数据
	currentData := models.Cart{
		Id:           goodsInfo.Id,
		Title:        goodsInfo.Title,
		Price:        goodsInfo.Price,
		GoodsVersion: goodsInfo.GoodsVersion,
		Uid:          0,
		Num:          1,
		GoodsGift:    goodsInfo.GoodsGift,
		GoodsFitting: "",
		GoodsColor:   goodsColor.ColorName,
		GoodsImg:     goodsInfo.GoodsImg,
		GoodsAttr:    "",
		Checked:      true,
	}

	// 判断购物车有没有数据
	cartList := []models.Cart{}
	cookie.Cookie.Get(c, "cartList", &cartList)
	if len(cartList) > 0 { // 购物车有数据
		// 判断当前购物车是否有该数据
		if HasCartData(cartList, currentData) {
			for i, l := 0, len(cartList); i < l; i++ {
				if cartList[i].Id == currentData.Id && cartList[i].GoodsColor == currentData.GoodsColor && cartList[i].GoodsAttr == currentData.GoodsAttr {
					cartList[i].Num += 1
				}
			}
			currentData.Num += 1
		} else {
			cartList = append(cartList, currentData)
		}
		cookie.Cookie.Set(c, "cartList", &cartList)
	} else { // 购物车没数据，把当前数据写入cookie
		cartList = append(cartList, currentData)
		cookie.Cookie.Set(c, "cartList", &cartList)
	}
	return nil
}

// HasCartData 判断当前购物车是否有该数据
func HasCartData(cartList []models.Cart, currentData models.Cart) bool {
	for i, l := 0, len(cartList); i < l; i++ {
		return cartList[i].Id == currentData.Id && cartList[i].GoodsColor == currentData.GoodsColor && cartList[i].GoodsAttr == currentData.GoodsAttr
	}
	return false
}
