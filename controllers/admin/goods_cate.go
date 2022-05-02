package admin

import (
	"errors"
	"strconv"
	logic "ziweiShop/logic/admin"
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
	// 业务逻辑
	TopGoodsCateWithGoodsCateList, err := logic.GoodsCateLogic{}.GetTopGoodsCateWithGoodsCateList()
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsCateController) Index(c *gin.Context)] [logic.GoodsCateLogic{}.GetTopGoodsCateWithGoodsCateList()] failed, ", zap.Error(err))
		con.error(c, CodeGetTopGoodsCateWithGoodsCateListErr)
		return
	}
	con.success(c, TopGoodsCateWithGoodsCateList)
}

// Add 增加商品分类页面的接口
func (con GoodsCateController) Add(c *gin.Context) {
	// 业务逻辑
	TopGoodsCateList, err := logic.GoodsCateLogic{}.GetTopGoodsCateList()
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsCateController) Add(c *gin.Context)] [logic.GoodsCateLogic{}.GetTopGoodsCateList()] failed, ", zap.Error(err))
		con.error(c, CodeGetTopGoodsCateListErr)
		return
	}
	con.success(c, TopGoodsCateList)
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

	cateImg, _ := tools.UploadImg(c, "cate_img")

	// 业务逻辑
	err := logic.GoodsCateLogic{}.DoAdd(p, cateImg)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsCateController) DoAdd(c *gin.Context)] [logic.GoodsCateLogic{}.DoAdd(p, cateImg)] failed, ", zap.Error(err))
		con.error(c, CodeAddGoodsCateErr)
		return
	}
	con.success(c, true)
}

// Edit 编辑商品分类页面的接口
func (con GoodsCateController) Edit(c *gin.Context) {
	// 解析参数
	goodsCateId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsCateController) Edit(c *gin.Context)] [strconv.Atoi(c.Query(\"id\"))] failed, ", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}
	// 业务逻辑
	goodsCateInfo, err := logic.GoodsCateLogic{}.GetGoodsCateInfo(goodsCateId)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsCateController) Edit(c *gin.Context)] [logic.GoodsCateLogic{}.GetGoodsCateInfo(goodsCateId)] failed, ", zap.Error(err))
		con.error(c, CodeGetGoodsCateInfoErr)
		return
	}
	con.success(c, goodsCateInfo)
}

// DoEdit 编辑商品分类的接口
func (con GoodsCateController) DoEdit(c *gin.Context) {
	// 解析参数
	id, err1 := strconv.Atoi(c.PostForm("id"))
	pid, err2 := strconv.Atoi(c.PostForm("pid"))
	status, err3 := strconv.Atoi(c.PostForm("status"))
	sort, err4 := strconv.Atoi(c.PostForm("sort"))

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsCateController) DoAdd(c *gin.Context)] [strconv.Atoi(c.PostForm())] failed, ", zap.Error(errors.New("Invalid params")))
		con.error(c, CodeInValidParams)
		return
	}
	p := &models.EditGoodsCateParams{
		Id:          id,
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

	cateImg, _ := tools.UploadImg(c, "cate_img")

	// 业务逻辑
	err := logic.GoodsCateLogic{}.DoEdit(p, cateImg)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsCateController) DoEdit(c *gin.Context)] [logic.GoodsCateLogic{}.DoEdit(p, cateImg)] failed, ", zap.Error(err))
		con.error(c, CodeEditGoodsCateErr)
		return
	}
	con.success(c, true)
}

// Delete 删除商品分类的接口
func (con GoodsCateController) Delete(c *gin.Context) {
	// 解析参数
	goodsCateId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsCateController) Delete(c *gin.Context)] [strconv.Atoi(c.Query(\"id\"))] failed, ", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}
	// 业务逻辑
	err = logic.GoodsCateLogic{}.Delete(goodsCateId)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsCateController) Delete(c *gin.Context)] [logic.GoodsCateLogic{}.Delete(goodsCateId)] failed, ", zap.Error(err))
		con.error(c, CodeDeleteGoodsCateErr)
		return
	}
	con.success(c, true)
}
