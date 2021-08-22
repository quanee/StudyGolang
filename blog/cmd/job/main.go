package main

import (
	"context"
	"errors"
	"flag"
	"os"
	"os/signal"
	"time"

	"blog/internal/conf"
	"blog/internal/service"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/sync/errgroup"
	"gopkg.in/yaml.v2"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(service *service.JobService) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		return service.UpdateCache(ctx)
	}
}

func main() {
	flag.Parse()
	logger := log.NewStdLogger(os.Stdout)

	cfg := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
		config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
			return yaml.Unmarshal(kv.Value, v)
		}),
	)
	if err := cfg.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := cfg.Scan(&bc); err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer func(ctx context.Context) {
		// Do not make the application hang when it is shutdown.
		ctx, cancel = context.WithTimeout(ctx, time.Second*5)
		defer cancel()
	}(ctx)

	run, cleanup, err := initApp(bc.Data, bc.Cache, bc.Mq, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return run(ctx)
	})

	g.Go(func() error {
		return sigProcess()(ctx)
	})

	if err := g.Wait(); err != nil {
		_ = logger.Log(log.LevelError, err)
	}
}

func sigProcess(sig ...os.Signal) func(context.Context) error {
	return func(ctx context.Context) error {

		if len(sig) == 0 {
			sig = append(sig, os.Interrupt)
		}

		done := make(chan os.Signal, len(sig))
		signal.Notify(done, sig...)

		var err error
		select {
		case <-ctx.Done():
		case s := <-done:
			err = errors.New("main: " + s.String())
		}

		signal.Stop(done)
		close(done)

		return err
	}
}
