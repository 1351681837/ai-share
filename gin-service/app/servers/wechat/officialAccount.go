package wechat

import (
	"context"
	"fmt"
	"github.com/putyy/ai-share/config"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/user"
	"log"
)

//Config config for 微信公众号
type Config struct {
	AppID          string `json:"app_id"`           //appid
	AppSecret      string `json:"app_secret"`       //appsecret
	Token          string `json:"token"`            //token
	EncodingAESKey string `json:"encoding_aes_key"` //EncodingAESKey
	Cache          cache.Cache
}

var officialApp *officialaccount.OfficialAccount

func init() {
	// 使用memcache保存access_token，也可选择redis或自定义cache
	wc := wechat.NewWechat()
	redis := config.Redis
	redisOpts := &cache.RedisOpts{
		Host:        redis.Host,     // redis host
		Password:    redis.Password, //redis password
		Database:    redis.Database, // redis db
		MaxActive:   10,             // 连接池最大活跃连接数
		MaxIdle:     10,             //连接池最大空闲连接数
		IdleTimeout: 60,             //空闲连接超时时间，单位：second
	}

	ctx := context.Background()
	redisCache := cache.NewRedis(ctx, redisOpts)
	wechatCof := config.Wechat
	fmt.Println(wechatCof)
	cfg := &offConfig.Config{
		AppID:     wechatCof.AppID,
		AppSecret: wechatCof.AppSecret,
		Token:     wechatCof.Token,
		//EncodingAESKey:  wechat.EncodingAESKey,
		Cache: redisCache,
	}
	officialApp = wc.GetOfficialAccount(cfg)
}

func OfficialApp() *officialaccount.OfficialAccount {
	return officialApp
}

func GetUserInfo(openID string) (userInfo *user.Info) {
	userInfo, err := officialApp.GetUser().GetUserInfo(openID)
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	return userInfo
}
