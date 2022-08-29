package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/aydogduyusuf/DBchain/access_refresh_tokens"
	db "github.com/aydogduyusuf/DBchain/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type getTransactionRequest struct {
	txID	uuid.UUID		`json:"tx_id"`
	fromAddress string		`json:"from_address"`
}

func (server *Server) createTransaction(ctx *gin.Context, txParams db.CreateTransactionParams) {

	result, err := server.store.CreateTransaction(ctx, txParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}


func (server *Server) getTransaction(ctx *gin.Context) {
	var req getTransactionRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}

	tx, err := server.store.GetTransaction(ctx, req.txID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*access_refresh_tokens.Payload)
	if tx.UID != authPayload.ID {
		err := errors.New("token does not belong to authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, tx)
}