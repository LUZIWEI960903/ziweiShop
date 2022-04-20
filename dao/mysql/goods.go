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
