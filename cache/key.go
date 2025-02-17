package cache

import (
	"fmt"
	"strconv"
)

const(
	RankKey = "rank"
)


func ProductViewKey(id uint) string{
	return fmt.Sprintf("view:procduct:%s",strconv.Itoa(int(id)))
}

func SeckillProductKey(id uint) string{
	return fmt.Sprintf("seckill:product:%s",strconv.Itoa(int(id)))
}