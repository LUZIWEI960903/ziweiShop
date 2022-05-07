package mysql

import (
	"math"
	"ziweiShop/models"
)

// GetGoodsByCateId 根据cateId 查找对应的商品列表  --- goods 表
func GetGoodsByCateId(cateId, page, pageSize int) (goodsList []models.Goods, pageNum float64, err error) {
	// 判断是否为顶级分类
	goodsCate := models.GoodsCate{}
	err1 := db.Where("is_deleted=0 AND id=?", cateId).First(&goodsCate).Error
	if err1 != nil {
		return nil, -1, err1
	}
	var cateIds []int
	cateIds = append(cateIds, cateId)
	if goodsCate.Pid == 0 { // 顶级分类，需要查询其子分类的ids
		goodsCateList := []models.GoodsCate{}
		db.Where("is_deleted=0 AND pid=?", cateId).Find(&goodsCateList)
		for _, v := range goodsCateList {
			cateIds = append(cateIds, v.Id)
		}
	}
	where := "is_deleted=0 AND cate_id IN ?"
	err = db.Where(where, cateIds).Select("id,title,sub_title,click_count,price,goods_number,market_price,goods_img").Offset(pageSize * (page - 1)).Limit(pageSize).Find(&goodsList).Error
	if err != nil {
		return nil, -1, err
	}
	// 分页
	var count int64
	db.Table("goods").Where(where, cateIds).Count(&count)
	pageNum = math.Ceil(float64(count) / float64(pageSize))
	return goodsList, pageNum, nil
}

// GetGoodsInfoById  根据 goodsId 查找 商品信息  --- goods 表
func GetGoodsInfoById(goodsId int) (goodsInfo models.Goods, err error) {
	return goodsInfo, db.Where("is_deleted=0 AND id=?", goodsId).First(&goodsInfo).Error
}

// GetRelationGoodsList 根据 relationGoodsIds 获取 商品列表  --- goods 表
func GetRelationGoodsList(relationGoodsIds []string) (relationGoodsList []models.Goods, err error) {
	return relationGoodsList, db.Where("is_deleted=0 AND id IN ?", relationGoodsIds).Select("id,title,price,goods_version").Find(&relationGoodsList).Error
}

// GetGoodsGiftList 根据 goodsGiftIds 获取 商品赠品列表  --- goods 表
func GetGoodsGiftList(goodsGiftIds []string) (goodsGiftList []models.Goods, err error) {
	return goodsGiftList, db.Where("is_deleted=0 AND id IN ?", goodsGiftIds).Select("id,title,price,goods_version").Find(&goodsGiftList).Error
}

// GetGoodsColorList1 根据 goodsColorIds 获取 商品颜色列表  --- goods_color 表
func GetGoodsColorList1(goodsColorIds []string) (goodsColorList []models.GoodsColor, err error) {
	return goodsColorList, db.Where("is_deleted=0 AND id IN ?", goodsColorIds).Select("id,color_name,color_value").Find(&goodsColorList).Error
}

// GetGoodsFittingList 根据 goodsFittingIds 获取 商品配件列表  --- goods 表
func GetGoodsFittingList(goodsFittingIds []string) (goodsFittingList []models.Goods, err error) {
	return goodsFittingList, db.Where("is_deleted=0 AND id IN ?", goodsFittingIds).Select("id,title,price,goods_version").Find(&goodsFittingList).Error
}

// GetGoodsImageList 根据 goodsId 获取 商品相关图片的列表  --- goods_image 表
func GetGoodsImageList(goodsId int) (goodsImageList []models.GoodsImage, err error) {
	return goodsImageList, db.Where("is_deleted=0 AND goods_id = ?", goodsId).Select("id,color_id,img_url").Find(&goodsImageList).Error
}

// GetGoodsAttrList 根据 goodsId 获取 商品相关规格包装的列表  --- goods_attr 表
func GetGoodsAttrList(goodsId int) (goodsAttrList []models.GoodsAttr, err error) {
	return goodsAttrList, db.Where("is_deleted=0 AND goods_id = ?", goodsId).Select("id,attribute_cate_id,attribute_id,attribute_type,attribute_title,attribute_value").Find(&goodsAttrList).Error
}
