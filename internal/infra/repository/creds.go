package repository

import "time"

type Creds struct {
	ConnectionString   string
	MaxOpenConns       int
	MaxIdleConns       int
	MaxConnLifetimeSec time.Duration
}
