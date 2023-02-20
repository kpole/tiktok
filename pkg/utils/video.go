package utils

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"offer_tiktok/biz/mw/minio"
	"strings"
)

func NewFileName(user_id int64, time int64) string {
	return fmt.Sprintf("%d.%d", user_id, time)
}

// URLconvert
/**
 * @description: 将数据库中存放的url转换为前端可访问的完整url
 * @param {context.Context} ctx
 * @param {*app.RequestContext} c
 * @param {string} path
 * @return {string} fullURL
 */
func URLconvert(ctx context.Context, c *app.RequestContext, path string) (fullURL string) {
	if len(path) == 0 {
		return ""
	}
	arr := strings.Split(path, "/")
	u, err := minio.GetObjURL(ctx, arr[0], arr[1])
	if err != nil {
		hlog.CtxInfof(ctx, err.Error())
		return ""
	}
	u.Scheme = string(c.URI().Scheme())
	u.Host = string(c.URI().Host())
	u.Path = "/src" + u.Path
	return u.String()
}
