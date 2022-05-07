package models

type GoodsImage struct {
	Id        int    `json:"id,omitempty"`         // 商品图片id
	GoodsId   int    `json:"goods_id,omitempty"`   // 商品id
	ColorId   int    `json:"color_id,omitempty"`   // 颜色id
	Sort      int    `json:"sort,omitempty"`       // 商品图片排序
	AddTime   int    `json:"add_time,omitempty"`   // 商品图片创建时间
	Status    int    `json:"status,omitempty"`     // 商品图片状态
	IsDeleted int    `json:"is_deleted,omitempty"` // 是否逻辑删除商品图片
	ImgUrl    string `json:"img_url,omitempty"`    // 商品图片路径
}

func (GoodsImage) TableName() string {
	return "goods_image"
}
