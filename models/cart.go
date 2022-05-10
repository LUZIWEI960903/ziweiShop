package models

type Cart struct {
	Id           int     `json:"id,omitempty"`            // 商品的id
	Title        string  `json:"title,omitempty"`         // 商品的名字
	Price        float64 `json:"price,omitempty"`         // 商品的价格
	GoodsVersion string  `json:"goods_version,omitempty"` // 商品的版本
	Uid          int     `json:"uid,omitempty"`           // 用户id
	Num          int     `json:"num,omitempty"`           // 商品的数量
	GoodsGift    string  `json:"goods_gift,omitempty"`    // 商品的赠品
	GoodsFitting string  `json:"goods_fitting,omitempty"` // 商品的配件
	GoodsColor   string  `json:"goods_color,omitempty"`   // 商品颜色
	GoodsImg     string  `json:"goods_img,omitempty"`     // 商品图片
	GoodsAttr    string  `json:"goods_attr,omitempty"`    // 商品的更多属性
	Checked      bool    `json:"checked,omitempty"`       // 是否选中
}
