package cache

import "strconv"

func View(productID uint) uint64 {
	countStr, _ := RedisClient.Get(ProductViewKey(productID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

func AddView(productID uint) {
	//Increase product click
    RedisClient.Incr(ProductViewKey(productID))
	RedisClient.ZIncrBy(RankKey, 1, strconv.Itoa((int(productID))))

}