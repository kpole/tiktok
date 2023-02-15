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

// path为数据库中存放的url，需要转为前端可访问的完整url，url为发送get请求中携带的URL去掉path的部分
func URLconvert(ctx context.Context, c *app.RequestContext, path string) (fullURL string) {
	arr := strings.Split(path, "/")
	u, err := minio.GetObjURL(ctx, arr[0], arr[1])
	if err != nil {
		hlog.CtxInfof(ctx, err.Error())
		return ""
	}
	fullURL = string(c.URI().Scheme()) + "://" + string(c.URI().Host()) + "/src" + u.Path
	return fullURL
}