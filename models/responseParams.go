package models

type NewRoleList struct {
	Id    int    `json:"id"`    // role的id
	Title string `json:"title"` // 角色名
}

type IndexManagerListRole struct {
	Id    int    `json:"id"`    // role的id
	Title string `json:"title"` // 角色名
}

type IndexManagerList struct {
	Id       int                  `json:"id"`       // manager的id
	Username string               `json:"username"` // manager的名字
	Mobile   string               `json:"mobile"`   // manager的电话
	Email    string               `json:"email"`    // manager的邮箱
	AddTime  string               `json:"add_time"` // manager的创建时间
	Role     IndexManagerListRole `json:"role"`     // manager对应role的信息
}

type EditManagerList struct {
	Id       int           `json:"id"`        // manager的id
	Username string        `json:"username"`  // manager的名字
	Password string        `json:"password"`  // manager的密码
	Mobile   string        `json:"mobile"`    // manager的电话
	Email    string        `json:"email"`     // manager的邮箱
	RoleList []NewRoleList `json:"role_list"` // 所有role的信息
}

type NewTopAccess struct {
	Id         int    `json:"id"`          // 权限id
	ModuleName string `json:"module_name"` // 模块名称
}

type NewAccessList struct {
	Id          int    `json:"id"`          // 权限id
	Type        int    `json:"type"`        // 模块类型: 1.表示模块 2.表示菜单 3.表示操作
	ModuleId    int    `json:"module_id"`   // 与当前Id自关联
	Sort        int    `json:"sort"`        // 排序
	ActionName  string `json:"action_name"` // 操作名称
	ModuleName  string `json:"module_name"` // 模块名称
	Url         string `json:"url"`         // 路由跳转地址
	Description string `json:"description"` // 描述
}

type NewTopAccessListWithAccessList struct {
	Id          int    `json:"id"`          // 权限id
	Type        int    `json:"type"`        // 模块类型: 1.表示模块 2.表示菜单 3.表示操作
	ModuleId    int    `json:"module_id"`   // 与当前Id自关联
	Sort        int    `json:"sort"`        // 排序
	ActionName  string `json:"action_name"` // 操作名称
	ModuleName  string `json:"module_name"` // 模块名称
	Url         string `json:"url"`         // 路由跳转地址
	Description string `json:"description"` // 描述

	AccessList []NewAccessList `json:"access_list"` // 该权限下的模块
}

type AccessType struct {
	Id       int    `json:"id"`        // access type id
	TypeName string `json:"type_name"` // 模块类型名字: 1.表示模块 2.表示菜单 3.表示操作
}

type EditAccessInfo struct {
	Id          int          `json:"id"`          // 权限id
	TypeList    []AccessType `json:"type_list"`   // 模块类型: 1.表示模块 2.表示菜单 3.表示操作
	ModuleId    int          `json:"module_id"`   // 与当前Id自关联
	Sort        int          `json:"sort"`        // 排序
	ActionName  string       `json:"action_name"` // 操作名称
	ModuleName  string       `json:"module_name"` // 模块名称
	Url         string       `json:"url"`         // 路由跳转地址
	Description string       `json:"description"` // 描述

	NewTopAccessList []NewTopAccess `json:"access_list"` // 所有的顶级模块
}

type ResponseAccessItemAuth struct {
	AccessId   int    `json:"access_id"`   // 权限id
	ActionName string `json:"action_name"` // 操作名称
	IsCheck    bool   `json:"is_check"`    // 是否已授权
}

type ResponseTopAccessItemAuth struct {
	AccessId   int                      `json:"access_id"`       // 权限id
	ModuleName string                   `json:"module_name"`     // 模块名字
	IsCheck    bool                     `json:"is_check"`        // 是否已授权
	AccessItem []ResponseAccessItemAuth `json:"son_access_list"` // 子模块列表
}

type ResponseAccessUrl struct {
	Url string // 路由跳转地址
}

type ResponseFocus struct {
	Id        int    `json:"id"`         // 轮播图id
	FocusType int    `json:"focus_type"` // 轮播图类型 1.web 2.app 3.小程序
	Sort      int    `json:"sort"`       // 排序方式
	Status    int    `json:"status"`     // 状态
	Title     string `json:"title"`      // 轮播图名称
	FocusImg  string `json:"focus_img"`  // 轮播图图片
	Link      string `json:"link"`       // 跳转地址
}

type FocusType struct {
	TypeId int    `json:"type_id"`
	Type   string `json:"type"`
}

type ResponseEditFocus struct {
	Id        int         `json:"id"`         // 轮播图id
	FocusType []FocusType `json:"focus_type"` // 轮播图类型 1.web 2.app 3.小程序
	Sort      int         `json:"sort"`       // 排序方式
	Status    int         `json:"status"`     // 状态
	Title     string      `json:"title"`      // 轮播图名称
	FocusImg  string      `json:"focus_img"`  // 轮播图图片
	Link      string      `json:"link"`       // 跳转地址
}

type TopGoodsCate struct {
	Id          int    `json:"id"`          // 商品分类id
	Pid         int    `json:"pid"`         // 二级商品分类id与Id自关联
	Status      int    `json:"status"`      // 商品分类状态
	Sort        int    `json:"sort"`        // 商品分类排序
	Title       string `json:"title"`       // 商品分类名
	CateImg     string `json:"cate_img"`    // 商品分类图片地址
	Link        string `json:"link"`        // 商品分类跳转地址
	Template    string `json:"template"`    // 商品分类加载的模板
	SubTitle    string `json:"sub_title"`   // 商品分类Seo标题
	Keywords    string `json:"keywords"`    // 商品分类Seo关键词
	Description string `json:"description"` // 商品分类描述
}

type GoodsCateWithGoodsCate struct {
	Id             int            `json:"id"`               // 商品分类id
	Pid            int            `json:"pid"`              // 二级商品分类id与Id自关联
	Status         int            `json:"status"`           // 商品分类状态
	Sort           int            `json:"sort"`             // 商品分类排序
	Title          string         `json:"title"`            // 商品分类名
	CateImg        string         `json:"cate_img"`         // 商品分类图片地址
	Link           string         `json:"link"`             // 商品分类跳转地址
	Template       string         `json:"template"`         // 商品分类加载的模板
	SubTitle       string         `json:"sub_title"`        // 商品分类Seo标题
	Keywords       string         `json:"keywords"`         // 商品分类Seo关键词
	Description    string         `json:"description"`      // 商品分类描述
	GoodsCateItems []TopGoodsCate `json:"goods_cate_items"` // 该商品分类下的子分类
}

type TopGoodsCate1 struct {
	Id    int    `json:"id"`    // 商品分类id
	Pid   int    `json:"pid"`   // 二级商品分类id与Id自关联
	Title string `json:"title"` // 商品分类名
}

type GoodsCateInfo struct {
	Id            int             `json:"id"`             // 商品分类id
	Status        int             `json:"status"`         // 商品分类状态
	Sort          int             `json:"sort"`           // 商品分类排序
	Title         string          `json:"title"`          // 商品分类名
	CateImg       string          `json:"cate_img"`       // 商品分类图片地址
	Link          string          `json:"link"`           // 商品分类跳转地址
	Template      string          `json:"template"`       // 商品分类加载的模板
	SubTitle      string          `json:"sub_title"`      // 商品分类Seo标题
	Keywords      string          `json:"keywords"`       // 商品分类Seo关键词
	Description   string          `json:"description"`    // 商品分类描述
	TopGoodsCate1 []TopGoodsCate1 `json:"top_goods_cate"` // 所有的顶级分类
}

type GetGoodsType struct {
	Id          int    `json:"id"`           // 商品类型id
	Status      int    `json:"status"`       // 商品类型状态
	Title       string `json:"title"`        // 商品类型名
	Description string `json:"descriptions"` // 商品类型描述
}

type GoodsTypeList1 struct {
	Id    int    `json:"id"`    // 商品类型id
	Title string `json:"title"` // 商品类型名
}

type GoodsTypeInfo struct {
	Id          int    `json:"id"`           // 商品类型id
	Status      int    `json:"status"`       // 商品类型状态
	Title       string `json:"title"`        // 商品类型名
	Description string `json:"descriptions"` // 商品类型描述
}

type GoodsTypeAttributeAddPageData struct {
	CateId         int             `json:"cate_id"`          // 对应的商品类型id
	GoodsTypeItems []GoodsTypeInfo `json:"goods_type_items"` // 商品类型列表
}

type GoodsTypeAttributeInfo struct {
	Id        int    `json:"id"`         // 商品类型属性的id
	Status    int    `json:"status"`     // 商品类型属性的状态
	Sort      int    `json:"sort"`       // 商品类型属性的排序
	AttrType  int    `json:"attr_type"`  // 商品类型属性的录入方式 1. 单行文本框 2. 多行文本框 3. 从下面的列表中选择（一行代表一个可选值 ）
	Title     string `json:"title"`      // 商品类型属性名
	AttrValue string `json:"attr_value"` // 商品类型属性的录入值
	AddTime   string `json:"add_time"`   // 创建商品类型属性的时间
}

type GoodsTypeAttributeIndexPageData struct {
	CateId                      int                      `json:"cate_id"`                    // 对应的商品类型id
	Title                       string                   `json:"goods_type"`                 // 对应的商品类型
	GoodsTypeAttributeInfoItems []GoodsTypeAttributeInfo `json:"goods_type_attribute_items"` // 商品类型属性列表
}

type GoodsTypeItems struct {
	Id    int    `json:"id"`    // 商品类型的id
	Title string `json:"title"` // 商品类型名
}

type GoodsTypeAttributeEditPageData struct {
	Id        int    `json:"id"`         // 商品类型属性的id
	Status    int    `json:"status"`     // 商品类型属性的状态
	Sort      int    `json:"sort"`       // 商品类型属性的排序
	AttrType  int    `json:"attr_type"`  // 商品类型属性的录入方式 1. 单行文本框 2. 多行文本框 3. 从下面的列表中选择（一行代表一个可选值 ）
	Title     string `json:"title"`      // 商品类型属性名
	AttrValue string `json:"attr_value"` // 商品类型属性的录入值

	GoodsTypeItems []GoodsTypeItems `json:"goods_type_items"` // 商品类型列表
}

type TopGoodsCate2 struct {
	Id    int    `json:"id"`    // 商品分类id
	Pid   int    `json:"pid"`   // 二级商品分类id与Id自关联
	Title string `json:"title"` // 商品分类名
}

type GoodsCateWithGoodsCate1 struct {
	Id                  int             `json:"id"`                     // 商品分类id
	Pid                 int             `json:"pid"`                    // 二级商品分类id与Id自关联
	Title               string          `json:"title"`                  // 商品分类名
	ChildGoodsCateItems []TopGoodsCate2 `json:"child_goods_cate_items"` // 该商品分类下的子分类
}

type GoodsAddPageData struct {
	GoodsCateItems  []GoodsCateWithGoodsCate1 `json:"goods_cate_items"`
	GoodsColorItems []GoodsColorList1         `json:"goods_color_items"`
	GoodsTypeItems  []GoodsTypeList1          `json:"goods_type_items"`
}

type GoodsColorList struct {
	Id         int    `json:"id"`          // 商品颜色id
	Status     int    `json:"status"`      // 商品颜色状态
	ColorName  string `json:"color_name"`  // 商品颜色名
	ColorValue string `json:"color_value"` // 商品颜色值
}

type GoodsColorList1 struct {
	Id        int    `json:"id"`         // 商品颜色id
	ColorName string `json:"color_name"` // 商品颜色名
}

type GoodsColorList2 struct {
	Id        int    `json:"id"`         // 商品颜色id
	ColorName string `json:"color_name"` // 商品颜色名
	IsCheck   bool   `json:"is_check"`   // 是否有选中
}

type AjaxGoodsTypeAttribute struct {
	Id        int    `json:"id"`         // 商品类型属性的id
	CateId    int    `json:"cate_id"`    // 对应的商品类型id
	Title     string `json:"title"`      // 商品类型属性名
	AttrType  int    `json:"attr_type"`  // 商品类型属性的录入方式 1. 单行文本框 2. 多行文本框 3. 从下面的列表中选择（一行代表一个可选值 ）
	AttrValue string `json:"attr_value"` // 商品类型属性的录入值
}

type GoodsIndexPage struct {
	Id          int     `json:"id"`           // 商品id
	Title       string  `json:"title"`        // 商品名
	Price       float64 `json:"price"`        // 商品价格
	MarketPrice float64 `json:"market_price"` // 商品市场价格
	IsHot       int     `json:"is_hot"`       // 是否热门商品
	IsBest      int     `json:"is_best"`      // 是否推荐商品
	IsNew       int     `json:"is_new"`       // 是否新的商品
	ClickCount  int     `json:"click_count"`  // 商品点击数
	GoodsNumber int     `json:"goods_number"` // 商品库存
	Sort        int     `json:"sort"`         // 商品排序
	Status      int     `json:"status"`       // 商品状态
}

type GoodsIndexPageData struct {
	GoodsItems []GoodsIndexPage `json:"goods_items"` // 商品列表
	PageCount  float64          `json:"page_count"`  // 商品总页数
	Page       int              `json:"page"`        // 第几页
}

type GoodsEdit struct {
	CateId        int     `json:"cate_id"`        // 商品的分类id
	GoodsNumber   int     `json:"goods_number"`   // 商品库存
	IsHot         int     `json:"is_hot"`         // 是否热门商品
	IsBest        int     `json:"is_best"`        // 是否推荐商品
	IsNew         int     `json:"is_new"`         // 是否新的商品
	GoodsTypeId   int     `json:"goods_type_id"`  // 商品关联类型
	Sort          int     `json:"sort"`           // 商品排序
	Status        int     `json:"status"`         // 商品状态
	Price         float64 `json:"price"`          // 商品价格
	MarketPrice   float64 `json:"market_price"`   // 商品市场价格
	Title         string  `json:"title"`          // 商品名
	SubTitle      string  `json:"sub_title"`      // 商品的二级标题
	GoodsSn       string  `json:"goods_sn"`       // 商品的sn号
	RelationGoods string  `json:"relation_goods"` // 关联商品
	GoodsAttr     string  `json:"goods_attr"`     // 商品属性
	GoodsVersion  string  `json:"goods_version"`  // 商品版本
	GoodsImg      string  `json:"goods_img"`      // 商品图片
	GoodsGift     string  `json:"goods_gift"`     // 商品赠品
	GoodsFitting  string  `json:"goods_fitting"`  // 商品配件
	GoodsKeywords string  `json:"goods_keywords"` // 商品关键词
	GoodsDesc     string  `json:"goods_desc"`     // 商品描述
	GoodsContent  string  `json:"goods_content"`  // 商品详情
}

type GoodsImageList struct {
	Id      int    `json:"id"`       // 商品图片id
	GoodsId int    `json:"goods_id"` // 商品id
	ImgUrl  string `json:"img_url"`  // 商品图片路径
}

type GoodsAttrList struct {
	Id              int    `json:"id"`                // 商品属性id
	GoodsId         int    `json:"goods_id"`          // 商品id
	AttributeCateId int    `json:"attribute_cate_id"` // 商品属性对应的cateId
	AttributeId     int    `json:"attribute_id"`      // 商品属性的Id
	AttributeType   int    `json:"attribute_type"`    // 商品属性的类型 1. 单行文本框 2. 多行文本框 3. 下拉框
	AttributeTitle  string `json:"attribute_title"`   // 商品属性的名
	AttributeValue  string `json:"attribute_value"`   // 商品属性的填写内容
}

type GoodsEditPageData struct {
	Goods           GoodsEdit                 `json:"goods"` // 商品详情
	GoodsCateItems  []GoodsCateWithGoodsCate1 `json:"goods_cate_items"`
	GoodsColorItems []GoodsColorList2         `json:"goods_color_items"`
	GoodsTypeItems  []GoodsTypeList1          `json:"goods_type_items"`
	GoodsImageItems []GoodsImageList          `json:"goods_image_items"` // 该商品图库信息
	GoodsAttrItems  []GoodsAttrList           `json:"goods_attr_items"`  // 该商品包装规格信息
}

type NavList struct {
	Id       int    `json:"id"`       // 导航栏id
	Position int    `json:"position"` // 导航栏位置
	Sort     int    `json:"sort"`     // 导航栏排序
	Status   int    `json:"status"`   // 导航栏状态
	Title    string `json:"title"`    // 导航栏名
	Link     string `json:"link"`     // 导航栏跳转链接
	Relation string `json:"relation"` // 相关联的商品id列表
}

type NavIndexPageData struct {
	NavItems  []NavList `json:"nav_items"`  // 导航栏列表
	PageCount float64   `json:"page_count"` // nav页数
	Page      int       `json:"page"`       // 当前页数
}

type Nav1 struct {
	Id        int    `json:"id"`         // 导航栏id
	Position  int    `json:"position"`   // 导航栏位置
	IsOpennew int    `json:"is_opennew"` // 是否打开新窗口
	Sort      int    `json:"sort"`       // 导航栏排序
	Status    int    `json:"status"`     // 导航栏状态
	Title     string `json:"title"`      // 导航栏名
	Link      string `json:"link"`       // 导航栏跳转链接
	Relation  string `json:"relation"`   // 相关联的商品id列表
}

type NavEditPageData struct {
	NavInfo Nav1 `json:"nav_info"` // 导航栏列表
}

type TopNav struct {
	Id        int    `json:"id"`         // 导航栏id
	IsOpennew int    `json:"is_opennew"` // 是否打开新窗口
	Sort      int    `json:"sort"`       // 导航栏排序
	Status    int    `json:"status"`     // 导航栏状态
	Title     string `json:"title"`      // 导航栏名
	Link      string `json:"link"`       // 导航栏跳转链接
	Relation  string `json:"relation"`   // 相关联的商品id列表
}

type MiddleNav struct {
	Id         int     `json:"id"`          // 导航栏id
	IsOpennew  int     `json:"is_opennew"`  // 是否打开新窗口
	Sort       int     `json:"sort"`        // 导航栏排序
	Status     int     `json:"status"`      // 导航栏状态
	Title      string  `json:"title"`       // 导航栏名
	Link       string  `json:"link"`        // 导航栏跳转链接
	Relation   string  `json:"relation"`    // 相关联的商品id列表
	GoodsItems []Goods `json:"goods_items"` // 关联的商品列表
}

type ShopIndexPageData struct {
	*ShopBaseData    `json:"shop_base_data"`
	FocusList        []ResponseFocus `json:"focus_list"`           // 轮播图列表
	CateId1GoodsList []Goods         `json:"cate_id_1_goods_list"` // cateid=1的商品列表
}

type ShopBaseData struct {
	TopNavList                 []TopNav                 `json:"top_nav_list"`    // 顶部导航列表
	GoodsCateWithGoodsCateList []GoodsCateWithGoodsCate `json:"goods_cate_list"` // 商品分类
	MiddleNavList              []MiddleNav              `json:"middle_nav_list"` // 中间导航列表
}

type SearchProductsByKeywordData struct {
	*ShopBaseData `json:"shop_base_data"` // 商城界面的基础数据
	GoodsList     []Goods                 `json:"goods_list"`   // 商品列表
	PageNum       float64                 `json:"page_num"`     // 总页码数量
	CurrentPage   int                     `json:"current_page"` // 当前页码
}
