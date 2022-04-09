package models

type Focus struct {
	Id        int    // 轮播图id
	FocusType int    // 轮播图类型 1.web 2.app 3.小程序
	Sort      int    // 排序方式
	Status    int    // 状态
	AddTime   int    // 创建时间
	IsDeleted int    // 是否逻辑删除
	Title     string // 轮播图名称
	FocusImg  string // 轮播图图片
	Link      string // 跳转地址
}

func (Focus) TableName() string {
	return "focus"
}
