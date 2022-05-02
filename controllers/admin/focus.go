package admin

import (
	"strconv"
	"ziweiShop/logic"
	"ziweiShop/models"
	"ziweiShop/pkg/tools"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type FocusController struct {
	BaseController
}

// Index 轮播图主页面的接口
func (con FocusController) Index(c *gin.Context) {
	// 业务逻辑
	focusList, err := logic.FocusLogic{}.GetFocusList()
	if err != nil {
		zap.L().Error("[pkg: admin] [(con FocusController) Index(c *gin.Context)] [logic.FocusLogic{}.GetFocusList()] failed, err:", zap.Error(err))
		con.error(c, CodeGetFocusListErr)
		return
	}
	con.success(c, focusList)
}

// Add 增加轮播图页面的接口
func (con FocusController) Add(c *gin.Context) {
	con.success(c, true)
}

// DoAdd 增加轮播图的接口
func (con FocusController) DoAdd(c *gin.Context) {
	// 解析参数
	focusType, err := strconv.Atoi(c.PostForm("focus_type"))
	if err != nil {
		zap.L().Error("[pkg: admin] [(con FocusController) DoAdd(c *gin.Context)] [strconv.Atoi(c.PostForm(\"focus_type\"))] failed, err:", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}
	sort, err := strconv.Atoi(c.PostForm("sort"))
	if err != nil {
		zap.L().Error("[pkg: admin] [(con FocusController) DoAdd(c *gin.Context)] [strconv.Atoi(c.PostForm(\"sort\"))] failed, err:", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}
	p := &models.AddFocusParams{
		FocusType: focusType,
		Sort:      sort,
		Title:     c.PostForm("title"),
		Link:      c.PostForm("link"),
	}

	// 上传文件
	focusImgSrc, err := tools.UploadImg(c, "focus_img")
	if err != nil {
		zap.L().Error("[pkg: admin] [(con FocusController) DoAdd(c *gin.Context)] [tools.UploadImg(c, \"focus_img\")] failed, err:", zap.Error(err))
		con.error(c, CodeUploadImgErr)
		return
	}
	// 业务逻辑
	err = logic.FocusLogic{}.DoAdd(p, focusImgSrc)
	if err != nil {
		zap.L().Error("[pkg: admin] [(con FocusController) DoAdd(c *gin.Context)] [logic.FocusLogic{}.DoAdd(p, focusImgSrc)] failed, err:", zap.Error(err))
		con.error(c, CodeAddFocusErr)
		return
	}
	con.success(c, true)
}

// Edit 编辑轮播图页面的接口
func (con FocusController) Edit(c *gin.Context) {
	// 解析参数
	focusIdStr := c.Query("id")
	focusId, err := strconv.Atoi(focusIdStr)
	if err != nil {
		zap.L().Error("[pkg: admin] [(con FocusController) Edit(c *gin.Context)] [strconv.Atoi(focusIdStr)] failed, err:", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}
	// 业务逻辑
	focusInfo, err := logic.FocusLogic{}.GetFocusById(focusId)
	if err != nil {
		zap.L().Error("[pkg: admin] [(con FocusController) Edit(c *gin.Context)] [logic.FocusLogic{}.GetFocusById(focusId)] failed, err:", zap.Error(err))
		con.error(c, CodeGetFocusErr)
		return
	}
	con.success(c, focusInfo)
}

// DoEdit 编辑轮播图的接口
func (con FocusController) DoEdit(c *gin.Context) {
	// 解析参数
	focusId, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		zap.L().Error("[pkg: admin] [(con FocusController) DoEdit(c *gin.Context)] [strconv.Atoi(c.PostForm(\"id\"))] failed, err:", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}
	focusType, err := strconv.Atoi(c.PostForm("focus_type"))
	if err != nil {
		zap.L().Error("[pkg: admin] [(con FocusController) DoEdit(c *gin.Context)] [strconv.Atoi(c.PostForm(\"focus_type\"))] failed, err:", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}
	sort, err := strconv.Atoi(c.PostForm("sort"))
	if err != nil {
		zap.L().Error("[pkg: admin] [(con FocusController) DoEdit(c *gin.Context)] [strconv.Atoi(c.PostForm(\"sort\"))] failed, err:", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}
	status, err := strconv.Atoi(c.PostForm("status"))
	if err != nil {
		zap.L().Error("[pkg: admin] [(con FocusController) DoEdit(c *gin.Context)] [strconv.Atoi(c.PostForm(\"status\"))] failed, err:", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}
	p := &models.EditFocusParams{
		Id:        focusId,
		FocusType: focusType,
		Sort:      sort,
		Status:    status,
		Title:     c.PostForm("title"),
		Link:      c.PostForm("link"),
	}
	// 上传文件
	focusImgSrc, _ := tools.UploadImg(c, "focus_img")

	// 业务逻辑
	err = logic.FocusLogic{}.DoEdit(p, focusImgSrc)
	if err != nil {
		zap.L().Error("[pkg: admin] [(con FocusController) DoAdd(c *gin.Context)] [logic.FocusLogic{}.DoAdd(p, focusImgSrc)] failed, err:", zap.Error(err))
		con.error(c, CodeEditFocusErr)
		return
	}
	con.success(c, true)
}

// Delete 删除轮播图的接口
func (con FocusController) Delete(c *gin.Context) {
	// 解析参数
	focusIdStr := c.Query("id")
	focusId, err := strconv.Atoi(focusIdStr)
	if err != nil {
		zap.L().Error("[pkg: admin] [(con FocusController) Delete(c *gin.Context)] [strconv.Atoi(focusIdStr)] failed, err:", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}
	// 业务逻辑
	err = logic.FocusLogic{}.DeleteFocus(focusId)
	if err != nil {
		zap.L().Error("[pkg: admin] [(con FocusController) Delete(c *gin.Context)] [logic.FocusLogic{}.DeleteFocus(focusId)] failed, err:", zap.Error(err))
		con.error(c, CodeDeleteFocusErr)
		return
	}
	con.success(c, true)
}
