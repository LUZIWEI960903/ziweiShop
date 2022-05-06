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
