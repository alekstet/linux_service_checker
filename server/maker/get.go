package maker

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func (impl *makerImpl) Get(active string) (*ServicesInfo, error) {
	rows, err := impl.getRows(active)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	servicesInfo := make(ServicesInfo)

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Println(err)
			return nil, err
		}

		var serviceInfo ServiceInfo

		serviceInfo.Name = values[0].(string)
		serviceInfo.Description = values[1].(string)
		serviceInfo.Loaded = values[2].(string)
		serviceInfo.Active = values[3].(string)
		serviceInfo.Journal = values[4].(string)

		servicesInfo[serviceInfo.Name] = serviceInfo

	}

	if impl.collectError != nil {
		return nil, fmt.Errorf("backend is not alive")
	}

	return &servicesInfo, nil
}

func (impl *makerImpl) getRows(active string) (pgx.Rows, error) {
	var rows pgx.Rows

	pool, err := impl.dbPool.Acquire(context.Background())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer pool.Release()

	switch active {
	case "all":
		selectStatement := `SELECT * FROM services`
		rows, err = pool.Query(context.Background(), selectStatement)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	case "active":
		selectStatement := `SELECT * FROM services WHERE active=$1 OR active=$2`
		rows, err = pool.Query(context.Background(), selectStatement, "active(exited)", "active(running)")
		if err != nil {
			log.Println(err)
			return nil, err
		}
	case "inactive":
		selectStatement := `SELECT * FROM services WHERE active=$1`
		rows, err = pool.Query(context.Background(), selectStatement, "inactive(dead)")
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	return rows, nil
}
