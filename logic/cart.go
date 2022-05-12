package logic

import (
	"fmt"
	"ziweiShop/dao/mysql"
	"ziweiShop/models"
	"ziweiShop/pkg/cookie"

	"github.com/gin-gonic/gin"
)

type CartLogic struct {
	BaseLogic
}

func (l CartLogic) GetCart(c *gin.Context) *models.CartData {
	// 获取基础信息
	baseData, _ := l.getBaseData()
	cartList := []models.Cart{}
	cookie.Cookie.Get(c, "cartList", &cartList)

	// 计算总消费
	var totalPrice float64
	for i, l := 0, len(cartList); i < l; i++ {
		if cartList[i].Checked {
			totalPrice += float64(cartList[i].Num) * cartList[i].Price
		}
	}
	return &models.CartData{
		CartList:     cartList,
		ShopBaseData: baseData,
		TotalPrice:   totalPrice,
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
	fmt.Println(cartList)
	if len(cartList) > 0 { // 购物车有数据
		// 判断当前购物车是否有该数据
		if HasCartData(cartList, currentData) {
			for i, l := 0, len(cartList); i < l; i++ {
				if cartList[i].Id == currentData.Id && cartList[i].GoodsColor == currentData.GoodsColor && cartList[i].GoodsAttr == currentData.GoodsAttr {
					cartList[i].Num += 1
				}
			}
		} else {
			cartList = append(cartList, currentData)
		}
		cookie.Cookie.Set(c, "cartList", &cartList)
	} else { // 购物车没数据，把当前数据写入cookie
		cartList = append(cartList, currentData)
		cookie.Cookie.Set(c, "cartList", &cartList)
	}
	fmt.Println(cartList)
	return nil
}

func (CartLogic) AddCartSuccess(goodsId int) (*models.AddCartSuccessData, error) {
	goodsInfo, err1 := mysql.GetGoodsById2(goodsId)
	if err1 != nil {
		return nil, err1
	}

	return &models.AddCartSuccessData{
		GoodsInfo: *goodsInfo,
	}, nil
}

func (CartLogic) IncCart(c *gin.Context, goodsId int, goodsColor string) *models.IncCartData {
	// 获取购物车信息
	cartList := []models.Cart{}
	cookie.Cookie.Get(c, "cartList", &cartList)
	var (
		num          int
		currentPrice float64
		TotalPrice   float64
	)
	if len(cartList) > 0 {
		for i, l := 0, len(cartList); i < l; i++ {
			if cartList[i].Id == goodsId && cartList[i].GoodsColor == goodsColor {
				cartList[i].Num += 1
				num = cartList[i].Num
				currentPrice += float64(cartList[i].Num) * cartList[i].Price
			}
			if cartList[i].Checked {
				TotalPrice += float64(cartList[i].Num) * cartList[i].Price
			}
		}
		cookie.Cookie.Set(c, "cartList", cartList)
	}
	return &models.IncCartData{
		Num:          num,
		CurrentPrice: currentPrice,
		TotalPrice:   TotalPrice,
	}

}

func (CartLogic) DecCart(c *gin.Context, goodsId int, goodsColor string) *models.DecCartData {
	// 获取当前购物车信息
	cartList := []models.Cart{}
	cookie.Cookie.Get(c, "cartList", &cartList)
	var (
		num          int
		currentPrice float64
		totalPrice   float64
	)
	if len(cartList) > 0 {
		for i, l := 0, len(cartList); i < l; i++ {
			if cartList[i].Id == goodsId && cartList[i].GoodsColor == goodsColor {
				if cartList[i].Num > 1 {
					cartList[i].Num -= 1
				}
				num += cartList[i].Num
				currentPrice += float64(cartList[i].Num) * cartList[i].Price

			}
			if cartList[i].Checked {
				totalPrice += float64(cartList[i].Num) * cartList[i].Price
			}
		}
		cookie.Cookie.Set(c, "cartList", cartList)
	}
	return &models.DecCartData{
		Num:          num,
		CurrentPrice: currentPrice,
		TotalPrice:   totalPrice,
	}
}

func (CartLogic) ChangeOneCart(c *gin.Context, goodsId int, goodsColor string) *models.ChangeOneCartData {
	// 获取当前购物车数据
	cartList := []models.Cart{}
	cookie.Cookie.Get(c, "cartList", &cartList)
	var totalPrice float64
	if len(cartList) > 0 {
		for i, l := 0, len(cartList); i < l; i++ {
			if cartList[i].Id == goodsId && cartList[i].GoodsColor == goodsColor {
				cartList[i].Checked = !cartList[i].Checked
			}
			if cartList[i].Checked {
				totalPrice += float64(cartList[i].Num) * cartList[i].Price
			}
		}
		cookie.Cookie.Set(c, "cartList", cartList)
	}
	return &models.ChangeOneCartData{
		CartList:   cartList,
		TotalPrice: totalPrice,
	}
}

func (CartLogic) ChangeAllCart(c *gin.Context, flag string) *models.ChangeAllCartData {
	// 获取购物车信息
	cartList := []models.Cart{}
	cookie.Cookie.Get(c, "cartList", &cartList)
	var totalPrice float64
	if len(cartList) > 0 {
		for i, l := 0, len(cartList); i < l; i++ {
			if flag == "1" {
				cartList[i].Checked = true
			} else {
				cartList[i].Checked = false
			}
			if cartList[i].Checked {
				totalPrice += float64(cartList[i].Num) * cartList[i].Price
			}
		}
		cookie.Cookie.Set(c, "cartList", cartList)
	}
	return &models.ChangeAllCartData{
		CartList:   cartList,
		TotalPrice: totalPrice,
	}
}

func (CartLogic) DeleteCart(c *gin.Context, goodsId int, goodsColor string) *models.DeleteCartData {
	// 获取当前购物车信息
	cartList := []models.Cart{}
	cookie.Cookie.Get(c, "cartList", &cartList)
	var totalPrice float64
	if len(cartList) > 0 {
		for i, l := 0, len(cartList); i < l; i++ {
			if cartList[i].Checked && !(cartList[i].Id == goodsId && cartList[i].GoodsColor == goodsColor) {
				totalPrice += float64(cartList[i].Num) * cartList[i].Price
			}
			if cartList[i].Id == goodsId && cartList[i].GoodsColor == goodsColor {
				cartList = append(cartList[:i], cartList[(i+1):]...)
				break
			}
		}
		cookie.Cookie.Set(c, "cartList", cartList)
	}

	return &models.DeleteCartData{
		CartList:   cartList,
		TotalPrice: totalPrice,
	}
}

// HasCartData 判断当前购物车是否有该数据
func HasCartData(cartList []models.Cart, currentData models.Cart) bool {
	for i, l := 0, len(cartList); i < l; i++ {
		return cartList[i].Id == currentData.Id && cartList[i].GoodsColor == currentData.GoodsColor && cartList[i].GoodsAttr == currentData.GoodsAttr
	}
	return false
}
