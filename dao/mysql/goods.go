package mysql

import (
	"math"
	"ziweiShop/models"
	"ziweiShop/pkg/tools"
)

// AddGoods 创建商品  --- goods 表
func AddGoods(p *models.AddGoodsParams) (goods *models.Goods, err error) {
	goods = &models.Goods{}
	goods.CateId = p.CateId
	goods.ClickCount = 0
	goods.GoodsNumber = p.GoodsNumber
	goods.IsHot = p.IsHot
	goods.IsBest = p.IsBest
	goods.IsNew = p.IsNew
	goods.GoodsTypeId = p.GoodsTypeId
	goods.Sort = p.Sort
	goods.Status = p.Status
	goods.AddTime = int(tools.GetUnix())
	goods.Price = p.Price
	goods.MarketPrice = p.MarketPrice
	goods.Title = p.Title
	goods.SubTitle = p.SubTitle
	goods.GoodsSn = p.GoodsSn
	goods.RelationGoods = p.RelationGoods
	goods.GoodsAttr = p.GoodsAttr
	goods.GoodsVersion = p.GoodsVersion
	goods.GoodsImg = p.GoodsImg
	goods.GoodsGift = p.GoodsGift
	goods.GoodsFitting = p.GoodsFitting
	goods.GoodsColor = p.GoodsColor
	goods.GoodsKeywords = p.GoodsKeywords
	goods.GoodsDesc = p.GoodsDesc
	goods.GoodsContent = p.GoodsContent

	return goods, db.Create(&goods).Error
}

// GetGoodsListByPage 分页查询 GoodsList --- goods 表
func GetGoodsListByPage(page, pageSize int) (oGoodsList []models.Goods, pageCount float64, err error) {
	oGoodsList = []models.Goods{}
	err = db.Where("is_deleted=0").Offset(pageSize * (page - 1)).Limit(pageSize).Find(&oGoodsList).Error

	var count int64
	// 获取goods总条数
	db.Table("goods").Where("is_deleted=0").Count(&count)
	// 获取总页数
	pageCount = math.Ceil(float64(count) / float64(pageSize))
	return oGoodsList, pageCount, err
}

// GetGoodsById  // 根据 goodId 查询 商品详情   --- goods 表
func GetGoodsById(goodsId int) (oGoodsInfo *models.Goods, err error) {
	goodsList := []models.Goods{}
	err = db.Where("id=? AND is_deleted=0", goodsId).First(&goodsList).Error
	if len(goodsList) < 1 || err != nil {
		return nil, ErrGetGoods
	}
	return &goodsList[0], nil
}

// EditGoods  修改 goods 信息   --- goods 表
func EditGoods(p *models.EditGoodsParams) (err error) {
	goodsList := []models.Goods{}
	err = db.Where("id=? AND is_deleted=0", p.Id).First(&goodsList).Error
	if len(goodsList) < 1 || err != nil {
		return ErrGetGoods
	}

	goodsList[0].CateId = p.CateId
	goodsList[0].GoodsNumber = p.GoodsNumber
	goodsList[0].IsHot = p.IsHot
	goodsList[0].IsBest = p.IsBest
	goodsList[0].IsNew = p.IsNew
	goodsList[0].GoodsTypeId = p.GoodsTypeId
	goodsList[0].Sort = p.Sort
	goodsList[0].Status = p.Status
	goodsList[0].Price = p.Price
	goodsList[0].MarketPrice = p.MarketPrice
	goodsList[0].Title = p.Title
	goodsList[0].SubTitle = p.SubTitle
	goodsList[0].GoodsSn = p.GoodsSn
	goodsList[0].RelationGoods = p.RelationGoods
	goodsList[0].GoodsAttr = p.GoodsAttr
	goodsList[0].GoodsVersion = p.GoodsVersion
	if p.GoodsImg != "" {
		goodsList[0].GoodsImg = p.GoodsImg
	}
	goodsList[0].GoodsGift = p.GoodsGift
	goodsList[0].GoodsFitting = p.GoodsFitting
	goodsList[0].GoodsColor = p.GoodsColor
	goodsList[0].GoodsKeywords = p.GoodsKeywords
	goodsList[0].GoodsDesc = p.GoodsDesc
	goodsList[0].GoodsContent = p.GoodsContent

	return db.Save(&goodsList).Error
}

// DeleteGoods  根据 goodsId 逻辑删除 goods  --- goods 表
func DeleteGoods(goodsId int) (imgSrc string, err error) {
	goodsList := []models.Goods{}
	err = db.Where("id=? AND is_deleted=0", goodsId).First(&goodsList).Error
	if len(goodsList) < 1 || err != nil {
		return "", ErrGetGoods
	}
	//fmt.Printf("goods:%#v\n", goodsList)
	goodsList[0].IsDeleted = 1
	return goodsList[0].GoodsImg, db.Save(&goodsList).Error
}
