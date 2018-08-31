package db

import (
	"context"
)

var driver Driver

// Register makes an database driver available
func Register(d Driver) {
	if d == nil {
		panic("gateway/db: Register driver is nil")
	}
	driver = d
}

func Open(ctx context.Context, connString string) (Conn, error) {
	if driver == nil {
		panic("gateway/db: Fail to open due to driver is nil")
	}
	return driver.Open(ctx, connString)
}