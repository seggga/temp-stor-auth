package application

import (
	"context"

	"github.com/seggga/temp-stor-auth/internal/adapters/grpc"
	"github.com/seggga/temp-stor-auth/internal/adapters/rest"
	"github.com/seggga/temp-stor-auth/internal/adapters/storage/memory"
	"github.com/seggga/temp-stor-auth/internal/domain/auth"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

var (
	logger   *zap.Logger
	rService *rest.Service
	gService *grpc.Service
)

// Start ...
func Start(ctx context.Context) {

	cfg := readConfig()
	logger = initLogger(cfg.Logger.Level)
	stor := memory.New()

	authService := auth.New(stor)
	gService = grpc.New(stor)
	rService = rest.New(authService, logger, cfg.RestPort)

	var g errgroup.Group
	g.Go(func() error {
		return gService.Start(ctx)
	})
	g.Go(func() error {
		return rService.Start(ctx)
	})

	logger.Info("app is started")
	err := g.Wait()
	if err != nil {
		logger.Sugar().Fatalf("applicateion failed: %v", err)
	}
}

// Stop ...
func Stop(ctx context.Context) {
	defer logger.Sync()
	_ = rService.Stop(ctx)
	// _ = gService.Stop(ctx)

	logger.Info("application has been stopped")

}
