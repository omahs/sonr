package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/sonrhq/core/testutil/keeper"
	"github.com/sonrhq/core/x/identity/keeper"
	"github.com/sonrhq/core/x/identity/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestDidDocumentMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.IdentityKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateDidDocument{Creator: creator,
			Document: types.NewBlankDocument(creator),
		}
		_, err := srv.CreateDidDocument(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetDidDocument(ctx,
			expected.Document.ID,
		)
		accAddr, err := rst.AccAddress()
		require.NoError(t, err)
		require.True(t, found)
		require.Equal(t, expected.Creator, accAddr)
	}
}

func TestDidDocumentMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateDidDocument
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateDidDocument{Creator: creator,
				Document: types.NewBlankDocument(creator),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateDidDocument{Creator: "B",
				Document: types.NewBlankDocument(creator),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateDidDocument{Creator: creator,
				Document: types.NewBlankDocument(creator),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.IdentityKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateDidDocument{Creator: creator,
				Document: types.NewBlankDocument(creator),
			}
			_, err := srv.CreateDidDocument(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateDidDocument(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetDidDocument(ctx,
					expected.Document.ID,
				)
				require.True(t, found)
				accAddr, err := rst.AccAddress()
				require.NoError(t, err)
				require.Equal(t, expected.Creator, accAddr)
			}
		})
	}
}

func TestDidDocumentMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteDidDocument
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteDidDocument{Creator: creator,
				Did: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteDidDocument{Creator: "B",
				Did: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteDidDocument{Creator: creator,
				Did: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.IdentityKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateDidDocument(wctx, &types.MsgCreateDidDocument{Creator: creator,
				Document: types.NewBlankDocument(creator),
			})
			require.NoError(t, err)
			_, err = srv.DeleteDidDocument(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetDidDocument(ctx,
					tc.request.Did,
				)
				require.False(t, found)
			}
		})
	}
}
