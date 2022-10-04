package subscriber

import (
	"L0/cache"
	"L0/database"
	"L0/models"
	"L0/repository"
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/stan.go"
)

func Subscribe() {
	or := models.Order{}
	database.Connect()

	sc, err := stan.Connect("test-cluster", "sub")
	if err != nil {
		panic(err)
	}

	sc.Subscribe("testchan2", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
		if err := json.Unmarshal(m.Data, &or); err != nil {
			log.Printf("[err] Nats-pub JSON: %s\n", err.Error())
		}

		cache.Set(or)
		repository.Insert(or)
	}, stan.DeliverAllAvailable())
}
