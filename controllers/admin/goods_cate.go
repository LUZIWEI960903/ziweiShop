package admin

import (
	"errors"
	"strconv"
	"ziweiShop/logic"
	"ziweiShop/models"
	"ziweiShop/pkg/tools"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type GoodsCateController struct {
	BaseController
}

// Index 商品分类页面的接口
func (con GoodsCateController) Index(c *gin.Context) {

}

// Add 增加商品分类页面的接口
func (con GoodsCateController) Add(c *gin.Context) {

}

// DoAdd 增加商品分类的接口
func (con GoodsCateController) DoAdd(c *gin.Context) {
	// 解析参数
	pid, err1 := strconv.Atoi(c.PostForm("pid"))
	status, err2 := strconv.Atoi(c.PostForm("status"))
	sort, err3 := strconv.Atoi(c.PostForm("sort"))

	if err1 != nil || err2 != nil || err3 != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsCateController) DoAdd(c *gin.Context)] [strconv.Atoi(c.PostForm())] failed, ", zap.Error(errors.New("Invalid params")))
		con.error(c, CodeInValidParams)
		return
	}
	p := &models.AddGoodsCateParams{
		Pid:         pid,
		Status:      status,
		Sort:        sort,
		Title:       c.PostForm("title"),
		Link:        c.PostForm("link"),
		Template:    c.PostForm("template"),
		SubTitle:    c.PostForm("sub_title"),
		Keywords:    c.PostForm("keywords"),
		Description: c.PostForm("description"),
	}

	cateImg, err := tools.UploadImg(c, "cate_img")
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsCateController) DoAdd(c *gin.Context)] [tools.UploadImg(c, \"cate_img\")] failed, ", zap.Error(err))
		con.error(c, CodeUploadImgErr)
		return
	}

	// 业务逻辑
	err = logic.GoodsCateLogic{}.DoAdd(p, cateImg)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsCateController) DoAdd(c *gin.Context)] [logic.GoodsCateLogic{}.DoAdd(p, cateImg)] failed, ", zap.Error(err))
		con.error(c, CodeAddGoodsCateErr)
		return
	}
	con.success(c, true)
}

// Edit 编辑商品分类页面的接口
func (con GoodsCateController) Edit(c *gin.Context) {

}

// DoEdit 编辑商品分类的接口
func (con GoodsCateController) DoEdit(c *gin.Context) {

}

// Delete 删除商品分类的接口
func (con GoodsCateController) Delete(c *gin.Context) {

}
