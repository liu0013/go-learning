package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

const shutdownTimeout = 10 * time.Second
const mockRequestTime = 5 * time.Second

type mockHttpHandler struct{}

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	server1 := &http.Server{Addr: "127.0.0.1:19000", Handler: mockHttpHandler{}}
	server2 := &http.Server{Addr: "127.0.0.1:19001", Handler: mockHttpHandler{}}
	start(server1, g, ctx)
	start(server2, g, ctx)
	handleSignal(g, ctx)
	if err := g.Wait(); err != nil {
		fmt.Println(err.Error())
	}
}

func (_ mockHttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(mockRequestTime)
	fmt.Fprintf(w, "Hello Golang\n")
}

func start(server *http.Server, g *errgroup.Group, ctx context.Context) {
	g.Go(func() error { return shutdown(server, ctx) })
	g.Go(func() error {
		fmt.Println("Server Start: ", server.Addr)
		return server.ListenAndServe()
	})
}

func shutdown(server *http.Server, ctx context.Context) error {
	<-ctx.Done()
	fmt.Println("Server Shutdown: ", server.Addr)
	shutdownCtx, _ := context.WithTimeout(context.Background(), shutdownTimeout)
	return server.Shutdown(shutdownCtx)
}

func handleSignal(g *errgroup.Group, ctx context.Context) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	g.Go(func() error {
		select {
		case s := <-c:
			fmt.Println("Signal: ", s)
			return errors.Wrapf(errors.New("Exit Signal"), "", s)
		case <-ctx.Done():
			return nil
		}
	})
}
