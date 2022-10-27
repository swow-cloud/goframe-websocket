package service

import (
	"context"
	"time"

	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"

	"goframe-websocket/internal/model"
)

var authService *jwt.GfJWTMiddleware

func Auth() *jwt.GfJWTMiddleware {
	return authService
}

func init() {

	var (
		cache       = gcache.New()
		ctx         = context.Background()
		redisConfig = &gredis.Config{
			Address: "127.0.0.1:6379", //todo 2022-10-27 待替换redisConfig
			Db:      0,
		}
	)
	gVar, _ := g.Cfg().Get(ctx, "jwt.key")
	key := gVar.Val().(string)
	// Create redis client object.
	redis, err := gredis.New(redisConfig)
	if err != nil {
		panic(err)
	}
	// Create redis cache adapter and set it to cache object.
	cache.SetAdapter(gcache.NewAdapterRedis(redis))
	// Create jwt
	auth := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte(key),
		Timeout:         time.Minute * 24 * 60,
		MaxRefresh:      time.Minute * 5,
		IdentityKey:     "id",
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Authenticator:   Authenticator,
		Unauthorized:    Unauthorized,
		PayloadFunc:     PayloadFunc,
		IdentityHandler: IdentityHandler,
		CacheAdapter:    cache,
	})
	authService = auth
}

// PayloadFunc is a callback function that will be called during login.
// Using this function it is possible to add additional payload data to the webtoken.
// The data is then made available during requests via c.Get("JWT_PAYLOAD").
// Note that the payload is not encrypted.
// The attributes mentioned on jwt.io can't be used as keys for the map.
// Optional, by default no additional data will be set.
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

// IdentityHandler get the identity from JWT and set the identity for every request
// Using this function, by r.GetParam("id") get identity
func IdentityHandler(ctx context.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return claims[authService.IdentityKey]
}

// Unauthorized is used to define customized Unauthorized callback function.
func Unauthorized(ctx context.Context, code int, message string) {
	r := g.RequestFromCtx(ctx)
	r.Response.WriteJson(g.Map{
		"code":    code,
		"message": message,
	})
	r.ExitAll()
}

// Authenticator is used to validate login parameters.
// It must return user data as user identifier, it will be stored in Claim Array.
// if your identityKey is 'id', your user data must have 'id'
// Check error (e) to determine the appropriate error message.
func Authenticator(ctx context.Context) (interface{}, error) {
	var (
		r  = g.RequestFromCtx(ctx)
		in model.UserLoginInput
	)
	//这里会在request获取bodyContent并且赋值到in
	if err := r.Parse(&in); err != nil {
		return "", err
	}
	if user, _ := User().GetUserByMobileAndPassword(ctx, in.Mobile, in.Password); user != nil {
		return g.Map{
			"id":     user.Id,
			"mobile": user.Mobile,
		}, nil
	}

	return nil, jwt.ErrFailedAuthentication
}
