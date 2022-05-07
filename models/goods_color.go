package models

type GoodsColor struct {
	Id         int    `json:"id,omitempty"`          // 商品颜色id
	Status     int    `json:"status,omitempty"`      // 商品颜色的状态
	IsDeleted  int    `json:"is_deleted,omitempty"`  // 是否逻辑删除商品颜色
	ColorName  string `json:"color_name,omitempty"`  // 商品颜色名
	ColorValue string `json:"color_value,omitempty"` // 商品颜色值
}

func (GoodsColor) TableName() string {
	return "goods_color"
}
