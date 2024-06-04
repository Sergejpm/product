package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/sergejpm/product/internal/domain/service/authorization"
	"github.com/sergejpm/product/internal/middleware"
	"net/http"
	_ "net/http/pprof"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/sync/errgroup"
	_ "google.golang.org/grpc/encoding/proto"

	"github.com/sergejpm/product/internal/infra/log"
	"github.com/sergejpm/product/pkg/api/product"
)

func runHTTPServer(ctx context.Context, port uint, server product.ProductServer, authService *authorization.Service) error {
	g, ctx := errgroup.WithContext(ctx)

	address := fmt.Sprintf(":%d", port)
	mux := runtime.NewServeMux()
	authMiddleware := middleware.AuthHandler(mux, authService)

	httpSvc := &http.Server{
		Addr:    address,
		Handler: authMiddleware,
	}

	err := product.RegisterProductHandlerServer(ctx, mux, server)

	if err != nil {
		log.Logger().Errorf("unable to register sber spasibo server as http handler: %v", err)
		return err
	}

	g.Go(func() error {
		errLS := httpSvc.ListenAndServe()
		if errLS != nil && !errors.Is(errLS, http.ErrServerClosed) {
			log.Logger().Errorf("unable to start http server: %v", errLS)
			return errLS
		}
		return nil
	})

	g.Go(func() error {
		<-ctx.Done()
		e := httpSvc.Shutdown(context.Background())
		if e == nil || !errors.Is(e, http.ErrServerClosed) {
			log.Logger().Infof("http server stopped :%s", e)
			return nil
		}
		return e
	})

	log.Logger().Infof("http server started at address: %s", address)

	if err = g.Wait(); err != nil {
		log.Logger().Errorf("server errgroup error :%s", err)
	}

	return nil
}
