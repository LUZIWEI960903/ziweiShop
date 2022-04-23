package admin

import (
	"strconv"
	"strings"
	"ziweiShop/logic"
	"ziweiShop/models"
	"ziweiShop/pkg/tools"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type GoodsController struct {
	BaseController
}

// Index 商品页面的接口
func (con GoodsController) Index(c *gin.Context) {
	// 业务逻辑
	data, err := logic.GoodsLogic{}.ShowIndexPageDataLogic()
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsController) Index(c *gin.Context)] [logic.GoodsLogic{}.ShowIndexPageDataLogic()] failed, ", zap.Error(err))
		con.error(c, CodeGetIndexPageDataErr)
		return
	}

	con.success(c, data)
}

// Add 增加商品页面的接口
func (con GoodsController) Add(c *gin.Context) {
	// 业务逻辑
	data, err := logic.GoodsLogic{}.ShowAddPageLogic()
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsController) Add(c *gin.Context)] [logic.GoodsLogic{}.ShowAddPageLogic()] failed, ", zap.Error(err))
		con.error(c, CodeGetAddPageDataErr)
		return
	}

	con.success(c, data)
}

// DoAdd 增加商品的接口
func (con GoodsController) DoAdd(c *gin.Context) {
	// 解析参数
	cateId, err1 := strconv.Atoi(c.PostForm("cate_id"))
	goodsNumber, err2 := strconv.Atoi(c.PostForm("goods_number"))
	markerPrice, err3 := strconv.ParseFloat(c.PostForm("market_price"), 64)
	price, err4 := strconv.ParseFloat(c.PostForm("price"), 64)
	isHot, err5 := strconv.Atoi(c.PostForm("is_hot"))
	isBest, err6 := strconv.Atoi(c.PostForm("is_best"))
	isNew, err7 := strconv.Atoi(c.PostForm("is_new"))
	goodsTypeId, err8 := strconv.Atoi(c.PostForm("goods_type_id"))
	sort, err9 := strconv.Atoi(c.PostForm("sort"))
	status, err10 := strconv.Atoi(c.PostForm("status"))

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil || err7 != nil || err8 != nil || err9 != nil || err10 != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsController) DoAdd(c *gin.Context)]")
		con.error(c, CodeInValidParams)
		return
	}

	goodsColorArrStr := strings.Join(c.PostFormArray("goods_color"), ",")

	goodsImg, _ := tools.UploadImg(c, "goods_img")

	// 获取图库信息
	goodsImageList := c.PostFormArray("goods_image_list")

	// 获取规格包装信息
	attrIdList := c.PostFormArray("attr_id_list")
	attrValueList := c.PostFormArray("attr_value_list")

	p := &models.AddGoodsParams{
		Title:    c.PostForm("title"),
		SubTitle: c.PostForm("sub_title"),
		GoodsSn:  c.PostForm("goods_sn"),
		CateId:   cateId,

		GoodsNumber: goodsNumber,
		MarketPrice: markerPrice,
		Price:       price,

		RelationGoods: c.PostForm("relation_goods"),
		GoodsAttr:     c.PostForm("goods_attr"),
		GoodsVersion:  c.PostForm("goods_version"),
		GoodsGift:     c.PostForm("goods_gift"),
		GoodsFitting:  c.PostForm("goods_fitting"),
		GoodsColor:    goodsColorArrStr,
		GoodsKeywords: c.PostForm("goods_keywords"),
		GoodsDesc:     c.PostForm("goods_desc"),
		GoodsContent:  c.PostForm("goods_content"),
		GoodsImg:      goodsImg,
		GoodsTypeId:   goodsTypeId,

		IsHot:   isHot,
		IsBest:  isBest,
		IsNew:   isNew,
		Sort:    sort,
		Status:  status,
		AddTime: int(tools.GetUnix()),
	}

	// 业务逻辑
	err := logic.GoodsLogic{}.AddGoodsLogic(p, goodsImageList, attrIdList, attrValueList)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsController) DoAdd(c *gin.Context)] [logic.GoodsLogic{}.AddGoodsLogic(p, goodsImageList, attrIdList, attrValueList)] failed, ", zap.Error(err))
		con.error(c, CodeDoAddLogicErr)
		return
	}

	con.success(c, true)
}

// Edit 编辑商品页面的接口
func (con GoodsController) Edit(c *gin.Context) {
	// 解析参数
	goodsId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsController) Edit(c *gin.Context)] [strconv.Atoi(c.Query(\"id\"))] failed, ", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}
	// 业务逻辑
	data, err1 := logic.GoodsLogic{}.ShowEditPageLogic(goodsId)
	if err1 != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsController) Edit(c *gin.Context)] [logic.GoodsLogic{}.ShowEditPageLogic(goodsId)] failed, ", zap.Error(err1))
		con.error(c, CodeGetEditPageDataErr)
		return
	}

	con.success(c, data)
}

// DoEdit 编辑商品的接口
func (con GoodsController) DoEdit(c *gin.Context) {
	// 解析参数
	goodsId, err11 := strconv.Atoi(c.PostForm("id"))
	cateId, err1 := strconv.Atoi(c.PostForm("cate_id"))
	goodsNumber, err2 := strconv.Atoi(c.PostForm("goods_number"))
	markerPrice, err3 := strconv.ParseFloat(c.PostForm("market_price"), 64)
	price, err4 := strconv.ParseFloat(c.PostForm("price"), 64)
	isHot, err5 := strconv.Atoi(c.PostForm("is_hot"))
	isBest, err6 := strconv.Atoi(c.PostForm("is_best"))
	isNew, err7 := strconv.Atoi(c.PostForm("is_new"))
	goodsTypeId, err8 := strconv.Atoi(c.PostForm("goods_type_id"))
	sort, err9 := strconv.Atoi(c.PostForm("sort"))
	status, err10 := strconv.Atoi(c.PostForm("status"))

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil || err7 != nil || err8 != nil || err9 != nil || err10 != nil || err11 != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsController) DoEdit(c *gin.Context)]")
		con.error(c, CodeInValidParams)
		return
	}

	goodsColorArrStr := strings.Join(c.PostFormArray("goods_color"), ",")

	goodsImg, _ := tools.UploadImg(c, "goods_img")

	// 获取图库信息
	goodsImageList := c.PostFormArray("goods_image_list")

	// 获取规格包装信息
	attrIdList := c.PostFormArray("attr_id_list")
	attrValueList := c.PostFormArray("attr_value_list")

	p := &models.EditGoodsParams{
		Id:       goodsId,
		Title:    c.PostForm("title"),
		SubTitle: c.PostForm("sub_title"),
		GoodsSn:  c.PostForm("goods_sn"),
		CateId:   cateId,

		GoodsNumber: goodsNumber,
		MarketPrice: markerPrice,
		Price:       price,

		RelationGoods: c.PostForm("relation_goods"),
		GoodsAttr:     c.PostForm("goods_attr"),
		GoodsVersion:  c.PostForm("goods_version"),
		GoodsGift:     c.PostForm("goods_gift"),
		GoodsFitting:  c.PostForm("goods_fitting"),
		GoodsColor:    goodsColorArrStr,
		GoodsKeywords: c.PostForm("goods_keywords"),
		GoodsDesc:     c.PostForm("goods_desc"),
		GoodsContent:  c.PostForm("goods_content"),
		GoodsImg:      goodsImg,
		GoodsTypeId:   goodsTypeId,

		IsHot:  isHot,
		IsBest: isBest,
		IsNew:  isNew,
		Sort:   sort,
		Status: status,
	}

	// 业务逻辑
	err := logic.GoodsLogic{}.EditGoodsLogic(p, goodsImageList, attrIdList, attrValueList)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsController) DoEdit(c *gin.Context)] [logic.GoodsLogic{}.EditGoodsLogic(p, goodsImageList, attrIdList, attrValueList)] failed, ", zap.Error(err))
		con.error(c, CodeDoEditLogicErr)
		return
	}

	con.success(c, true)
}

// ImageUpload wysiwyg-editor上传图片的接口
func (con GoodsController) ImageUpload(c *gin.Context) {
	ImgSrc, err := tools.UploadImg(c, "file")
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsController) ImageUpload(c *gin.Context)] [tools.UploadImg(c, \"file\")] failed, ", zap.Error(err))
		con.error(c, CodeUploadImgErr)
		return
	}

	con.success(c, gin.H{"link": "/" + ImgSrc})
}

// GoodsTypeAttribute 使用Ajax 根据cateId 动态获取对应的商品类型属性的接口
func (con GoodsController) GoodsTypeAttribute(c *gin.Context) {
	// 解析参数
	cateId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsController) GoodsTypeAttribute(c *gin.Context)] [strconv.Atoi(c.Query(\"id\"))] failed, ", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}

	// 业务逻辑
	data, err := logic.GoodsLogic{}.AjaxGetGoodsTypeAttributeLogic(cateId)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsController) GoodsTypeAttribute(c *gin.Context)] [logic.GoodsLogic{}.AjaxGetGoodsTypeAttributeLogic(cateId)] failed, ", zap.Error(err))
		con.error(c, CodeAjaxErr)
		return
	}

	con.success(c, data)
}

// ChangeGoodsImageColor 使用Ajax 根据 imageId 改变其 colorId的接口
func (con GoodsController) ChangeGoodsImageColor(c *gin.Context) {
	// 解析参数
	goodsImageId, err1 := strconv.Atoi(c.Query("goods_image_id"))
	colorId, err2 := strconv.Atoi(c.Query("color_id"))
	if err1 != nil || err2 != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsController) ChangeGoodsImageColor(c *gin.Context)]")
		con.error(c, CodeInValidParams)
		return
	}

	// 业务逻辑
	err3 := logic.GoodsLogic{}.AjaxChangeGoodsImageLogic(goodsImageId, colorId)
	if err3 != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsController) ChangeGoodsImageColor(c *gin.Context)] [logic.GoodsLogic{}.AjaxChangeGoodsImageLogic(goodsImageId, colorId)] failed, ", zap.Error(err3))
		con.error(c, CodeAjaxErr)
		return
	}

	con.success(c, true)
}
