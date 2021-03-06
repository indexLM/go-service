package config

type Server struct {
	System System `json:"system" yaml:"system"`
	Redis  Redis  `json:"redis" yaml:"redis"`
	Log    Log    `json:"log" yaml:"log"`
	Jwt    Jwt    `json:"jwt" yaml:"jwt"`
}

// System 系统配置
type System struct {
	Port int    `json:"port" yaml:"port"`
	Name string `json:"name" yaml:"name"`
}

// Redis 缓存配置
type Redis struct {
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	Db       int    `json:"db" yaml:"db"`
	Password string `json:"password" yaml:"password"`
	PoolSize int    `json:"poolSize" yaml:"poolSize"`
	Cache    struct {
		TokenExpired int `json:"tokenExpired" yaml:"tokenExpired"`
	}
}

// Log 日志配置
type Log struct {
	Prefix  string `json:"prefix" yaml:"prefix"`
	LogFile bool   `json:"log_file" yaml:"log-file"`
	Stdout  string `json:"stdout" yaml:"stdout"`
	File    string `json:"file" yaml:"file"`
}

// Jwt 签名结构
type Jwt struct {
	SignKey string `json:"signKey" yaml:"signKey"`
	Expires int64  `json:"expires" yaml:"expires"`
}
