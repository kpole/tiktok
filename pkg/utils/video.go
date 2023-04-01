package utils

import (
	"context"
	"fmt"
	"offer_tiktok/biz/mw/minio"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

// NewFileName Splicing user_id and time to make unique filename
func NewFileName(user_id, time int64) string {
	return fmt.Sprintf("%d.%d", user_id, time)
}

// URLconvert Convert the path in the database into a complete url accessible by the front end
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
