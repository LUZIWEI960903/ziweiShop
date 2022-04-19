package admin

import (
	"strconv"
	"ziweiShop/logic"
	"ziweiShop/pkg/tools"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type GoodsController struct {
	BaseController
}

// Index 商品页面的接口
func (con GoodsController) Index(c *gin.Context) {

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
