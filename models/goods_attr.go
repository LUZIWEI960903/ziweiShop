package models

type GoodsAttr struct {
	Id              int    // 商品属性id
	GoodsId         int    // 商品id
	AttributeCateId int    // 商品属性对应的cateId
	AttributeId     int    // 商品属性的Id
	AttributeType   int    // 商品属性的类型 1. 单行文本框 2. 多行文本框 3. 下拉框
	AddTime         int    // 商品属性的创建时间
	Status          int    // 商品属性的状态
	IsDeleted       int    // 是否逻辑删除商品属性
	AttributeTitle  string // 商品属性的名
	AttributeValue  string // 商品属性的填写内容
	Sort            string // 商品属性排序
}

func (GoodsAttr) TableName() string {
	return "goods_attr"
}
