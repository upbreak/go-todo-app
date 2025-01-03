package main

import (
	"context"
	"fmt"
	"github.com/upbreak/go-todo-app/config"
	"golang.org/x/sync/errgroup"
	"log"
	"net"
	"net/http"
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

	server := &http.Server{
		//Addr:    ":18080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello World! %s", r.URL.Path[1:])
		}),
	}

	// 다른 고루틴에서 http서버를 실행
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		if err := server.Serve(l); err != nil && err != http.ErrServerClosed {
			log.Printf("failed to close %+v", err)
			return err
		}
		return nil
	})

	//채널로부터 알림(종료)을 기다린다
	<-ctx.Done()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("failed to shutdown server: %v", err)
	}

	// Go 메서드로 실행한 다른 고루팅의 종료를 기다림
	return eg.Wait()

}
