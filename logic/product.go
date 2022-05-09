package logic

import (
	"strings"
	"ziweiShop/dao/mysql"
	"ziweiShop/models"
)

type ProductLogic struct {
	BaseLogic
}

func (l ProductLogic) SearchProductsById(cateId, page, pageSize int) (*models.SearchProductsByKeywordData, error) {
	// 获取基础数据
	baseData, err1 := l.getBaseData()
	if err1 != nil {
		return nil, err1
	}

	// 根据cateId 获取商品
	goodsList, pageNum, err2 := mysql.GetGoodsByCateId(cateId, page, pageSize)
	if err2 != nil {
		return nil, err2
	}

	return &models.SearchProductsByKeywordData{
		ShopBaseData: baseData,
		GoodsList:    goodsList,
		PageNum:      pageNum,
		CurrentPage:  page,
	}, nil
}

func (l ProductLogic) GetGoodsInfoData(goodsId int) (*models.GoodsInforData, error) {
	// 获取基础数据
	baseData, err1 := l.getBaseData()
	if err1 != nil {
		return nil, err1
	}

	// 获取商品信息
	goodsInfo, err2 := mysql.GetGoodsInfoById(goodsId)
	if err2 != nil {
		return nil, err2
	}

	// 获取关联商品
	relationGoodsIdsStr := strings.ReplaceAll(goodsInfo.RelationGoods, "，", ",")
	relationGoodsIds := strings.Split(relationGoodsIdsStr, ",")
	relationGoodsList, err3 := mysql.GetRelationGoodsList(relationGoodsIds)
	if err3 != nil {
		return nil, err3
	}

	// 获取关联赠品
	goodsGiftIdsStr := strings.ReplaceAll(goodsInfo.GoodsGift, "，", ",")
	goodsGiftIds := strings.Split(goodsGiftIdsStr, ",")
	goodsGiftList, err4 := mysql.GetGoodsGiftList(goodsGiftIds)
	if err4 != nil {
		return nil, err4
	}

	// 获取关联颜色
	goodsColorIdsStr := strings.ReplaceAll(goodsInfo.GoodsColor, "，", ",")
	goodsColorIds := strings.Split(goodsColorIdsStr, ",")
	goodsColorList, err5 := mysql.GetGoodsColorList1(goodsColorIds)
	if err5 != nil {
		return nil, err5
	}

	// 获取关联配件
	goodsFittingIdsStr := strings.ReplaceAll(goodsInfo.GoodsFitting, "，", ",")
	goodsFittingIds := strings.Split(goodsFittingIdsStr, ",")
	goodsFittingList, err6 := mysql.GetGoodsFittingList(goodsFittingIds)
	if err6 != nil {
		return nil, err6
	}

	// 获取商品关联的图片
	goodsImageList, err7 := mysql.GetGoodsImageList(goodsId)
	if err7 != nil {
		return nil, err7
	}

	// 获取规格参数
	goodsAttrList, err8 := mysql.GetGoodsAttrList(goodsId)
	if err8 != nil {
		return nil, err8
	}

	// 获取更多属性
	/*
		颜色:xxx,xxx,xxx | 尺码:xxx,xxx,xxx
	*/
	goodsItemAttr := goodsInfo.GoodsAttr
	goodsItemAttr = strings.ReplaceAll(goodsItemAttr, " ", "")
	goodsItemAttr = strings.ReplaceAll(goodsItemAttr, "，", ",")
	goodsItemAttr = strings.ReplaceAll(goodsItemAttr, "：", ":")
	goodsItemAttrSlice := strings.Split(goodsItemAttr, "|")
	goodsItemAttrList := make([]models.GoodsItemAttr, 0)
	for _, goodsItemAttrStr := range goodsItemAttrSlice {
		if strings.Contains(goodsItemAttrStr, ":") {
			goodsItemAttrStrSlice := strings.Split(goodsItemAttrStr, ":")
			newGoodsItemAttr := models.GoodsItemAttr{
				Cate: goodsItemAttrStrSlice[0],
				List: strings.Split(goodsItemAttrStrSlice[1], ","),
			}
			goodsItemAttrList = append(goodsItemAttrList, newGoodsItemAttr)
		}
	}

	return &models.GoodsInforData{
		ShopBaseData:      baseData,
		GoodsInfo:         goodsInfo,
		RelationGoodsList: relationGoodsList,
		GoodsGiftList:     goodsGiftList,
		GoodsColorList:    goodsColorList,
		GoodsFittingList:  goodsFittingList,
		GoodsImageList:    goodsImageList,
		GoodsAttrList:     goodsAttrList,
		GoodsItemAttrList: goodsItemAttrList,
	}, nil
}

func (l ProductLogic) GetImgList(goodsId, colorId int) (*models.GetImgListData, error) {
	imgList, err1 := mysql.GetImgListByGoodsIdAndColorId(goodsId, colorId)
	if err1 != nil {
		return nil, err1
	}

	// 如果imgList没有数据，需要返回所有图片
	if len(imgList) < 1 {
		imgList, err1 = mysql.GetGoodsImageListById(goodsId)
		if err1 != nil {
			return nil, err1
		}
	}

	return &models.GetImgListData{
		ImgList: imgList,
	}, nil
}
