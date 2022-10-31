package maker

import (
	"context"
	"fmt"
	"log"
)

func (impl *makerImpl) GetOne(name string) (*ServiceInfo, error) {
	pool, err := impl.dbPool.Acquire(context.Background())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer pool.Release()

	var service *ServiceInfo
	selectStatement := `SELECT * FROM services WHERE name=$1`
	_ = pool.QueryRow(context.Background(), selectStatement, name).Scan(service)
	if err != nil || service == nil {
		log.Println(err)
		return nil, err
	}

	if impl.collectError != nil {
		return nil, fmt.Errorf("backend is not alive")
	}

	return service, nil
}
