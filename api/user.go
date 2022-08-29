package api

import (
	"database/sql"
	"log"
	"math/big"
	"net/http"
	"time"

	"github.com/aydogduyusuf/DBchain/access_refresh_tokens"
	"github.com/aydogduyusuf/DBchain/blockchain"
	db "github.com/aydogduyusuf/DBchain/db/sqlc"
	"github.com/aydogduyusuf/DBchain/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

var secretKey = []byte("secretsecretsecretsecretsecretsecretse")

type createUserRequest struct {
	Username    string `json:"username" binding:"required,alphanum" `
	Password 	string `json:"password" binding:"required,min=6" `
	Fullname 	string `json:"full_name" binding:"required" `
	Email 		string `json:"email" binding:"required,email" `
}

type userResponse struct {
	Username          string    `json:"username"`
	Fullname          string    `json:"full_name"`
	Email             string    `json:"email"`
	UpdateTime 		  time.Time `json:"update_time"`
	CreateTime        time.Time `json:"create_time"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username: 	user.Username,
		Fullname: 	user.FullName,
		Email: 		user.Email,
		UpdateTime: user.UpdateTime,
		CreateTime: user.CreateTime,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	publicKey, privateKeyECDSA := blockchain.CreateWallet()
	privateKeyBytes := crypto.FromECDSA(privateKeyECDSA)
	privateKey := hexutil.Encode(privateKeyBytes)[2:]
	
	privateKeyEncrypted, err := util.Encrypt([]byte(privateKey), secretKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
 

	arg := db.CreateUserParams{
		Username: 				req.Username,
		HashedPassword: 		hashedPassword,
		FullName: 				req.Fullname,
		Email: 					req.Email,
		WalletPublicAddress: 	publicKey.String(),
		WalletPrivateAddress:   string(privateKeyEncrypted),
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation" : 
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

type loginUserRequest struct {
	Username    string `json:"username" binding:"required,alphanum" `
	Password 	string `json:"password" binding:"required,min=6" `
}

type loginUserResponse struct {
	SessionID					uuid.UUID			`json:"session_id"`
	AccessToken 				string				`json:"access_token"`
	AccessTokenExpiresAt 		time.Time 			`json:"access_token_expires_at"` 
	RefreshToken 				string				`json:"refresh_token"`
	RefreshTokenExpiresAt 		time.Time 			`json:"refresh_token_expires_at"` 
	User 						userResponse		`json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(user.Username, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID: 				refreshPayload.ID,
		Username:			user.Username,
		RefreshToken:		refreshToken,
		UserAgent: 			ctx.Request.UserAgent(),   
		ClientIp:     		ctx.ClientIP(),
		IsBlocked:			false,   
		ExpiresAt:			refreshPayload.ExpiredAt,    
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginUserResponse{
		SessionID: 					session.ID,
		AccessToken: 				accessToken,
		AccessTokenExpiresAt: 		accessPayload.ExpiredAt,
		RefreshToken: 				refreshToken,
		RefreshTokenExpiresAt: 		refreshPayload.ExpiredAt,
		User: 						newUserResponse(user),
	}
	
	ctx.JSON(http.StatusOK, rsp)
}

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