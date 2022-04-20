package models

type GoodsImage struct {
	Id        int    // 商品图片id
	GoodsId   int    // 商品id
	ImgUrl    string // 商品图片路径
	ColorId   int    // 颜色id
	Sort      int    // 商品图片排序
	AddTime   int    // 商品图片创建时间
	Status    int    // 商品图片状态
	IsDeleted int    // 是否逻辑删除商品图片
}

func (GoodsImage) TableName() string {
	return "goods_image"
}
