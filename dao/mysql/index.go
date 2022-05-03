package mysql

import "ziweiShop/models"

// GetTopNavList 获取顶部导航栏列表  --- nav 表
func GetTopNavList() (oTopNavList []models.Nav, err error) {
	return oTopNavList, db.Where("is_deleted=0 AND position=1").Find(&oTopNavList).Error
}

// GetMiddleNavList 获取中间导航栏列表  --- nav 表
func GetMiddleNavList() (oMiddleNavList []models.Nav, err error) {
	return oMiddleNavList, db.Where("is_deleted=0 AND position=2").Find(&oMiddleNavList).Error
}

// GetGoodsListByIds 根据ids 获取对应 goods  --- goods 表
func GetGoodsListByIds(relationIds []string) (goodsList []models.Goods, err error) {
	return goodsList, db.Where("is_deleted=0 AND id IN ?", relationIds).Select("id,title,goods_img,price").Find(&goodsList).Error
}

// GetGoodsByCategory 根据 分类获取 goods  --- goods 表
func GetGoodsByCategory(cateId int, displayByWhat string, limitNum int) (goodsList []models.Goods) {
	// 判断 cate 是否为顶级分类 --- goods_cate 表
	goodsCate := models.GoodsCate{Id: cateId}
	db.First(&goodsCate)

	var goodsCateIds []int
	goodsCateIds = append(goodsCateIds, goodsCate.Id)
	if goodsCate.Pid == 0 { // 顶级分类
		// 获取顶级分类下面的二级分类
		goodsCateList := []models.GoodsCate{}
		db.Where("pid=? AND is_deleted=0", goodsCate.Id).Find(&goodsCateList)
		for _, v := range goodsCateList {
			goodsCateIds = append(goodsCateIds, v.Id)
		}
	}

	where := "cate_id in ?"
	switch displayByWhat {
	case "hot":
		where += " AND is_hot=1"
	case "best":
		where += " AND is_best=1"
	case "new":
		where += " AND is_new=1"
	default:
		where += ""
	}

	db.Where(where, goodsCateIds).Select("id,title,goods_img,price,click_count,sub_title").Limit(limitNum).Order("sort DESC").Find(&goodsList)
	return
}
