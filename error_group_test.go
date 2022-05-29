package go_advanced

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

var srv = &http.Server{Addr: ":8080"}

func startHttpServer() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world\n")
	})
	return srv.ListenAndServe()
}

func notifySignal() error {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for s := range c {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			return errors.New("Exiting Program")
		default:
			fmt.Println("other signal", s)
		}
	}
	return nil
}

func TestErrorGroup(t *testing.T) {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(startHttpServer)
	g.Go(notifySignal)

	err := g.Wait()
	println(err)
	println(ctx.Err())
	if err != nil {
		srv.Shutdown(ctx)
	}
}
