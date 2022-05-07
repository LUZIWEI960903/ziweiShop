package models

type GoodsAttr struct {
	Id              int    `json:"id,omitempty"`                // 商品属性id
	GoodsId         int    `json:"goods_id,omitempty"`          // 商品id
	AttributeCateId int    `json:"attribute_cate_id,omitempty"` // 商品属性对应的cateId
	AttributeId     int    `json:"attribute_id,omitempty"`      // 商品属性的Id
	AttributeType   int    `json:"attribute_type,omitempty"`    // 商品属性的类型 1. 单行文本框 2. 多行文本框 3. 下拉框
	AddTime         int    `json:"add_time,omitempty"`          // 商品属性的创建时间
	Status          int    `json:"status,omitempty"`            // 商品属性的状态
	Sort            int    `json:"sort,omitempty"`              // 商品属性排序
	IsDeleted       int    `json:"is_deleted,omitempty"`        // 是否逻辑删除商品属性
	AttributeTitle  string `json:"attribute_title,omitempty"`   // 商品属性的名
	AttributeValue  string `json:"attribute_value,omitempty"`   // 商品属性的填写内容
}

func (GoodsAttr) TableName() string {
	return "goods_attr"
}
