package models

type GoodsType struct {
	Id          int    // 商品类型id
	Title       string // 商品类型名
	Description string // 商品类型描述
	Status      int    // 商品类型状态
	AddTime     int    // 商品类型创建时间
	IsDeleted   int    // 商品类型是否逻辑删除
}

func (GoodsType) TableName() string {
	return "goods_type"
}
