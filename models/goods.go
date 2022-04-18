package models

type Goods struct {
	Id            int     // 商品id
	CateId        int     // 商品的分类id
	ClickCount    int     // 商品点击数
	GoodsNumber   int     // 商品库存
	IsDeleted     int     // 是否逻辑删除商品
	IsHot         int     // 是否热门商品
	IsBest        int     // 是否推荐商品
	IsNew         int     // 是否新的商品
	GoodsTypeId   int     // 商品关联类型
	Sort          int     // 商品排序
	Status        int     // 商品状态
	AddTime       int     // 商品增加时间
	Price         float64 // 商品价格
	MarketPrice   float64 // 商品市场价格
	Title         string  // 商品名
	SubTitle      string  // 商品的二级标题
	GoodsSn       string  // 商品的sn号
	RelationGoods string  // 关联商品
	GoodsAttr     string  // 商品属性
	GoodsVersion  string  // 商品版本
	GoodsImg      string  // 商品图片
	GoodsGift     string  // 商品赠品
	GoodsFitting  string  // 商品配件
	GoodsColor    string  // 商品颜色
	GoodsKeywords string  // 商品关键词
	GoodsDesc     string  // 商品描述
	GoodsContent  string  // 商品详情
}

func (Goods) TableName() string {
	return "goods"
}
