package models

type LoginParams struct {
	Username     string `json:"username"`      // 用户名
	Password     string `json:"password"`      // 用户密码
	CaptchaId    string `json:"captcha_id"`    // 验证码id
	CaptchaValue string `json:"captcha_value"` // 验证码的值
}

type AddRoleParams struct {
	Title       string `json:"title"`       // 角色名
	Description string `json:"description"` // 角色详情
}

type EditRoleParams struct {
	Id          int    `json:"id"`          // role的id
	Title       string `json:"title"`       // 角色名
	Description string `json:"description"` // 角色详情
}

type AddManagerParams struct {
	RoleId   int    `json:"role_id"`  // 管理员所属角色（部门）
	Username string `json:"username"` // 管理员姓名
	Password string `json:"password"` // 管理员密码
	Mobile   string `json:"mobile"`   // 管理员手机
	Email    string `json:"email"`    // 管理员邮箱
}

type EditManagerParams struct {
	Id       int    `json:"id"`       // 管路员id
	RoleId   int    `json:"role_id"`  // 管理员所属角色（部门）
	Password string `json:"password"` // 管理员密码
	Mobile   string `json:"mobile"`   // 管理员手机
	Email    string `json:"email"`    // 管理员邮箱
}

type AddAccessParams struct {
	Type        int    `json:"type"`        // 模块类型: 1.表示模块 2.表示菜单 3.表示操作
	ModuleId    int    `json:"module_id"`   // 与当前Id自关联
	Sort        int    `json:"sort"`        // 排序
	Status      int    `json:"status"`      // 模块状态
	ActionName  string `json:"action_name"` // 操作名称
	ModuleName  string `json:"module_name"` // 模块名称
	Url         string `json:"url"`         // 路由跳转地址
	Description string `json:"description"` // 描述
}

type EditAccessParams struct {
	Id          int    `json:"id"`          // 权限id
	Type        int    `json:"type"`        // 模块类型: 1.表示模块 2.表示菜单 3.表示操作
	ModuleId    int    `json:"module_id"`   // 与当前Id自关联
	Sort        int    `json:"sort"`        // 排序
	Status      int    `json:"status"`      // 模块状态
	ActionName  string `json:"action_name"` // 操作名称
	ModuleName  string `json:"module_name"` // 模块名称
	Url         string `json:"url"`         // 路由跳转地址
	Description string `json:"description"` // 描述
}

type DoAuthParams struct {
	Id         int                         `json:"role_id"`     // 角色id
	AccessItem []ResponseTopAccessItemAuth `json:"access_item"` // 权限列表
}

type AddFocusParams struct {
	FocusType int    `json:"focus_type"` // 轮播图类型 1.web 2.app 3.小程序
	Sort      int    `json:"sort"`       // 排序方式
	Title     string `json:"title"`      // 轮播图名称
	Link      string `json:"links"`      // 跳转地址
}

type EditFocusParams struct {
	Id        int    `json:"id"`         // 轮播图id
	FocusType int    `json:"focus_type"` // 轮播图类型 1.web 2.app 3.小程序
	Sort      int    `json:"sort"`       // 排序方式
	Status    int    `json:"status"`     // 状态
	Title     string `json:"title"`      // 轮播图名称
	Link      string `json:"link"`       // 跳转地址
}

type AddGoodsCateParams struct {
	Pid         int    `json:"pid"`         // 二级商品分类id与Id自关联
	Status      int    `json:"status"`      // 商品分类状态
	Sort        int    `json:"sort"`        // 商品分类排序
	Title       string `json:"title"`       // 商品分类名
	Link        string `json:"link"`        // 商品分类跳转地址
	Template    string `json:"template"`    // 商品分类加载的模板
	SubTitle    string `json:"sub_title"`   // 商品分类Seo标题
	Keywords    string `json:"keywords"`    // 商品分类Seo关键词
	Description string `json:"description"` // 商品分类描述
}

type EditGoodsCateParams struct {
	Id          int    `json:"id"`          // 商品分类id
	Pid         int    `json:"pid"`         // 二级商品分类id与Id自关联
	Status      int    `json:"status"`      // 商品分类状态
	Sort        int    `json:"sort"`        // 商品分类排序
	Title       string `json:"title"`       // 商品分类名
	Link        string `json:"link"`        // 商品分类跳转地址
	Template    string `json:"template"`    // 商品分类加载的模板
	SubTitle    string `json:"sub_title"`   // 商品分类Seo标题
	Keywords    string `json:"keywords"`    // 商品分类Seo关键词
	Description string `json:"description"` // 商品分类描述
}

type AddGoodsTypeParams struct {
	Status      int    `json:"status"`      // 商品类型状态
	Title       string `json:"title"`       // 商品类型名
	Description string `json:"description"` // 商品类型描述
}

type EditGoodsTypeParams struct {
	Id          int    `json:"id"`          // 商品类型id
	Status      int    `json:"status"`      // 商品类型状态
	Title       string `json:"title"`       // 商品类型名
	Description string `json:"description"` // 商品类型描述
}

type AddGoodsTypeAttributeParams struct {
	CateId    int    `json:"cate_id"`    // 对应的商品类型id
	Status    int    `json:"status"`     // 商品类型属性的状态
	Sort      int    `json:"sort"`       // 商品类型属性的排序
	AttrType  int    `json:"attr_type"`  // 商品类型属性的录入方式 1. 单行文本框 2. 多行文本框 3. 从下面的列表中选择（一行代表一个可选值 ）
	Title     string `json:"title"`      // 商品类型属性名
	AttrValue string `json:"attr_value"` // 商品类型属性的录入值
}

type EditGoodsTypeAttributeParams struct {
	Id        int    `json:"id"`         // 商品类型属性的id
	CateId    int    `json:"cate_id"`    // 对应的商品类型id
	Status    int    `json:"status"`     // 商品类型属性的状态
	Sort      int    `json:"sort"`       // 商品类型属性的排序
	AttrType  int    `json:"attr_type"`  // 商品类型属性的录入方式 1. 单行文本框 2. 多行文本框 3. 从下面的列表中选择（一行代表一个可选值 ）
	Title     string `json:"title"`      // 商品类型属性名
	AttrValue string `json:"attr_value"` // 商品类型属性的录入值
}

type AddGoodsColorParams struct {
	ColorName  string `json:"color_name"`  // 商品颜色名
	ColorValue string `json:"color_value"` // 商品颜色值
}

type EditGoodsColorParams struct {
	Id         int    `json:"id"`          // 商品颜色id
	Status     int    `json:"status"`      // 商品颜色状态
	ColorName  string `json:"color_name"`  // 商品颜色名
	ColorValue string `json:"color_value"` // 商品颜色值
}

type AddGoodsParams struct {
	CateId        int     `json:"cate_id"`        // 商品的分类id
	GoodsNumber   int     `json:"goods_number"`   // 商品库存
	IsHot         int     `json:"is_hot"`         // 是否热门商品
	IsBest        int     `json:"is_best"`        // 是否推荐商品
	IsNew         int     `json:"is_new"`         // 是否新的商品
	GoodsTypeId   int     `json:"goods_type_id"`  // 商品关联类型
	Sort          int     `json:"sort"`           // 商品排序
	Status        int     `json:"status"`         // 商品状态
	AddTime       int     `json:"add_time"`       // 商品增加时间
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
	GoodsColor    string  `json:"goods_color"`    // 商品颜色
	GoodsKeywords string  `json:"goods_keywords"` // 商品关键词
	GoodsDesc     string  `json:"goods_desc"`     // 商品描述
	GoodsContent  string  `json:"goods_content"`  // 商品详情
}

type EditGoodsParams struct {
	Id            int     `json:"id"`             // 商品id
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
	GoodsColor    string  `json:"goods_color"`    // 商品颜色
	GoodsKeywords string  `json:"goods_keywords"` // 商品关键词
	GoodsDesc     string  `json:"goods_desc"`     // 商品描述
	GoodsContent  string  `json:"goods_content"`  // 商品详情
}
