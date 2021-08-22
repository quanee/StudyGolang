package client

import (
	"context"

	pb "blog/api/blog/v1"
	"blog/internal/conf"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
)

func NewBlogClient(conf *conf.Service) pb.BlogHTTPClient {
	conn, err := transhttp.NewClient(
		context.Background(),
		transhttp.WithMiddleware(
			recovery.Recovery(),
		),
		transhttp.WithEndpoint(conf.ServiceMap["Blog"]),
	)
	if err != nil {
		panic(err)
	}

	return pb.NewBlogHTTPClient(conn)
}
