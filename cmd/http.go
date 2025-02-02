package cmd

import (
	"ewallet-wallet/helpers"
	"ewallet-wallet/internal/api"
	"ewallet-wallet/internal/interfaces"
	"ewallet-wallet/internal/repository"
	"ewallet-wallet/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServerHTTP() {
	d := dependencyInject()

	healtcheckSvc := &services.Healthcheck{
	}
	healtcheckAPI := &api.Healthcheck{
		HealthcheckServices: healtcheckSvc,
	}

	r := gin.New()
	r.GET("/health", healtcheckAPI.HealthCheckHandlerHTTP)

	walletV1 := r.Group("/wallet/v1")
	walletV1.POST("/", d.WalletAPI.Create)

	err := r.Run(":" + helpers.GetEnv("PORT","8082"))
	if err != nil {
		log.Fatal(err.Error())
	}
}

type Dependency struct {
	HealthcheckAPI interfaces.IHealthcheckAPI
	WalletAPI      interfaces.IWalletAPI
// 	External       interfaces.IExternal
}

func dependencyInject() Dependency {
	healthcheckSvc := &services.Healthcheck{}
	healthcheckAPI := &api.Healthcheck{
		HealthcheckServices: healthcheckSvc,
	}

	walletRepo := &repository.WalletRepo{
		DB: helpers.DB,
	}

	walletSvc := &services.WalletService{
		WalletRepo: walletRepo,
	}
	walletAPI := &api.WalletAPI{
		WalletService: walletSvc,
	}

	// external := &external.External{}

	return Dependency{
		HealthcheckAPI: healthcheckAPI,
		WalletAPI:      walletAPI,
		// External:       external,
	}
}