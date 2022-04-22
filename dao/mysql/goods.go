package mysql

import "ziweiShop/models"

// AddGoods 创建商品  --- goods 表
func AddGoods(p *models.AddGoodsParams) (goods models.Goods, err error) {
	goods = models.Goods{
		CateId:        p.CateId,
		ClickCount:    0,
		GoodsNumber:   p.GoodsNumber,
		IsHot:         p.IsHot,
		IsBest:        p.IsBest,
		IsNew:         p.IsNew,
		GoodsTypeId:   p.GoodsTypeId,
		Sort:          p.Sort,
		Status:        p.Status,
		AddTime:       p.AddTime,
		Price:         p.Price,
		MarketPrice:   p.MarketPrice,
		Title:         p.Title,
		SubTitle:      p.SubTitle,
		GoodsSn:       p.GoodsSn,
		RelationGoods: p.RelationGoods,
		GoodsAttr:     p.GoodsAttr,
		GoodsVersion:  p.GoodsVersion,
		GoodsImg:      p.GoodsImg,
		GoodsGift:     p.GoodsGift,
		GoodsFitting:  p.GoodsFitting,
		GoodsColor:    p.GoodsColor,
		GoodsKeywords: p.GoodsKeywords,
		GoodsDesc:     p.GoodsDesc,
		GoodsContent:  p.GoodsContent,
	}
	return goods, db.Create(&goods).Error
}

// GetGoodsList 查询所有 goods   --- goods 表
func GetGoodsList() (oGoodsList []models.Goods, err error) {
	oGoodsList = []models.Goods{}
	err = db.Where("is_deleted=0").Find(&oGoodsList).Error
	return oGoodsList, err
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
