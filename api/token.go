package api

import (
	"database/sql"
	"errors"
	"log"
	"math/big"
	"net/http"

	"github.com/aydogduyusuf/DBchain/access_refresh_tokens"
	"github.com/aydogduyusuf/DBchain/blockchain"
	db "github.com/aydogduyusuf/DBchain/db/sqlc"
	"github.com/aydogduyusuf/DBchain/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type deployTokenRequest struct {
	name			string
	symbol			string
	supply			int64
}

type deployTokenResponse struct {
	TokenAddress		string
}

func (server *Server) deployToken(ctx *gin.Context) {
	var req deployTokenRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*access_refresh_tokens.Payload)
	arg := db.CreateTokenParams{
		UID: 				authPayload.ID,
		TokenName: 			req.name,
		Symbol: 			req.symbol,
		Supply: 			req.supply,
	}

	_, err := server.store.CreateToken(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation" : 
				ctx.JSON(http.StatusForbidden, errorResponse(err))
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	user, err := server.store.GetUserFromID(ctx, authPayload.ID)
	privateKeyBytes, err := util.Decrypt([]byte(user.WalletPrivateAddress), secretKey)
	privateKey := string(privateKeyBytes)
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}
	contractAddress, _ := blockchain.DeployContract(common.BytesToAddress([]byte(user.WalletPublicAddress)), privateKeyECDSA, arg.TokenName, arg.Symbol, big.NewInt(arg.Supply))
	txArg := db.CreateTransactionParams{
		TransactionType: "deploy",
		FromAddress: user.WalletPublicAddress,
		ToAddress: "",
		TransferData: contractAddress.String(),
		HashValue: contractAddress.Hash().String(),
	}
	server.store.CreateTransaction(ctx, txArg)
	ctx.JSON(http.StatusOK, contractAddress)
}

type getTokenRequest struct {
	ID 	uuid.UUID 	`uri:"id" binding:"required,min=1"`
}

func (server *Server) getToken(ctx *gin.Context) {
	var req getTokenRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}

	token, err := server.store.GetToken(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*access_refresh_tokens.Payload)
	if token.UID != authPayload.ID {
		err := errors.New("token does not belong to authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, token)
}

type listTokensRequest struct {
	UserID   uuid.UUID `json:"user_id" binding:"required,min=1"`
}

func (server *Server) listTokens(ctx *gin.Context) {
	var req listTokensRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*access_refresh_tokens.Payload)
	
	tokens, err := server.store.ListTokens(ctx, authPayload.ID)
	if err != nil {
		
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, tokens)
}