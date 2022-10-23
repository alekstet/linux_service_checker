package collector

import (
	"context"
	"log"
)

type Collector interface {
	Collect() (*ServicesInfo, error)
}

func (pool *PGXPool) Collect() (*ServicesInfo, error) {
	servicesInfo := make(ServicesInfo)

	newPool, err := pool.DBPool.Acquire(context.Background())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer newPool.Release()

	rows, err := newPool.Query(context.Background(), "SELECT * FROM services")
	if err != nil {
		log.Println(err)
		return nil, err
	}

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

	return &servicesInfo, nil
}
