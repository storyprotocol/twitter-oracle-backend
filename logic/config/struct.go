package config

type Config struct {
	Debug   *DebugConfig
	Event   *EventConfig
	Twitter *TwitterConfig
	Wallet  *WalletConfig
}

type DebugConfig struct {
	Enable  bool
	Verbose bool
}

type EventConfig struct {
	Endpoint        string
	TwitterHookAddr string
}

type TwitterConfig struct {
	BearerToken       string
	ApiKey            string
	ApiSecret         string
	AccessToken       string
	AccessTokenSecret string
}

type WalletConfig struct {
	PrivateKey string
}
