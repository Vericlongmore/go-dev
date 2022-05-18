package utils

import (
	"time"
)

const (
	DefaultGrpcDialTimeout       = 5 * time.Second
	DefaultGrpcMinConnectTimeout = 3 * time.Second
	DefaultGrpcKeepAliveTime     = 10 * time.Second
	DefaultGrpcKeepAliveTimeout  = 3 * time.Second
)
