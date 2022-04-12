package models

type GoodsCate struct {
	Id          int    // 商品分类id
	Pid         int    // 二级商品分类id与Id自关联
	Status      int    // 商品分类状态
	Sort        int    // 商品分类排序
	AddTime     int    // 创建商品分类的时间戳
	IsDeleted   int    // 商品分类逻辑删除
	Title       string // 商品分类名
	CateImg     string // 商品分类图片地址
	Link        string // 商品分类跳转地址
	Template    string // 商品分类加载的模板
	SubTitle    string // 商品分类Seo标题
	Keywords    string // 商品分类Seo关键词
	Description string // 商品分类描述
}

func (GoodsCate) TableName() string {
	return "goods_cate"
}
