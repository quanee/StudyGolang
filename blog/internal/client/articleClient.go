package client

import (
	"context"

	pb "blog/api/article/v1"
	"blog/internal/conf"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
)

func NewArticleClient(conf *conf.Service) pb.ArticleServiceHTTPClient {
	conn, err := transhttp.NewClient(
		context.Background(),
		transhttp.WithMiddleware(
			recovery.Recovery(),
		),
		transhttp.WithEndpoint(conf.ServiceMap["Article"]),
	)
	if err != nil {
		panic(err)
	}

	return pb.NewArticleServiceHTTPClient(conn)
}
