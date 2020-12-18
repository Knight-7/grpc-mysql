package config

import "fmt"

func (c *Config) GetServerAddr() string {
	addr := fmt.Sprintf("%s:%d", c.Server.Address, c.Server.Port)
	return addr
}

func (c *Config) GetClientTarget() string {
	addr := fmt.Sprintf("%s:%d", c.Client.Address, c.Client.Port)
	return addr
}

func (c *Config) GetMySQLDsn() string {
	user := c.MySQL.UserName
	passwd := c.MySQL.Password
	addr := c.MySQL.Address
	port := c.MySQL.Port
	name := c.MySQL.Database

	mysqlDsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
		user, passwd, addr, port, name)

	return mysqlDsn
}

func (c *Config) GetLogLevel() string {
	return c.Log.Level
}

func (c *Config) GetLogFilePath() string {
	return c.PWD + c.Log.FilePath
}

func (c *Config) GetLogTimeFormatter() string {
	return c.Log.TimeFormatter
}

func (c *Config) GetLogDisableTimestamp() bool {
	return c.Log.DisableTimestamp
}

func (c *Config) GetCertFile() string {
	return c.PWD + "keys/" + c.TLS.CertFile
}

func (c *Config) GetKeyFile() string {
	return c.PWD + "keys/" + c.TLS.KeyFile
}

func (c *Config) GetServerName() string {
	return c.TLS.ServerName
}
