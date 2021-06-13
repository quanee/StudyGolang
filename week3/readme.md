```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("hello world!"))
}

// 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
func main() {
	// 用于取消全部的 Context
	ctx, cancel := context.WithCancel(context.Background())
	// 创建 error Group 和 error Context
	eg, errCtx := errgroup.WithContext(ctx)

	server := &http.Server{Addr: ":8080"}
	// 启动服务
	eg.Go(func() error {
		http.HandleFunc("/hello", Hello)
		fmt.Println("http server start")
		return server.ListenAndServe()
	})

	// 处理错误
	eg.Go(func() error {
		<-errCtx.Done()
		fmt.Println("http server stop")
		return server.Shutdown(errCtx) // 关闭 http server
	})

	// 处理信号
	eg.Go(func() error {
		// chan buffer 必须大于 0
		sChan := make(chan os.Signal, 1)
		// 信号注册
		signal.Notify(sChan)
		for {
			select {
			case <-errCtx.Done(): // 错误处理
				fmt.Println("context done")
				return errCtx.Err()
			case <-sChan: // 信号处理
				// 调用取消函数
				cancel()
				fmt.Println("signal cancel")
				return nil
			}
		}
	})

	if err := eg.Wait(); err != nil {
		fmt.Println("error group error: ", err)
	}
	fmt.Println("all error group done!")
}
```
