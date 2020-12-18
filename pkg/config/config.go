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

type LogOptions struct {
	FilePath         string `json:"file-path" yaml:"file-path"`
	Level            string `json:"level" yaml:"level"`
	TimeFormatter    string `json:"time-formatter" yaml:"time-formatter"`
	DisableTimestamp bool   `json:"disable-timestamp" yaml:"disable-timestamp"`
}

type TLSOptions struct {
	CertFile   string `json:"cert-file" yaml:"cert-file"`
	KeyFile    string `json:"key-file" yaml:"key-file"`
	ServerName string `json:"server-name" yaml:"server-name"`
}

type Config struct {
	MySQL  MySQLOptions     `json:"mysql" yaml:"mysql"`
	Server DAOServerOptions `json:"server" yaml:"server"`
	Client DAOClientOptions `json:"client" yaml:"client"`
	Log    LogOptions       `json:"log" yaml:"log"`
	TLS    TLSOptions       `json:"tls" yaml:"tls"`
	PWD    string
}
