package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/ruhollahh/go-ecom/internal/service/admin"
	"github.com/ruhollahh/go-ecom/internal/service/admin/stores/admindb"
	authsvc "github.com/ruhollahh/go-ecom/internal/service/auth"
	productsvc "github.com/ruhollahh/go-ecom/internal/service/product"
	"github.com/ruhollahh/go-ecom/internal/service/product/stores/productdb"
	usersvc "github.com/ruhollahh/go-ecom/internal/service/user"
	"github.com/ruhollahh/go-ecom/internal/service/user/stores/userdb"
	"github.com/ruhollahh/go-ecom/pkg/logger"
)

type Config struct {
	AdminAuthSvcCfg authsvc.Config
	UserAuthSvcCfg  authsvc.Config
}

type Service struct {
	AdminAuthSvc *authsvc.Service
	AdminSvc     *adminsvc.Service
	UserAuthSvc  *authsvc.Service
	UserSvc      *usersvc.Service
	ProductSvc   *productsvc.Service
}

func New(cfg Config, log *logger.Logger, db *sqlx.DB) Service {
	adminStorer := admindb.NewStore(log, db)
	adminAuthSvc := authsvc.NewService(cfg.AdminAuthSvcCfg)
	adminSvc := adminsvc.New(log, adminAuthSvc, adminStorer)
	userStorer := userdb.NewStore(log, db)
	userAuthSvc := authsvc.NewService(cfg.UserAuthSvcCfg)
	userSvc := usersvc.New(log, userAuthSvc, userStorer)
	productStorer := productdb.NewStore(log, db)
	productSvc := productsvc.New(productStorer)

	return Service{
		AdminAuthSvc: adminAuthSvc,
		AdminSvc:     adminSvc,
		UserAuthSvc:  userAuthSvc,
		UserSvc:      userSvc,
		ProductSvc:   productSvc,
	}
}
