// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package wire

import (
	"dev/wire/internal/dao"
	"dev/wire/internal/server"
	"dev/wire/internal/service"
	"github.com/google/wire"
)

func InitializeServer(daos dao.Dao) (*server.Server, error) {
	wire.Build(service.NewService, server.NewServer)
	return nil, nil
}
func InitDB() (dao.Dao, error) {
	wire.Build(dao.NewDB, dao.NewDao)
	return nil, nil
}
