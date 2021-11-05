package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	svr := http.Server{Addr: ":8088"}
	// http server
	g.Go(func() error {
		go func() {
			<-ctx.Done()
			fmt.Println("http done")
			_ = svr.Shutdown(ctx)
		}()
		return svr.ListenAndServe()
	})

	// signal
	g.Go(func() error {
		exitSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}
		sig := make(chan os.Signal, len(exitSignals))
		signal.Notify(sig, exitSignals...)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("signal done")
				return ctx.Err()
			case <-sig:
				return nil
			}
		}
	})

	// make error
	g.Go(func() error {
		fmt.Println("make error start")
		time.Sleep(time.Second)
		fmt.Println("make finish")
		return errors.New("make error wrapper")
	})

	err := g.Wait()
	fmt.Println(err, "\r\n first error return , all exit")
}
