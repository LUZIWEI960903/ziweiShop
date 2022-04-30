package models

type Setting struct {
	Id              int    `json:"id" form:"id"`                             // 设置的id
	OssStatus       int    `json:"oss_status" form:"oss_status"`             // 网站的oss状态
	SiteTitle       string `json:"site_title" form:"site_title"`             // 网站的标题
	SiteLogo        string `json:"site_logo"`                                // 网站的logo
	SiteKeywords    string `json:"site_keywords" form:"site_keywords"`       // 网站的关键词
	SiteDescription string `json:"site_description" form:"site_description"` // 网站的描述
	DefaultPic      string `json:"default_pic"`                              // 没图片时网站的默认图片
	SiteIcp         string `json:"site_icp" form:"site_icp"`                 // 网站的icp
	SiteTel         string `json:"site_tel" form:"site_tel"`                 // 网站的联系电话
	SearchKeywords  string `json:"search_keywords" form:"search_keywords"`   // 网站的搜索关键词
	TongjiCode      string `json:"tongji_code" form:"tongji_code"`           // 网站的统计代码
	AppId           string `json:"app_id" form:"app_id"`                     // 网站的统appid
	AppSecret       string `json:"app_secret" form:"app_secret"`             // 网站的app密钥
	EndPoint        string `json:"end_point" form:"end_point"`               // 网站的endpoint
	BucketName      string `json:"bucket_name" form:"bucket_name"`           // 网站的bucketname
}

func (Setting) TableName() string {
	return "setting"
}
