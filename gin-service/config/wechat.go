package config

type WechatConfig struct {
	AppID     string `env:"WECHAT_APP_ID"`
	AppSecret string `env:"WECHAT_APP_SECRET"`
	Token string `env:"WECHAT_TOKEN"`
	EncodingAESKey string `env:"WECHAT_ENCODING_AES_KEY"`
}
