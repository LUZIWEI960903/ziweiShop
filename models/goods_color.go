package models

type GoodsColor struct {
	Id         int    // 商品颜色id
	Status     int    // 商品颜色的状态
	IsDeleted  int    // 是否逻辑删除商品颜色
	ColorName  string // 商品颜色名
	ColorValue string // 商品颜色值
}

func (GoodsColor) TableName() string {
	return "goods_color"
}
