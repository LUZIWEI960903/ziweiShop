package redis

import (
	"fmt"
	"ziweiShop/models"
)

// CacheTopNavList 查询topNavList 缓存
func CacheTopNavList(oTopNavList *[]models.Nav) bool {
	return CacheDB.Get(fmt.Sprintf("%vtopNavList", CACHEKEYINDEX), oTopNavList)
}

// SetCacheTopNavList 设置topNavList 缓存
func SetCacheTopNavList(oTopNavList *[]models.Nav) {
	CacheDB.Set(fmt.Sprintf("%vtopNavList", CACHEKEYINDEX), oTopNavList, CACHETIME)
}

// CacheFocusList 查询focusList 缓存
func CacheFocusList(focusList *[]models.Focus) bool {
	return CacheDB.Get(fmt.Sprintf("%vfocusList", CACHEKEYINDEX), focusList)
}

// SetCacheFocusList 设置focusList 缓存
func SetCacheFocusList(focusList *[]models.Focus) {
	CacheDB.Set(fmt.Sprintf("%vfocusList", CACHEKEYINDEX), focusList, CACHETIME)
}

// CacheTopGoodsCateWithGoodsCateList 查询oTopGoodsCateWithGoodsCateList 缓存
func CacheTopGoodsCateWithGoodsCateList(oTopGoodsCateWithGoodsCateList *[]models.GoodsCate) bool {
	return CacheDB.Get(fmt.Sprintf("%voTopGoodsCateWithGoodsCateList", CACHEKEYINDEX), oTopGoodsCateWithGoodsCateList)
}

// SetCacheTopGoodsCateWithGoodsCateList 设置oTopGoodsCateWithGoodsCateList 缓存
func SetCacheTopGoodsCateWithGoodsCateList(oTopGoodsCateWithGoodsCateList *[]models.GoodsCate) {
	CacheDB.Set(fmt.Sprintf("%voTopGoodsCateWithGoodsCateList", CACHEKEYINDEX), oTopGoodsCateWithGoodsCateList, CACHETIME)
}

// CacheMiddleNavList 查询middleNavList 缓存
func CacheMiddleNavList(oMiddleNavList *[]models.Nav) bool {
	return CacheDB.Get(fmt.Sprintf("%vmiddleNavList", CACHEKEYINDEX), oMiddleNavList)
}

// SetCacheMiddleNavList 设置middleNavList 缓存
func SetCacheMiddleNavList(oMiddleNavList *[]models.Nav) {
	CacheDB.Set(fmt.Sprintf("%vmiddleNavList", CACHEKEYINDEX), oMiddleNavList, CACHETIME)
}

// CacheGoodsList 查询goodsList 缓存
func CacheGoodsList(goodsList *[]models.Goods) bool {
	return CacheDB.Get(fmt.Sprintf("%vgoodsList", CACHEKEYINDEX), goodsList)
}

// SetCacheGoodsList 设置goodsList 缓存
func SetCacheGoodsList(goodsList *[]models.Goods) {
	CacheDB.Set(fmt.Sprintf("%vgoodsList", CACHEKEYINDEX), goodsList, CACHETIME)
}

// CacheGoodsListByCategory 查询goodsList_X 缓存
func CacheGoodsListByCategory(goodsList1 *[]models.Goods, cateId int) bool {
	return CacheDB.Get(fmt.Sprintf("%vgoodsList_%v", CACHEKEYINDEX, cateId), goodsList1)
}

// SetCacheGoodsListByCategory 设置goodsList_X 缓存
func SetCacheGoodsListByCategory(goodsList *[]models.Goods, cateId int) {
	CacheDB.Set(fmt.Sprintf("%vgoodsList_%v", CACHEKEYINDEX, cateId), goodsList, CACHETIME)
}
