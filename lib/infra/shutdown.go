package infra

import (
	"context"

	"_gdc_/lib/infra/mysql"
)

func Shutdown(ctx context.Context) error {
	_ = mysql.Shutdown(MysqlClient)

	return nil
}
