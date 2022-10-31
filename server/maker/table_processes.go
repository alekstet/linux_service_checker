package maker

import (
	"context"
	"log"
)

func (impl *makerImpl) checkEmptyTable() bool {
	pool, err := impl.dbPool.Acquire(context.Background())
	if err != nil {
		log.Println(err)
		return false
	}

	defer pool.Release()

	var quentity int
	err = pool.QueryRow(context.Background(), "SELECT count(*) FROM services").Scan(&quentity)
	if err != nil {
		log.Println(err)
		return false
	}

	if quentity == 0 {
		return true
	}

	return false
}

func (impl *makerImpl) setTable() {
	for _, service := range impl.config.MonitoringServer.ServicesNames {
		insertStatement := "INSERT INTO services (name, description, active, loaded, journal) values ($1, $2, $3, $4, $5)"
		_, err := impl.dbPool.Exec(context.Background(), insertStatement, service, "", "", "", "")
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func (impl *makerImpl) TruncateTable() {
	_, err := impl.dbPool.Exec(context.Background(), "TRUNCATE services")
	if err != nil {
		log.Println(err)
		return
	}
}
