package main

import (
	"container/ring"
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

type MovingTimeWindow struct {
	bucket    *ring.Ring
	window    time.Duration
	precision time.Duration
	ticker    *time.Ticker
	mu        sync.Mutex
}

func GenerateMovingTimeWindow(w time.Duration, p time.Duration) *MovingTimeWindow {
	return &MovingTimeWindow{
		bucket:    ring.New(int(w.Seconds() / p.Seconds())),
		window:    w,
		precision: p,
		ticker:    time.NewTicker(p),
	}
}

func (m *MovingTimeWindow) init() {
	r := m.bucket
	for i := 0; i < m.bucket.Len(); i++ {
		r.Value = 0
		r = r.Next()
	}
}

func (m *MovingTimeWindow) GetRollingSum() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	sum := 0
	r := m.bucket
	for i := 0; i < m.bucket.Len(); i++ {
		sum += r.Value.(int)
		r = r.Next()
	}
	return sum
}

func (m *MovingTimeWindow) Add(inc int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.bucket.Value = m.bucket.Value.(int) + inc
	m.debug()
}

func (m *MovingTimeWindow) debug() {
	result := make([]int, m.bucket.Len())
	r := m.bucket
	for i := 0; i < m.bucket.Len(); i++ {
		result[i] = r.Value.(int)
		r = r.Next()
	}
	fmt.Println("debug to print the whole window buckets: ", result)
}

func (m *MovingTimeWindow) Start() {
	// init bucket
	m.init()
	go func() {
		for range m.ticker.C {
			m.next()
		}
	}()
}

func (m *MovingTimeWindow) Stop() {
	m.ticker.Stop()
}

// move to next timeframe
func (m *MovingTimeWindow) next() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.bucket = m.bucket.Next()
	m.bucket.Value = 0
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	rollingNumber := GenerateMovingTimeWindow(60*time.Second, 5*time.Second)
	start(rollingNumber, g, ctx)
	for i := 0; i < 120; i++ {
		time.Sleep(1 * time.Second)
		rollingNumber.Add(1)
		fmt.Println("the time window sum: ", rollingNumber.GetRollingSum())
	}
	handleSignal(g, ctx)
	if err := g.Wait(); err != nil {
		fmt.Println(err.Error())
	}
}

func start(timeWindow *MovingTimeWindow, g *errgroup.Group, ctx context.Context) {
	g.Go(func() error { return shutdownHook(timeWindow, ctx) })
	g.Go(func() error {
		timeWindow.Start()
		fmt.Println("time window start!")
		return nil
	})
}

func shutdownHook(timeWindow *MovingTimeWindow, ctx context.Context) error {
	<-ctx.Done()
	timeWindow.Stop()
	fmt.Println("time window shutdown!")
	return nil
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
