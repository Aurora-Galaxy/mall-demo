package cache

import (
	"fmt"
	"strconv"
)

// ProductKeyView 标识商品的点击次数,redis以key—val存储，使用该key标识商品点击数
func ProductKeyView(id uint) string {
	return fmt.Sprintf("view:product:%s", strconv.Itoa(int(id)))
}
