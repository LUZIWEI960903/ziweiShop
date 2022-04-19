package admin

import (
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
