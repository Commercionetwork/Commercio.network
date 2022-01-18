package keeper

import (
	"testing"
	"time"

	"github.com/commercionetwork/commercionetwork/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, Keeper, sdk.Context) {
	keeper, ctx := setupKeeper(t)

	return NewMsgServerImpl(*keeper), *keeper, ctx
}

func Test_SetDidDocument(t *testing.T) {
	srv, k, ctx := setupMsgServer(t)

	// create
	dateString := "2019-03-23T06:35:22Z"
	createdTimestamp, err := time.Parse(types.ComplaintW3CTime, dateString)
	require.NoError(t, err)
	ctx = ctx.WithBlockTime(createdTimestamp.UTC())

	sdkCtx := sdk.WrapSDKContext(ctx)

	msg := types.MsgSetDidDocument{
		DidDocument: types.ValidIdentity.DidDocument,
	}

	did := msg.DidDocument.ID

	_, err = k.GetLastIdentityOfAddress(ctx, did)
	assert.Error(t, err)

	resp, err := srv.SetIdentity(sdkCtx, &msg)
	require.NoError(t, err)
	assert.Equal(t, &types.MsgSetDidDocumentResponse{}, resp)

	firstIdentity, err := k.GetLastIdentityOfAddress(ctx, did)
	assert.NoError(t, err)
	require.Equal(t, msg.DidDocument, firstIdentity.DidDocument)
	expectedFirstMetadata := types.Metadata{
		Created: dateString,
		Updated: dateString,
	}
	require.Equal(t, &expectedFirstMetadata, firstIdentity.Metadata)

	// update
	ctx = sdk.UnwrapSDKContext(sdkCtx)
	dateUpdated := "2023-08-10T13:40:06Z"
	assert.True(t, dateUpdated > dateString)
	updatedTimestamp, err := time.Parse(types.ComplaintW3CTime, dateUpdated)
	require.NoError(t, err)
	ctx = ctx.WithBlockTime(updatedTimestamp)

	sdkCtx = sdk.WrapSDKContext(ctx)

	newMsg := msg
	newMsg.DidDocument.AssertionMethod = []string{"#key-1"}

	resp, err = srv.SetIdentity(sdkCtx, &newMsg)
	require.NoError(t, err)
	assert.Equal(t, &types.MsgSetDidDocumentResponse{}, resp)

	identityUpdated, err := k.GetLastIdentityOfAddress(ctx, did)
	assert.NoError(t, err)
	require.Equal(t, newMsg.DidDocument, identityUpdated.DidDocument)

	require.Equal(t, firstIdentity.Metadata.Created, identityUpdated.Metadata.Created)
	require.NotEqual(t, firstIdentity.Metadata.Updated, identityUpdated.Metadata.Updated)

}
