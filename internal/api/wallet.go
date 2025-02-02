package api

import (
	"ewallet-wallet/helpers"
	"ewallet-wallet/internal/interfaces"
	"ewallet-wallet/internal/models"
	"net/http"

	"ewallet-wallet/constants"

	"github.com/gin-gonic/gin"
)

type WalletAPI struct { 
	WalletService interfaces.IWalletService
}

func (api *WalletAPI) Create(c *gin.Context) {
	var (
		log = helpers.Loggger
		req  models.Wallet
		
	)
	if err := c.ShouldBindJSON(&req); err != nil { 
		log.Error("failed to parse request: ", err)
		helpers.SendResponseHttp(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}
	if req.UserID == 0 {
		log.Error("user id is empty")
		helpers.SendResponseHttp(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}
	err := api.WalletService.Create(c.Request.Context(), &req)
	if err != nil {
		log.Error("failed to create wallet: ", err)
		helpers.SendResponseHttp(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}
	helpers.SendResponseHttp(c, http.StatusOK, constants.SuccessMessage, req)
}