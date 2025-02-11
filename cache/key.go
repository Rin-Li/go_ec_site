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