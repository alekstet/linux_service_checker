package maker

import (
	"sync"

	"github.com/alekstet/linux_service_checker/server/conf"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/ssh"
)

var _ Maker = (*makerImpl)(nil)

type makerImpl struct {
	config  *conf.Config
	client  *ssh.Client
	mutex   sync.Mutex
	dbPool  *pgxpool.Pool
	isAlive bool
}

func NewMaker(config *conf.Config, dbPool *pgxpool.Pool, client *ssh.Client) *makerImpl {
	return &makerImpl{
		dbPool: dbPool,
		config: config,
		mutex:  sync.Mutex{},
		client: client,
	}
}
