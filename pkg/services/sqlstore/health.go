package sqlstore

import (
	"fmt"
	"github.com/elitecodegroovy/gnetwork/pkg/bus"
	m "github.com/elitecodegroovy/gnetwork/pkg/models"
)

func init() {
	bus.AddHandler("sql", GetDBHealthQuery)
	fmt.Println("Initialized sqlstore DB health....")
}

func GetDBHealthQuery(query *m.GetDBHealthQuery) error {
	return x.Ping()
}
