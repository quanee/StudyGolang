package server

import (
	"context"
	"errors"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/quanee/go-training/week4/api/article/v1"
	"github.com/quanee/go-training/week4/internal/service"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type Server struct {
	service *service.Service
}

func NewServer(s *service.Service) *Server {
	return &Server{service: s}
}

func (srv *Server) Run() error {
	lis, err := net.Listen("tcp", viper.GetString("grpc.port"))
	if err != nil {
		return err
	}
	egrp, ctx := errgroup.WithContext(context.Background())
	rpcs := grpc.NewServer()
	egrp.Go(func() error {
		go func() {
			<-ctx.Done()
			rpcs.GracefulStop()
			log.Printf("Shutdown Server")
		}()
		pb.RegisterArticleServer(rpcs, srv.service)
		return rpcs.Serve(lis)
	})
	egrp.Go(func() error {
		signals := []os.Signal{syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT}
		sigChan := make(chan os.Signal, len(signals))
		signal.Notify(sigChan, signals...)
		for {
			select {
			case <-ctx.Done():
				return nil
			case sig := <-sigChan:
				log.Printf("get a signal %s", sig.String())
				switch sig {
				case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
					return errors.New("Close by signal " + sig.String())
				case syscall.SIGHUP:
				default:
					return errors.New("Undefined signal")
				}
			}
		}
	})
	return egrp.Wait()
}
