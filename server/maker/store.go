package maker

import (
	"sync"

	"github.com/alekstet/linux_service_checker/server/conf"
	"github.com/alekstet/linux_service_checker/server/db"
	ssh2 "github.com/alekstet/linux_service_checker/server/ssh"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/ssh"
)

var _ Maker = (*makerImpl)(nil)

type makerImpl struct {
	config *conf.Config
	client *ssh.Client
	mutex  sync.Mutex
	dbPool *pgxpool.Pool
}

func NewMaker(config *conf.Config) (*makerImpl, error) {
	client, err := ssh2.GetClient(config)
	if err != nil {
		return nil, err
	}

	connectionPool, err := db.GetDBConnectionPool(config.Database.ConnectionString)
	if err != nil {
		return nil, err
	}

	return &makerImpl{
		dbPool: connectionPool,
		config: config,
		mutex:  sync.Mutex{},
		client: client,
	}, nil
}
