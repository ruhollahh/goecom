package main

import (
	"context"
	"errors"
	"expvar"
	"fmt"
	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/ardanlabs/conf/v3"
	_ "github.com/go-sql-driver/mysql"
	httpserver "github.com/ruhollahh/go-ecom/delivery/http"
	_ "github.com/ruhollahh/go-ecom/docs"
	"github.com/ruhollahh/go-ecom/internal/clients/dbpostgre"
	"github.com/ruhollahh/go-ecom/internal/clients/redis"
	"github.com/ruhollahh/go-ecom/internal/service"
	"github.com/ruhollahh/go-ecom/internal/service/auth"
	"github.com/ruhollahh/go-ecom/pkg/logger"
	"github.com/ruhollahh/go-ecom/pkg/validate"
	"github.com/ruhollahh/go-ecom/pkg/web/debug"
	"github.com/ruhollahh/go-ecom/pkg/web/meta"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

var build = "develop"

// @title GoEcom API
// @version 1.0
// @description an e-commerce web API.
// @termsOfService https://goecom.com/terms/

// @contact.name API Support
// @contact.url https://www.goecom.com/support
// @contact.email ruhollahh01@gmail.com

// @securityDefinitions.apikey AuthBearerAdmin
// @in header
// @name Authorization
// @description Type the word 'Bearer' followed by a space and admin access token

// @securityDefinitions.apikey AuthBearerCustomer
// @in header
// @name Authorization
// @description Type the word 'Bearer' followed by a space and user access token

// @BasePath /
func main() {
	var log *logger.Logger

	traceIDFn := func(ctx context.Context) string {
		return httpmeta.GetTraceID(ctx)
	}

	log = logger.New(os.Stdout, logger.LevelInfo, "GOECOM", traceIDFn)

	// -------------------------------------------------------------------------

	ctx := context.Background()

	if err := run(ctx, log); err != nil {
		log.Error(ctx, "startup", "msg", err)
		return
	}
}

func run(ctx context.Context, log *logger.Logger) error {

	// -------------------------------------------------------------------------
	// GOMAXPROCS

	log.Info(ctx, "startup", "GOMAXPROCS", runtime.GOMAXPROCS(0), "build", build)

	// -------------------------------------------------------------------------
	// Configuration

	cfg := struct {
		conf.Version
		Web struct {
			ShutdownTimeout time.Duration `conf:"default:20s"`
			APIPort         string        `conf:"default:3000"`
			APIHost         string        `conf:"default:0.0.0.0"`
			DebugPort       string        `conf:"default:4000"`
			DebugHost       string        `conf:"default:0.0.0.0"`
		}
		UserAuth struct {
			SignKey               string        `conf:"default:usersecret123,mask"`
			AccessExpirationTime  time.Duration `conf:"default:24h"`
			RefreshExpirationTime time.Duration `conf:"default:72h"`
			AccessSubject         string        `conf:"default:access"`
			RefreshSubject        string        `conf:"default:refresh"`
		}
		AdminAuth struct {
			SignKey               string        `conf:"default:adminsecret123,mask"`
			AccessExpirationTime  time.Duration `conf:"default:24h"`
			RefreshExpirationTime time.Duration `conf:"default:72h"`
			AccessSubject         string        `conf:"default:access"`
			RefreshSubject        string        `conf:"default:refresh"`
		}
		Redis struct {
			Host     string
			Port     int
			Password string `conf:"mask"`
			DB       int
		}
		Postgres struct {
			User         string
			Password     string `conf:"mask"`
			Host         string
			Port         string
			Name         string
			MaxIdleConns int           `conf:"default:25"`
			MaxOpenConns int           `conf:"default:25"`
			MaxIdleTime  time.Duration `conf:"default:15m"`
		}
		Minio struct {
			Host   string
			Port   string
			Access string
			Secret string `conf:"mask"`
			Bucket string `conf:"default:goecom"`
		}
	}{
		Version: conf.Version{
			Build: build,
			Desc:  "copyright information here",
		},
	}

	const prefix = "GOECOM"
	help, err := conf.Parse(prefix, &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil
		}
		return fmt.Errorf("parsing config: %w", err)
	}

	// -------------------------------------------------------------------------
	// App Starting

	log.Info(ctx, "starting service", "version", build)
	defer log.Info(ctx, "shutdown complete")

	out, err := conf.String(&cfg)
	if err != nil {
		return fmt.Errorf("generating config for output: %w", err)
	}
	log.Info(ctx, "startup", "config", out)

	expvar.NewString("build").Set(build)

	err = validate.Init()
	if err != nil {
		return err
	}

	// -------------------------------------------------------------------------
	// Database Support

	log.Info(ctx, "startup", "status", "initializing database support", "host", cfg.Postgres.Host)

	db, err := dbpostgre.Open(dbpostgre.Config{
		User:         cfg.Postgres.User,
		Password:     cfg.Postgres.Password,
		Host:         cfg.Postgres.Host,
		Port:         cfg.Postgres.Port,
		Name:         cfg.Postgres.Name,
		MaxIdleConns: cfg.Postgres.MaxIdleConns,
		MaxOpenConns: cfg.Postgres.MaxOpenConns,
		MaxIdleTime:  cfg.Postgres.MaxIdleTime,
	})
	if err != nil {
		return fmt.Errorf("connecting to db: %w", err)
	}
	defer func() {
		log.Info(ctx, "shutdown", "status", "stopping database support", "host", cfg.Postgres.Host)
		db.Close()
	}()

	// -------------------------------------------------------------------------
	// Storage Support
	//minioClient, err := minio.New(fmt.Sprintf("%s:%s", cfg.Minio.Host, cfg.Minio.Port), &minio.Options{
	//	Creds:  credentials.NewStaticV4(cfg.Minio.Access, cfg.Minio.Secret, ""),
	//	Secure: false,
	//})
	//if err != nil {
	//	return fmt.Errorf("creating minio client: %w", err)
	//}
	//
	//err = minioClient.MakeBucket(ctx, cfg.Minio.Bucket, minio.MakeBucketOptions{})
	//if err != nil {
	//	return fmt.Errorf("making bucket: %w", err)
	//}
	//
	//policy := fmt.Sprintf(`{
	//	"Version": "2012-10-17",
	//	"Statement": [
	//		{
	//			"Effect": "Allow",
	//			"Principal": "*",
	//			"Action": "s3:GetObject",
	//			"Resource": "arn:aws:s3:::%s/*"
	//		}
	//	]
	//}`, cfg.Minio.Bucket)
	//
	//err = minioClient.SetBucketPolicy(ctx, cfg.Minio.Bucket, policy)
	//if err != nil {
	//	return fmt.Errorf("setting bucket policy: %w", err)
	//}

	// -------------------------------------------------------------------------
	// Initialize Services

	log.Info(ctx, "startup", "status", "initializing authentication support")

	redisClient := redisclient.New(redisclient.Config{
		Host:     cfg.Redis.Host,
		Port:     cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	defer func() {
		redisClient.Close()
	}()

	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Store = redisstore.New(redisClient)

	services := service.New(service.Config{
		UserAuthSvcCfg: authsvc.Config{
			SignKey:               []byte(cfg.UserAuth.SignKey),
			AccessExpirationTime:  cfg.UserAuth.AccessExpirationTime,
			RefreshExpirationTime: cfg.UserAuth.RefreshExpirationTime,
			AccessSubject:         cfg.UserAuth.AccessSubject,
			RefreshSubject:        cfg.UserAuth.RefreshSubject,
		},
		AdminAuthSvcCfg: authsvc.Config{
			SignKey:               []byte(cfg.AdminAuth.SignKey),
			AccessExpirationTime:  cfg.AdminAuth.AccessExpirationTime,
			RefreshExpirationTime: cfg.AdminAuth.RefreshExpirationTime,
			AccessSubject:         cfg.AdminAuth.AccessSubject,
			RefreshSubject:        cfg.AdminAuth.RefreshSubject,
		},
	}, log, db)

	// -------------------------------------------------------------------------
	// Start Debug Service

	go func() {
		debugAddr := fmt.Sprintf("%s:%s", cfg.Web.DebugHost, cfg.Web.DebugPort)
		log.Info(ctx, "startup", "status", "debug v1 router started", "host", debugAddr)

		if err := http.ListenAndServe(debugAddr, debug.Mux()); err != nil {
			log.Error(ctx, "shutdown", "status", "debug v1 router closed", "host", debugAddr, "msg", err)
		}
	}()

	// -------------------------------------------------------------------------
	// Start API Service

	log.Info(ctx, "startup", "status", "initializing V1 API support")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	httpServer := httpserver.New(httpserver.Config{
		Build:   build,
		APIPort: cfg.Web.APIPort,
		APIHost: cfg.Web.APIHost,
	}, log, db, services, sessionManager)

	serverErrors := make(chan error, 1)

	go func() {
		log.Info(ctx, "startup", "status", "api router started", "host", fmt.Sprintf("%s:%s", cfg.Web.APIHost, cfg.Web.APIPort))
		serverErrors <- httpServer.Serve()
	}()

	// -------------------------------------------------------------------------
	// Shutdown

	select {
	case err = <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		log.Info(ctx, "shutdown", "status", "shutdown started", "signal", sig)
		defer log.Info(ctx, "shutdown", "status", "shutdown complete", "signal", sig)

		ctx, cancel := context.WithTimeout(ctx, cfg.Web.ShutdownTimeout)
		defer cancel()

		if err = httpServer.Shutdown(ctx); err != nil {
			httpServer.Close()
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil
}
