package db

import (
	"context"
	"testing"

	"github.com/aydogduyusuf/DBchain/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User{
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	publicAdr ,err := util.HashPassword(util.RandomString(42))
	require.NoError(t, err)

	privateAdr ,err := util.HashPassword(util.RandomString(64))
	require.NoError(t, err)

	
	arg := CreateUserParams{
		Username: util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName: util.RandomOwner(),
		Email: util.RandomEmail(),
		WalletPublicAddress: publicAdr,
		WalletPrivateAddress: privateAdr,
	}

	user, err := testQueries.CreateUser(context.Background(), arg) 
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.WalletPublicAddress, user.WalletPublicAddress)
	require.Equal(t, arg.WalletPrivateAddress, user.WalletPrivateAddress)

	require.Empty(t, user.UpdateTime)
	require.Empty(t, user.DeleteTime)
	require.NotZero(t, user.CreateTime)
	require.True(t, user.IsActive)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, 						user2.ID)
	require.Equal(t, user1.Username, 				user2.Username)
	require.Equal(t, user1.HashedPassword, 			user2.HashedPassword)
	require.Equal(t, user1.FullName, 				user2.FullName)
	require.Equal(t, user1.Email, 				 	user2.Email)
	require.Equal(t, user1.WalletPublicAddress, 	user2.WalletPublicAddress)
	require.Equal(t, user1.WalletPrivateAddress, 	user2.WalletPrivateAddress)

	require.Empty(t, user2.UpdateTime)
	require.Empty(t, user2.DeleteTime)
	require.NotZero(t, user2.CreateTime)
	require.True(t, user2.IsActive)
}
