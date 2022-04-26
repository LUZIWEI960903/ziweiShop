package models

type Nav struct {
	Id        int    // 导航栏id
	Position  int    // 导航栏位置
	IsOpennew int    // 是否打开新窗口
	Sort      int    // 导航栏排序
	Status    int    // 导航栏状态
	AddTime   int    // 导航栏创建时间
	IsDeleted int    // 是否逻辑删除导航栏
	Title     string // 导航栏名
	Link      string // 导航栏跳转链接
	Relation  string // 相关联的商品id列表
}

func (Nav) TableName() string {
	return "nav"
}
