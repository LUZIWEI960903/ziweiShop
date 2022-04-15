package logic

import (
	"ziweiShop/dao/mysql"
	"ziweiShop/models"
)

type GoodsTypeAttributeLogic struct {
}

func (GoodsTypeAttributeLogic) ShowAddPageLogic(cateId int) (data *models.AddGoodsTypeAttributePageData, err error) {
	// 查询 cateId 是否存在
	_, err = mysql.GetGoodsCateById(cateId)
	if err != nil {
		return nil, err
	}

	// 查询所有 商品类型
	oGoodsTypeList, err := mysql.GetGoodsTypeList()
	if err != nil {
		return nil, err
	}

	// 构造返回数据
	GoodsTypeItems := make([]models.GoodsTypeInfo, 0)
	GoodsTypeItems1 := make([]models.GoodsTypeInfo, 0)
	GoodsTypeItems2 := make([]models.GoodsTypeInfo, 0)
	for _, oGoodsType := range oGoodsTypeList {
		GoodsType := models.GoodsTypeInfo{
			Id:          oGoodsType.Id,
			Status:      oGoodsType.Status,
			Title:       oGoodsType.Title,
			Description: oGoodsType.Description,
		}
		if oGoodsType.Id == cateId {
			GoodsTypeItems1 = append(GoodsTypeItems1, GoodsType)
		} else {
			GoodsTypeItems2 = append(GoodsTypeItems2, GoodsType)
		}
	}
	GoodsTypeItems = append(GoodsTypeItems1, GoodsTypeItems2...)

	return &models.AddGoodsTypeAttributePageData{
		CateId:         cateId,
		GoodsTypeItems: GoodsTypeItems,
	}, nil
}

func (GoodsTypeAttributeLogic) DoAddLogic(p *models.AddGoodsTypeAttributeParams) (err error) {
	// 查看 cateId 是否存在
	_, err = mysql.GetGoodsTypeById(p.CateId)
	if err != nil {
		return err
	}
	return mysql.AddGoodsTypeAttribute(p)
}
