package config

import "time"

type All struct {
	App    App    `yaml:"App"`
	Log    Log    `yaml:"Log"`
	MySQL  MySQL  `yaml:"MySQL"`
	Page   Page   `yaml:"Page"`
	Server Server `yaml:"Server"`
	Email  Email  `yaml:"Email"`
	Token  Token  `yaml:"Token"`
	Redis  Redis  `yaml:"Redis"`
	Rule   Rule   `yaml:"Rule"`
}
type App struct {
	Name      string    `yaml:"Name"`
	StartTime time.Time `yaml:"StartTime"`
	Version   string    `yaml:"version"`
}

type Log struct {
	Level         string `yaml:"Level"`
	LogSavePath   string `yaml:"LogSavePath"`
	LowLevelFile  string `yaml:"LowLevelFile"`
	LogFileExt    string `yaml:"LogFileExt"`
	MaxSize       int    `yaml:"MaxSize"`
	MaxAge        int    `yaml:"MaxAge"`
	MaxBackups    int    `yaml:"MaxBackups"`
	Compress      bool   `yaml:"Compress"`
	HighLevelFile string `yaml:"HighLevelFile"`
}

type MySQL struct {
	Host         string `yaml:"Host"`
	User         string `yaml:"User"`
	Password     string `yaml:"Password"`
	DbName       string `yaml:"Dbname"`
	Port         string `yaml:"Port"`
	MaxOpenConns int    `yaml:"MaxOpenConns"`
	MaxIdleConns int    `yaml:"MaxIdleConns"`
}

type Page struct {
	DefaultPageSize int32  `yaml:"DefaultPageSize"` //默认每页所放置的个数
	MaxPageSize     int32  `yaml:"MaxPageSize"`     //每页最大放多少个
	GetSizeKey      string `yaml:"GetSizeKey"`      //根据key去获取当前放置的个数
	GetPageKey      string `yaml:"GetPageKey"`      //根据key去获取当前页数
}

type Server struct {
	Addr                  string        `yaml:"Port"`
	Mode                  string        `yaml:"Mode"`
	ReadTimeOut           time.Duration `yaml:"ReadTimeOut"`
	WriteTimeOut          time.Duration `yaml:"WriteTimeOut"`
	DefaultContextTimeout time.Duration `yaml:"DefaultContextTimeout"`
}
type Email struct {
	ValidEmail string `yaml:"ValidEmail"`
	SmtpHost   string `yaml:"SmtpHost"`
	SmtpEmail  string `yaml:"SmtpEmail"` //你的邮箱
	SmtpPass   string `yaml:"SmtpPass"`  //你邮箱的通行码
}

type Token struct {
	Key                  string        `yaml:"Key"`
	AssessTokenDuration  time.Duration `yaml:"AssessTokenDuration"`
	RefreshTokenDuration time.Duration `yaml:"RefreshTokenDuration"`
	AuthorizationKey     string        `yaml:"AuthorizationKey"`
	AuthorizationType    string        `yaml:"AuthorizationType"`
	ExpireTime           int           `yaml:"ExpireTime"`
}

type Redis struct {
	Address   string        `yaml:"Address"`
	DB        int           `yaml:"DB"`
	Password  string        `yaml:"Password"`
	PoolSize  int           `yaml:"PoolSize"`
	CacheTime time.Duration `yaml:"CacheTime"`
}

type Rule struct {
	MaxUsernameLen int    `yaml:"MaxUsernameLen"`
	MinUsernameLen int    `yaml:"MinUsernameLen"`
	MaxPasswordLen int    `yaml:"MaxPasswordLen"`
	MinPasswordLen int    `yaml:"MinPasswordLen"`
	Avatar         string `yaml:"Avatar"`
}
