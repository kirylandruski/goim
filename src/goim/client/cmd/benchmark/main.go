package main

import (
	"flag"
	"goim/client"
	"log"
	"golang.org/x/net/context"
	"goim/rpc"
	"sync"
	"time"
	"sync/atomic"
	"goim/common"
)

var username = flag.String("username", "", "username")
var password = flag.String("password", "", "password")
var address = flag.String("address", "", "server address. example: localhost:5432")
var seconds = flag.Int("duration", 1, "duration of test in seconds")
var connections = flag.Int("connections", 1, "count of parallel connections")

const testInputsCount = 10000
const testItemLen = 100

func main() {
	testInputs := common.GenerateRandomStrings(testInputsCount, testItemLen)

	flag.Parse()
	durationTime := time.Duration(*seconds) * time.Second

	ctx, _ := context.WithTimeout(context.Background(), durationTime)
	log.Printf("starting %v parallel connections", *connections)
	log.Printf("running for %v seconds", *seconds)

	wg := sync.WaitGroup{}
	wg.Add(*connections)

	total := int64(0)

	for i := 0; i < *connections; i++ {
		go func() {
			c := client.NewClient()
			if err := c.Connect(context.Background(), *address, nil, *username, *password); err != nil {
				log.Fatal(err)
			}

			counter := int64(0)
			for {
				select {
				case <-ctx.Done():
					c.Disconnect()
					atomic.AddInt64(&total, counter)
					wg.Done()
					return
				default:
					var status int32
					var err error
					if counter%2 == 0 {
						_, err, status = c.SetStr(&testInputs[counter%testInputsCount], &testInputs[(counter+1)%testInputsCount], 50)
					} else {
						_, _, err, status = c.GetStr(&testInputs[counter%testInputsCount])
					}

					if err != nil {
						log.Fatal(err)
					}
					if status != rpc.SuccessStatus {
						log.Fatal("bad response status")
					}
					counter++
				}
			}

		}()
	}

	wg.Wait()

	rate := total / int64(*seconds)
	log.Printf("done %v requests per second", rate)
}
