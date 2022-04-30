package models

type Setting struct {
	Id              int    // 设置的id
	OssStatus       int    // 网站的oss状态
	SiteTitle       string // 网站的标题
	SiteLogo        string // 网站的logo
	SiteKeywords    string // 网站的关键词
	SiteDescription string // 网站的描述
	DefaultPic      string // 没图片时网站的默认图片
	SiteIcp         string // 网站的icp
	SiteTel         string // 网站的联系电话
	SearchKeywords  string // 网站的搜索关键词
	TongjiCode      string // 网站的统计代码
	AppId           string // 网站的统appid
	AppSecret       string // 网站的app密钥
	EndPoint        string // 网站的endpoint
	BucketName      string // 网站的bucketname
}

func (Setting) TableName() string {
	return "setting"
}
