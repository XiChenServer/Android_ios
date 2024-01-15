/**
* @Auth:ShenZ
* @Description:
* @CreateDate:2022/06/15 14:57:55
 */
package models

import (
	"Android_ios/dao"
	"context"

	"time"
)

/*
*
设置在线用户到redis缓存
*
*/
func SetUserOnlineInfo(key string, val []byte, timeTTL time.Duration) {
	ctx := context.Background()
	dao.RDB.Set(ctx, key, val, timeTTL)
}
