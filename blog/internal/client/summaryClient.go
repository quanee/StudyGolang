package client

import (
	"context"

	pb "blog/api/summary/v1"
	"blog/internal/conf"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
)

func NewSummaryClient(conf *conf.Service) pb.SummaryServiceHTTPClient {
	conn, err := transhttp.NewClient(
		context.Background(),
		transhttp.WithMiddleware(
			recovery.Recovery(),
		),
		transhttp.WithEndpoint(conf.ServiceMap["Summary"]),
	)
	if err != nil {
		panic(err)
	}

	return pb.NewSummaryServiceHTTPClient(conn)
}
