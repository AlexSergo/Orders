package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/nats-io/stan.go"
)

func main() {
	// e := godotenv.Load()
	// if e != nil {
	// 	fmt.Print(e)
	// }

	sc, err := stan.Connect("test-cluster", "test-prod", stan.NatsURL(stan.DefaultNatsURL),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Printf("Connection lost, reason: %v", reason)
		}))
	if err != nil {
		log.Printf("[err] Nats-pub: %s\n", err.Error())
		return
	}
	defer sc.Close()

	for i := 1; i <= 3; i++ {
		file, err := os.Open("C:/Users/Alexandr/Desktop/projects/GO/L0/test_models/model" + strconv.Itoa(i) + ".json")
		if err != nil {
			log.Printf("[err] Nats-pub open file: %s\n", err.Error())
			return
		}
		data, err := ioutil.ReadAll(file)
		if err != nil {
			log.Printf("[err] Nats-pub ioutil: %s\n", err.Error())
			return
		}

		err = sc.Publish("testchan2", data)

		if err != nil {
			log.Printf("[err] Nats-pub: %s\n", err.Error())
		}
		fmt.Println("Publish was successful!")
		time.Sleep(1 * time.Second)
	}
}
