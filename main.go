package main

import (
	"fmt"
	"net/http"
	//"os"
	"context"
	"log"
	"golang.org/x/sync/errgroup"
)

func main() {
	if err:= run(context.Background()) ; err != nil {
		log.Printf("failed to terminate server: %v", err)
	}
}

func run(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":18080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			fmt.Fprintf(w, "Hello World! %s", r.URL.Path[1:])
		}),
	}

	// 다른 고루틴에서 http서버를 실행
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("failed to close %v", err);
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