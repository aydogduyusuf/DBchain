package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"

	//"strconv"

	"github.com/aydogduyusuf/DBchain/access_refresh_tokens"
	"github.com/aydogduyusuf/DBchain/blockchain"
	db "github.com/aydogduyusuf/DBchain/db/sqlc"
	"github.com/aydogduyusuf/DBchain/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"

	//"github.com/google/uuid"
	"github.com/lib/pq"
)

type deployTokenRequest struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Supply int64  `json:"supply"`
}

type deployTokenResponse struct {
	TokenAddress string `json:"token_address"`
}

func (server *Server) deployToken(ctx *gin.Context) {
	var req deployTokenRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*access_refresh_tokens.Payload)

	user, err := server.store.GetUser(ctx, authPayload.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	privateKey, err := util.Decrypt(user.WalletPrivateAddress, secretKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	contractAddress, _ := blockchain.DeployContract(common.HexToAddress(user.WalletPublicAddress), privateKeyECDSA, req.Name, req.Symbol, big.NewInt(req.Supply))

	arg := db.CreateTokenParams{
		UID:             user.ID,
		TokenName:       req.Name,
		Symbol:          req.Symbol,
		Supply:          req.Supply,
		ContractAddress: contractAddress.String(),
	}
	_, err = server.store.CreateToken(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	txArg := db.CreateTransactionParams{
		TransactionType: "deploy",
		FromAddress:     user.WalletPublicAddress,
		ToAddress:       "",
		TransferData:    contractAddress.String(),
		HashValue:       contractAddress.Hash().String(),
	}
	_, err = server.store.CreateTransaction(ctx, txArg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	res := deployTokenResponse{
		TokenAddress: contractAddress.String(),
	}
	ctx.JSON(http.StatusOK, res)
}

type transferTokenRequest struct {
	ContractAddress string `json:"contract_address"`
	ToAddress       string `json:"to_address"`
	Amount          int64  `json:"amount"`
}

func (server *Server) transferToken(ctx *gin.Context) {
	x, _ := ioutil.ReadAll(ctx.Request.Body)
	values := string(x)
	var data transferTokenRequest
	err := json.Unmarshal([]byte(values), &data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*access_refresh_tokens.Payload)
	user, err := server.store.GetUser(ctx, authPayload.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	/* arg := db.GetTokenByUIDAndContractParams{
		UID: user.ID,
		ContractAddress: req.ContractAddress,
	}

	token, err := server.store.GetTokenByUIDAndContract(ctx, arg)
	if token.ContractAddress != req.ContractAddress {
		ctx.JSON(http.StatusBadRequest, "wrong token address")
	}
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation" :
				ctx.JSON(http.StatusForbidden, errorResponse(err))
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	} */

	privateKey, err := util.Decrypt(user.WalletPrivateAddress, secretKey)
	if err != nil {
		log.Fatal(err)
	}

	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(data.ContractAddress)
	toAddress := common.HexToAddress(data.ToAddress)
/* 
	fmt.Println("privateKeyECDSA: ", privateKeyECDSA)
	fmt.Println("contractAddress:", contractAddress)
	fmt.Println("toAddress: ", toAddress)
	fmt.Println("big.NewInt(req.Amount): ", big.NewInt(data.Amount)) */

	hash, err := blockchain.TransferContract(privateKeyECDSA, common.HexToAddress(user.WalletPublicAddress), contractAddress, toAddress, big.NewInt(data.Amount))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	txArg := db.CreateTransactionParams{
		TransactionType: "transfer",
		FromAddress:     user.WalletPublicAddress,
		ToAddress:       data.ToAddress,
		TransferData:    "amount:" + string(rune(data.Amount)) + "token:" + data.ContractAddress,
		HashValue:       hash.String(),
	}
	_, err = server.store.CreateTransaction(ctx, txArg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, "token transferred")
}

type getTokenRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
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

func (server *Server) listTokens(ctx *gin.Context) {

	authPayload := ctx.MustGet(authorizationPayloadKey).(*access_refresh_tokens.Payload)

	user, err := server.store.GetUser(ctx, authPayload.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	tokens, err := server.store.ListTokens(ctx, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, tokens)
}

type getTokenBalanceRequest struct {
	TokenAddress string `json:"tokenaddress"`
}

type getTokenBalanceResponse struct {
	Balance *big.Int `json:"balance"`
}

func (server *Server) getTokenBalance(ctx *gin.Context) {

	x, _ := ioutil.ReadAll(ctx.Request.Body)
	values := string(x)
	var data getTokenBalanceRequest
	err := json.Unmarshal([]byte(values), &data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*access_refresh_tokens.Payload)

	_, err = server.store.GetTokenByAddress(ctx, data.TokenAddress)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	_, err = server.store.GetUser(ctx, authPayload.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	balance, err := blockchain.GetTokenBalance(common.HexToAddress(data.TokenAddress))
	if err != nil {
		fmt.Println("get token err")
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, getTokenBalanceResponse{
		Balance: balance,
	})

}
