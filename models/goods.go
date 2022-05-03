package models

type Goods struct {
	Id            int     `json:"id,omitempty"`             // 商品id
	CateId        int     `json:"cate_id,omitempty"`        // 商品的分类id
	ClickCount    int     `json:"click_count,omitempty"`    // 商品点击数
	GoodsNumber   int     `json:"goods_number,omitempty"`   // 商品库存
	IsDeleted     int     `json:"is_deleted,omitempty"`     // 是否逻辑删除商品
	IsHot         int     `json:"is_hot,omitempty"`         // 是否热门商品
	IsBest        int     `json:"is_best,omitempty"`        // 是否推荐商品
	IsNew         int     `json:"is_new,omitempty"`         // 是否新的商品
	GoodsTypeId   int     `json:"goods_type_id,omitempty"`  // 商品关联类型
	Sort          int     `json:"sort,omitempty"`           // 商品排序
	Status        int     `json:"status,omitempty"`         // 商品状态
	AddTime       int     `json:"add_time,omitempty"`       // 商品增加时间
	Price         float64 `json:"price,omitempty"`          // 商品价格
	MarketPrice   float64 `json:"market_price,omitempty"`   // 商品市场价格
	Title         string  `json:"title,omitempty"`          // 商品名
	SubTitle      string  `json:"sub_title,omitempty"`      // 商品的二级标题
	GoodsSn       string  `json:"goods_sn,omitempty"`       // 商品的sn号
	RelationGoods string  `json:"relation_goods,omitempty"` // 关联商品
	GoodsAttr     string  `json:"goods_attr,omitempty"`     // 商品属性
	GoodsVersion  string  `json:"goods_version,omitempty"`  // 商品版本
	GoodsImg      string  `json:"goods_img,omitempty"`      // 商品图片
	GoodsGift     string  `json:"goods_gift,omitempty"`     // 商品赠品
	GoodsFitting  string  `json:"goods_fitting,omitempty"`  // 商品配件
	GoodsColor    string  `json:"goods_color,omitempty"`    // 商品颜色
	GoodsKeywords string  `json:"goods_keywords,omitempty"` // 商品关键词
	GoodsDesc     string  `json:"goods_desc,omitempty"`     // 商品描述
	GoodsContent  string  `json:"goods_content,omitempty"`  // 商品详情
}

func (Goods) TableName() string {
	return "goods"
}
