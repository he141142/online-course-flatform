package configs

import "time"

type AppConfig struct {
	ShutdownTimeout time.Duration
	AppName         string
	Appversion      string
}

type HttpConfig struct {
	HostName string
	Port     int
	IPAddr   string
}