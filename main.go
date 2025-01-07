package main

import (
	"context"
	"fmt"
	"github.com/upbreak/go-todo-app/config"
	"log"
	"net"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
	}
}

func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
	}

	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)

	dbCfg, err := config.DBNew()
	if err != nil {
		return err
	}

	mux, cleanup, err := NewMux(ctx, dbCfg)
	defer cleanup()
	if err != nil {
		return err
	}

	server := NewServer(l, mux)

	return server.Run(ctx)

}
