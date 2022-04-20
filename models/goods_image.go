package models

type GoodsImage struct {
	Id        int    // 商品图片id
	GoodsId   int    // 商品id
	ColorId   int    // 颜色id
	Sort      int    // 商品图片排序
	AddTime   int    // 商品图片创建时间
	Status    int    // 商品图片状态
	IsDeleted int    // 是否逻辑删除商品图片
	ImgUrl    string // 商品图片路径
}

func (GoodsImage) TableName() string {
	return "goods_image"
}
