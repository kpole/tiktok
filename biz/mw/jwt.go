package mw

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	db "offer_tiktok/biz/dal/db"
	"offer_tiktok/biz/model/basic/user"
	_ "offer_tiktok/pkg/errno"
	"offer_tiktok/pkg/utils"
	"time"

	"github.com/hertz-contrib/jwt"
)

var JwtMiddleware *jwt.HertzJWTMiddleware

func Init() {
	JwtMiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
		Key:         []byte("tiktok secret key"),
		TokenLookup: "query: token, form: token",
		IdentityKey: "user_id",
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginRequest user.DouyinUserLoginRequest
			if err := c.BindAndValidate(&loginRequest); err != nil {
				return nil, err
			}
			password, err := utils.MD5(loginRequest.Password)
			user_id, err := db.VerifyUser(loginRequest.Username, password)
			if err != nil {
				return nil, err
			}
			c.Set("user_id", user_id)
			return user_id, nil
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
			c.Set("token", token)
		},
	})

}
