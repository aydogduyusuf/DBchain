package api

import (
	"database/sql"
	"log"
	"time"

	//"errors"
	//"fmt"
	"net/http"

	"github.com/aydogduyusuf/DBchain/access_refresh_tokens"
	db "github.com/aydogduyusuf/DBchain/db/sqlc"
	"github.com/gin-gonic/gin"
	//"github.com/golang/protobuf/ptypes/timestamp"
	//"github.com/google/uuid"
)

/* func (server *Server) createTransaction(ctx *gin.Context, txParams db.CreateTransactionParams) {

	result, err := server.store.CreateTransaction(ctx, txParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
} */

func (server *Server) getDeploys(ctx *gin.Context) {

	authPayload := ctx.MustGet(authorizationPayloadKey).(*access_refresh_tokens.Payload)
	user, err := server.store.GetUser(ctx, authPayload.Username)
	if err != nil {
		log.Fatal(err)
	}

	if ctx.Query("token") != "" {
		token, err := server.store.GetTokenByAddress(ctx, ctx.Query("token"))
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		tx, err := server.store.GetTransactionByAddress(ctx, token.ContractAddress)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, tx)
		return
	} else if ctx.Query("from") != "" && ctx.Query("to") != "" {
		startTime, err := time.Parse("2006-01-02", ctx.Query("from"))
		if err != nil {
			log.Fatal(err)
		}
		endTime, err := time.Parse("2006-01-02", ctx.Query("to"))
		if err != nil {
			log.Fatal(err)
		}
		arg := db.ListDeploysByTimeParams{
			CreateTime:      startTime,
			CreateTime_2:    endTime,
			TransactionType: "deploy",
			FromAddress:     user.WalletPublicAddress,
		}
		txs, err := server.store.ListDeploysByTime(ctx, arg)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, txs)
		return
	} else {
		arg := db.ListDeploysByUserParams{
			FromAddress:     user.WalletPublicAddress,
			TransactionType: "deploy",
		}
		txs, err := server.store.ListDeploysByUser(ctx, arg)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, txs)
		return
	}
}

func (server *Server) getTransfers(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*access_refresh_tokens.Payload)
	user, _ := server.store.GetUser(ctx, authPayload.Username)

	if ctx.Query("type") == "from" {
		arg := db.ListTransactionsByTypeFromParams{
			FromAddress:     user.WalletPublicAddress,
			TransactionType: "transfer",
		}

		txs, err := server.store.ListTransactionsByTypeFrom(ctx, arg)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, txs)
		return

	} else if ctx.Query("type") == "to" {
		arg := db.ListTransactionsByTypeToParams{
			ToAddress:       user.WalletPublicAddress,
			TransactionType: "transfer",
		}

		txs, err := server.store.ListTransactionsByTypeTo(ctx, arg)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, txs)
		return
	} else if ctx.Query("from") != "" && ctx.Query("to") != "" {
		startTime, err := time.Parse("2006-01-02", ctx.Query("from"))
		if err != nil {
			log.Fatal(err)
		}
		endTime, err := time.Parse("2006-01-02", ctx.Query("to"))
		if err != nil {
			log.Fatal(err)
		}
		arg := db.ListTransfersByTimeFromParams{
			CreateTime:      startTime,
			CreateTime_2:    endTime,
			TransactionType: "transfer",
			FromAddress:     user.WalletPublicAddress,
		}
		txs, err := server.store.ListTransfersByTimeFrom(ctx, arg)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, txs)

		arg2 := db.ListTransfersByTimeToParams{
			CreateTime:      startTime,
			CreateTime_2:    endTime,
			TransactionType: "transfer",
			ToAddress:       user.WalletPublicAddress,
		}
		txs, err = server.store.ListTransfersByTimeTo(ctx, arg2)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, txs)
		return
	} else {
		arg := db.ListTransactionsByTypeFromParams{
			FromAddress:     user.WalletPublicAddress,
			TransactionType: "transfer",
		}

		txs, err := server.store.ListTransactionsByTypeFrom(ctx, arg)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, txs)

		arg2 := db.ListTransactionsByTypeToParams{
			ToAddress:       user.WalletPublicAddress,
			TransactionType: "transfer",
		}

		txs, err = server.store.ListTransactionsByTypeTo(ctx, arg2)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, txs)
	}

}
