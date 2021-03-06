package models

type GoodsTypeAttribute struct {
	Id        int    // 商品类型属性的id
	CateId    int    // 对应的商品类型id
	Status    int    // 商品类型的状态
	Sort      int    // 商品类型的排序
	AddTime   int    // 商品类型的创建时间
	IsDeleted int    // 是否逻辑删除商品类型
	AttrType  int    // 商品类型属性的录入方式 1. 单行文本框 2. 多行文本框 3. 从下面的列表中选择（一行代表一个可选值 ）
	Title     string // 商品类型属性名
	AttrValue string // 商品类型属性的录入值
}

func (GoodsTypeAttribute) TableName() string {
	return "goods_type_attribute"
}
