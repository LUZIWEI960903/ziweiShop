package redis

import (
	"fmt"
	"ziweiShop/models"
)

// CacheTopNavList 查询topNavList 缓存
func CacheTopNavList(oTopNavList *[]models.Nav) bool {
	return CacheDB.Get("topNavList", oTopNavList)
}

// SetCacheTopNavList 设置topNavList 缓存
func SetCacheTopNavList(oTopNavList *[]models.Nav) {
	CacheDB.Set("topNavList", oTopNavList, 60*60)
}

// CacheFocusList 查询focusList 缓存
func CacheFocusList(focusList *[]models.Focus) bool {
	return CacheDB.Get("focusList", focusList)
}

// SetCacheFocusList 设置focusList 缓存
func SetCacheFocusList(focusList *[]models.Focus) {
	CacheDB.Set("focusList", focusList, 60*60)
}

// CacheTopGoodsCateWithGoodsCateList 查询oTopGoodsCateWithGoodsCateList 缓存
func CacheTopGoodsCateWithGoodsCateList(oTopGoodsCateWithGoodsCateList *[]models.GoodsCate) bool {
	return CacheDB.Get("oTopGoodsCateWithGoodsCateList", oTopGoodsCateWithGoodsCateList)
}

// SetCacheTopGoodsCateWithGoodsCateList 设置oTopGoodsCateWithGoodsCateList 缓存
func SetCacheTopGoodsCateWithGoodsCateList(oTopGoodsCateWithGoodsCateList *[]models.GoodsCate) {
	CacheDB.Set("oTopGoodsCateWithGoodsCateList", oTopGoodsCateWithGoodsCateList, 60*60)
}

// CacheMiddleNavList 查询middleNavList 缓存
func CacheMiddleNavList(oMiddleNavList *[]models.Nav) bool {
	return CacheDB.Get("middleNavList", oMiddleNavList)
}

// SetCacheMiddleNavList 设置middleNavList 缓存
func SetCacheMiddleNavList(oMiddleNavList *[]models.Nav) {
	CacheDB.Set("middleNavList", oMiddleNavList, 60*60)
}

// CacheGoodsList 查询goodsList 缓存
func CacheGoodsList(goodsList *[]models.Goods) bool {
	return CacheDB.Get("goodsList", goodsList)
}

// SetCacheGoodsList 设置goodsList 缓存
func SetCacheGoodsList(goodsList *[]models.Goods) {
	CacheDB.Set("goodsList", goodsList, 60*60)
}

// CacheGoodsListByCategory 查询goodsList_X 缓存
func CacheGoodsListByCategory(goodsList1 *[]models.Goods, cateId int) bool {
	return CacheDB.Get(fmt.Sprintf("goodsList_%v", cateId), goodsList1)
}

// SetCacheGoodsListByCategory 设置goodsList_X 缓存
func SetCacheGoodsListByCategory(goodsList *[]models.Goods, cateId int) {
	CacheDB.Set(fmt.Sprintf("goodsList_%v", cateId), goodsList, 60*60)
}
