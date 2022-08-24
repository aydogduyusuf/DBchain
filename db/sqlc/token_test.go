package db

import (
	"context"
	"testing"
	"time"

	"github.com/aydogduyusuf/DBchain/util"
	"github.com/stretchr/testify/require"
)

func createRandomToken(t *testing.T) Token{
	user := createRandomUser(t)
	
	arg := CreateTokenParams{
		UID: 				user.ID,
		TokenName: 			util.RandomString(6),
		Symbol: 			util.RandomString(3),
		Supply: 			util.RandomMoney(),
		ContractAddress: 	util.RandomString(42),
	}

	token, err := testQueries.CreateToken(context.Background(), arg) 
	require.NoError(t, err)
	require.NotEmpty(t, token)

	require.Equal(t, arg.UID, token.UID)
	require.Equal(t, arg.TokenName, token.TokenName)
	require.Equal(t, arg.Symbol, token.Symbol)
	require.Equal(t, arg.Supply, token.Supply)
	require.Equal(t, arg.ContractAddress, token.ContractAddress)

	require.NotZero(t, token.ID)
	require.NotZero(t, token.IsActive)

	return token
}

func TestCreateToken(t *testing.T) {
	createRandomToken(t)
}

func TestGetToken(t *testing.T) {
	token1 := createRandomToken(t)
	token2, err := testQueries.GetToken(context.Background(), token1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, token2)

	require.Equal(t, token1.ID, 		token2.ID)
	require.Equal(t, token1.UID, 		token2.UID)
	require.Equal(t, token1.TokenName, 	token2.TokenName)
	require.Equal(t, token1.Symbol, 	token1.Symbol)
	require.WithinDuration(t, token1.CreateTime, token2.CreateTime, time.Second)
}

/* func TestUpdateAccount(t *testing.T) {
	token1 := createRandomToken(t)

	arg := UpdateTokenParams {
		ID: token1.ID,
		ContractAddress: "0x4E4a2D84DD82dC5f3fb17a938e963BaB9204a92C",
	}

	token2, err := testQueries.UpdateToken(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, token2)

	require.Equal(t, token1.ID, token2.ID)
	require.Equal(t, token1.UID, token2.UID)
	require.Equal(t, token1.TokenName, token2.TokenName)
	require.Equal(t, token1.Symbol, token2.Symbol)
	require.Equal(t, token1.ID, token2.ID)
	require.Equal(t, token1.Owner, token2.Owner)
	require.Equal(t, arg.Balance, token2.Balance)
	require.Equal(t, token1.Currency, token2.Currency)
	require.WithinDuration(t, token1.CreatedAt, token2.CreatedAt, time.Second)
} */

func TestDeleteToken(t *testing.T) {
	token1 := createRandomToken(t)

	err := testQueries.DeleteToken(context.Background(), token1.ID)
	require.NoError(t, err)

	token1, err = testQueries.GetToken(context.Background(), token1.ID)
	require.NoError(t, err)
	//require.Equal(t, token1.DeleteTime, time.Now())
	require.False(t, token1.IsActive)
}

func TestListTokens(t *testing.T) {
	user := createRandomUser(t)
	var lastToken Token
	for i := 0; i < 10; i++ {
		lastToken = createRandomToken(t)
	}

	tokens, err := testQueries.ListTokens(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, tokens)
	
	for _, token := range tokens {
		require.NotEmpty(t, token)
		require.Equal(t, lastToken.UID, token.UID)
	}
	
}