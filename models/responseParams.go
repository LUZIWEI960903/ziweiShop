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
