package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"context"
	"goim/server"
	"runtime/pprof"
	"goim/mem_store"
	"goim/network"
	"goim/auth"
)

var usersPath = flag.String("users", "", "path to file containing users. example: ./users.json")
var certPath = flag.String("cert", "", "path to server.crt file. example: ./server.crt")
var keyPath = flag.String("key", "", "path to server.key file. example: ./server.key")
var address = flag.String("address", "", "server address. example: localhost:5432")
var cpuProfile = flag.String("cpu_profile", "", "path to cpu profile")

func handleInterrupt(cancel func()) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt)
	go func() {
		for range interruptChan {
			cancel()
		}
	}()
}

func initUserManager(users string) *auth.UserManager {
	m := auth.NewUserManager()

	file, err := os.OpenFile(users, os.O_RDONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Read(file); err != nil {
		log.Fatal(err)
	}

	return m
}

func main() {
	flag.Parse()

	if len(*cpuProfile) > 0 {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	ctx, cancel := context.WithCancel(context.Background())
	handleInterrupt(cancel)

	handler := server.AppHandler{}
	handler.Authenticator = initUserManager(*usersPath)
	handler.Storer = mem_store.NewSyncStore(context.Background(), true)

	s, err := network.NewServer(&handler, *certPath, *keyPath)
	if err != nil {
		log.Fatalf("failed to initiate server: %v\n", err.Error())
	}

	if err := s.Start(ctx, *address); err != nil {
		log.Fatal(err)
	}
}
