package mw

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	db "offer_tiktok/biz/dal/db"
	"offer_tiktok/biz/model/basic/user"
	_ "offer_tiktok/biz/model/basic/user"
	_ "offer_tiktok/pkg/errno"
	"time"

	"github.com/hertz-contrib/jwt"
)

var JwtMiddleware *jwt.HertzJWTMiddleware

func Init() {
	JwtMiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
		Key:         []byte("tiktok secret key"),
		Timeout:     12 * time.Hour,
		MaxRefresh:  12 * time.Hour,
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		IdentityKey: "user_id",
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginStruct struct {
				username string `form:"account" json:"account" query:"username" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
				password string `form:"password" json:"password" query:"password" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
			}
			if err := c.BindAndValidate(&loginStruct); err != nil {
				return nil, err
			}
			user, err := db.QueryUser(loginStruct.username)
			if err != nil {
				return nil, err
			}
			if user != nil {
				return nil, err
			}

			return loginStruct.username, nil
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(string); ok {
				return jwt.MapClaims{
					jwt.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},

		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {

			c.JSON(http.StatusOK, user.DouyinUserLoginResponse{
				StatusCode: 1,
				UserId:     123,
				Token:      token,
			})
		},
	})

}
