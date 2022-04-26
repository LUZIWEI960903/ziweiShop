package mysql

import "ziweiShop/models"

// AddGoodsImage 创建 goodsImage   --- goods_image 表
func AddGoodsImage(goodsImageObj *models.GoodsImage) error {
	return db.Create(goodsImageObj).Error
}

// GetGoodsImageListByGoodsId  通过 goodsId 查询 其所有的 GoodsImage  --- goods_image 表
func GetGoodsImageListByGoodsId(goodsId int) (oGoodsImageList []models.GoodsImage, err error) {
	oGoodsImageList = []models.GoodsImage{}
	err = db.Where("goods_id=? AND is_deleted=0", goodsId).Find(&oGoodsImageList).Error
	return oGoodsImageList, err
}

// EditColorByGoodsImageId  根据 goodsImageId 修改 colorId  --- goods_image 表
func EditColorByGoodsImageId(goodsImageId, colorId int) (err error) {
	goodsImageList := []models.GoodsImage{}
	err = db.Where("id=? AND is_deleted=0", goodsImageId).First(&goodsImageList).Error
	if len(goodsImageList) < 1 || err != nil {
		return ErrGetGoodsImage
	}
	goodsImageList[0].ColorId = colorId
	return db.Save(&goodsImageList).Error
}

// AjaxRemoveGoodsImageByGoodsImageId 根据 goodsImageId 删除 商品图库信息  --- goods_image 表
func AjaxRemoveGoodsImageByGoodsImageId(goodsImageId int) (imgSrc string, err error) {
	goodsImageList := []models.GoodsImage{}
	err = db.Where("id=? AND is_deleted=0", goodsImageId).First(&goodsImageList).Error
	if len(goodsImageList) < 1 || err != nil {
		return "", ErrGetGoodsImage
	}
	goodsImageList[0].IsDeleted = 1
	return goodsImageList[0].ImgUrl, db.Save(&goodsImageList).Error
}
