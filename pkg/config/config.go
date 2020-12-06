package config

type MySQLOptions struct {
	Address      string `json:"address" yaml:"address"`
	Port         int    `json:"port" yaml:"port"`
	Database     string `json:"database" yaml:"database"`
	UserName     string `json:"username" yaml:"username"`
	Password     string `json:"password" yaml:"password"`
	MaxOpenConns int    `json:"max-open-connections" yaml:"max-open-connections"`
	MaxIdleConns int    `json:"max-idle-connections" yaml:"max-idle-connections"`
	MaxLifetime  int    `json:"max-connections-lifetime" yaml:"max-connections-lifetime"`
}

type DAOServerOptions struct {
	Address string `json:"address" yaml:"address"`
	Port    int    `json:"port" yaml:"port"`
}

type DAOClientOptions struct {
	Address string `json:"address" address:"address"`
	Port    int    `json:"port" port:"port"`
}

type Config struct {
	MySQL  MySQLOptions     `json:"mysql" yaml:"mysql"`
	Server DAOServerOptions `json:"server" yaml:"server"`
	Client DAOClientOptions `json:"client" yaml:"client"`
}
